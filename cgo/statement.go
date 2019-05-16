package cgo

//#include <stdlib.h>
//#include "ctlib.h"
import "C"
import (
	"context"
	"database/sql/driver"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/SAP/go-ase/libase"
	"github.com/SAP/go-ase/libase/types"
)

type statement struct {
	name     string
	argCount int
	cmd      *Command
}

// Interface satisfaction checks
var (
	_ driver.Stmt             = (*statement)(nil)
	_ driver.StmtExecContext  = (*statement)(nil)
	_ driver.StmtQueryContext = (*statement)(nil)
)

var (
	statementCounter  uint
	statementCounterM = sync.Mutex{}
)

func (conn *Connection) Prepare(query string) (driver.Stmt, error) {
	return conn.prepare(query)
}

func (conn *Connection) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	recvStmt := make(chan driver.Stmt, 1)
	recvErr := make(chan error, 1)
	go func() {
		stmt, err := conn.prepare(query)
		recvStmt <- stmt
		close(recvStmt)
		recvErr <- err
		close(recvErr)
	}()

	for {
		select {
		case <-ctx.Done():
			go func() {
				stmt := <-recvStmt
				if stmt != nil {
					stmt.Close()
				}
			}()
			return nil, ctx.Err()
		case err := <-recvErr:
			stmt := <-recvStmt
			return stmt, err
		}
	}
}

func (conn *Connection) prepare(query string) (driver.Stmt, error) {
	stmt := &statement{}

	stmt.argCount = strings.Count(query, "?")

	statementCounterM.Lock()
	statementCounter++
	stmt.name = fmt.Sprintf("stmt%d", statementCounter)
	statementCounterM.Unlock()

	cmd, err := conn.Dynamic(stmt.name, query)
	if err != nil {
		stmt.Close()
		return nil, err
	}

	for err = nil; err != io.EOF; _, _, err = cmd.Response() {
		if err != nil {
			stmt.Close()
			cmd.Cancel()
			return nil, err
		}
	}

	stmt.cmd = cmd

	return stmt, nil
}

func (stmt *statement) Close() error {
	if stmt.cmd != nil {
		name := C.CString(stmt.name)
		defer C.free(unsafe.Pointer(name))

		retval := C.ct_dynamic(stmt.cmd.cmd, C.CS_DEALLOC, name, C.CS_NULLTERM, nil, C.CS_UNUSED)
		if retval != C.CS_SUCCEED {
			return makeError(retval, "C.ct_dynamic with C.CS_DEALLOC failed")
		}

		retval = C.ct_send(stmt.cmd.cmd)
		if retval != C.CS_SUCCEED {
			return makeError(retval, "C.ct_send failed")
		}

		var err error
		for err = nil; err != io.EOF; _, _, err = stmt.cmd.Response() {
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (stmt *statement) NumInput() int {
	return stmt.argCount
}

func (stmt *statement) exec(args []driver.NamedValue) error {
	if len(args) != stmt.argCount {
		return fmt.Errorf("Mismatched argument count - expected %d, got %d",
			stmt.argCount, len(args))
	}

	name := C.CString(stmt.name)
	defer C.free(unsafe.Pointer(name))

	retval := C.ct_dynamic(stmt.cmd.cmd, C.CS_EXECUTE, name, C.CS_NULLTERM, nil, C.CS_UNUSED)
	if retval != C.CS_SUCCEED {
		return makeError(retval, "C.ct_dynamic with CS_EXECUTE failed")
	}

	for i, arg := range args {
		datafmt := (*C.CS_DATAFMT)(C.calloc(1, C.sizeof_CS_DATAFMT))
		defer C.free(unsafe.Pointer(datafmt))
		datafmt.status = C.CS_INPUTVALUE
		datafmt.namelen = C.CS_NULLTERM
		asetype, err := types.FromGoType(arg.Value)
		if err != nil {
			return fmt.Errorf("Failed to retrieve ASEType for driver.Value: %v", err)
		}
		datafmt.datatype = (C.CS_INT)(asetype)

		datalen := 0
		var ptr unsafe.Pointer
		switch value := arg.Value.(type) {
		case int64:
			i := (C.CS_BIGINT)(value)
			ptr = unsafe.Pointer(&i)
		case uint64:
			i := (C.CS_UBIGINT)(value)
			ptr = unsafe.Pointer(&i)
		case float64:
			i := (C.CS_FLOAT)(value)
			ptr = unsafe.Pointer(&i)
		case bool:
			b := (C.CS_BOOL)(0)
			if value {
				b = (C.CS_BOOL)(1)
			}
			ptr = unsafe.Pointer(&b)
			datalen = 1
		case []byte:
			if len(value) == 0 {
				ptr = C.CBytes([]byte{})
				defer C.free(ptr)
				datalen = 0
			} else {
				ptr = C.CBytes(arg.Value.([]byte))
				defer C.free(ptr)
				datalen = len(arg.Value.([]byte))
			}
			datafmt.format = C.CS_FMT_PADNULL
		case string:
			if len(value) == 0 {
				ptr = unsafe.Pointer(C.CString(""))
				defer C.free(ptr)
				datalen = 0
			} else {
				ptr = unsafe.Pointer(C.CString(value))
				defer C.free(ptr)
				datalen = len(value)
			}
			datafmt.format = C.CS_FMT_NULLTERM
			datafmt.maxlength = C.CS_MAX_CHAR
		case time.Time:
			microseconds := (C.CS_UBIGINT)(libase.TimeToMicroseconds(arg.Value.(time.Time)))
			ptr = unsafe.Pointer(&microseconds)
		default:
			return fmt.Errorf("Unable to transform to Client-Library: %v", arg.Value)
		}

		var csDatalen C.CS_INT
		if datalen != C.CS_UNUSED {
			csDatalen = (C.CS_INT)(datalen)
		} else {
			csDatalen = C.CS_UNUSED
		}

		retval = C.ct_param(stmt.cmd.cmd, datafmt, ptr, csDatalen, 0)
		if retval != C.CS_SUCCEED {
			return makeError(retval, "C.ct_param on parameter %d failed with argument '%v'", i, arg)
		}
	}

	retval = C.ct_send(stmt.cmd.cmd)
	if retval != C.CS_SUCCEED {
		return makeError(retval, "C.ct_send failed")
	}

	return nil
}

func (stmt *statement) execContext(ctx context.Context, args []driver.NamedValue) error {
	recvErr := make(chan error, 1)
	go func() {
		recvErr <- stmt.exec(args)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-recvErr:
			return err
		}
	}
}

func (stmt *statement) execResults() (driver.Result, error) {
	var resResult driver.Result

	for {
		_, result, err := stmt.cmd.Response()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if result != nil {
			resResult = result
		}
	}

	return resResult, nil
}

func (stmt *statement) Exec(args []driver.Value) (driver.Result, error) {
	err := stmt.exec(libase.ValuesToNamedValues(args))
	if err != nil {
		return nil, err
	}

	return stmt.execResults()
}

func (stmt *statement) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	err := stmt.execContext(ctx, args)
	if err != nil {
		return nil, err
	}

	recvResult := make(chan driver.Result, 1)
	recvErr := make(chan error, 1)
	go func() {
		res, err := stmt.execResults()
		recvResult <- res
		close(recvResult)
		recvErr <- err
	}()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case err := <-recvErr:
			res := <-recvResult
			return res, err
		}
	}
}

func (stmt *statement) Query(args []driver.Value) (driver.Rows, error) {
	err := stmt.exec(libase.ValuesToNamedValues(args))
	if err != nil {
		return nil, err
	}

	rows, _, err := stmt.cmd.Response()
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (stmt *statement) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	err := stmt.execContext(ctx, args)
	if err != nil {
		return nil, err
	}

	recvRows := make(chan driver.Rows, 1)
	recvErr := make(chan error, 1)
	go func() {
		rows, _, err := stmt.cmd.Response()
		recvRows <- rows
		close(recvRows)
		recvErr <- err
		close(recvErr)
	}()

	for {
		select {
		case <-ctx.Done():
			go func() {
				rows := <-recvRows
				if rows != nil {
					rows.Close()
				}
			}()
			return nil, ctx.Err()
		case err := <-recvErr:
			rows := <-recvRows
			return rows, err
		}
	}
}

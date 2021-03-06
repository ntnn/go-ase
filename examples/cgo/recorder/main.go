// SPDX-FileCopyrightText: 2020 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// This example shows how the cgo.MessageRecorder can be used to process
// messages from the TDS server.
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SAP/go-ase/cgo"
	"github.com/SAP/go-ase/libase/libdsn"
)

func main() {
	err := DoMain()
	if err != nil {
		log.Fatal(err)
	}
}

func DoMain() error {
	dsn, err := libdsn.NewInfoFromEnv("")
	if err != nil {
		return fmt.Errorf("error reading DSN info from env: %w", err)
	}

	fmt.Println("Opening database")
	db, err := sql.Open("ase", dsn.AsSimple())
	if err != nil {
		return fmt.Errorf("failed to open connection to database: %w", err)
	}
	defer db.Close()

	// Execute a ping here - the connection through database/sql will
	// only be created once a query is performed.
	// This causes the server to send messages regarding the context
	// switches, which we do not want to test for.
	// To prevent the context switch messages being recorded a query is
	// performed before attaching the recorder to the message broker.
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Creating MessageRecorder")
	recorder := cgo.NewMessageRecorder()
	fmt.Println("Registering handler with server message broker")
	cgo.GlobalServerMessageBroker.RegisterHandler(recorder.HandleMessage)

	fmt.Println("Enable traceflag 3604")
	_, err = db.Exec("dbcc traceon(3604)")
	if err != nil {
		return fmt.Errorf("failed to enable traceflag 3604: %w", err)
	}

	fmt.Println("Received messages:")
	for _, line := range recorder.Text() {
		fmt.Print(line)
	}

	fmt.Println("Listing enabled traceflags")
	recorder.Reset()
	_, err = db.Exec("dbcc traceflags")
	if err != nil {
		return fmt.Errorf("failed to list traceflags: %w", err)
	}

	fmt.Println("Received messages:")
	for _, line := range recorder.Text() {
		fmt.Print(line)
	}

	return nil
}

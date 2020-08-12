package tds

import "fmt"

type ErrorPackage struct {
	Length       uint16
	ErrorNumber  int32
	State        uint8
	Class        uint8
	MsgLength    uint16
	ErrorMsg     string
	ServerLength uint8
	ServerName   string
	ProcLength   uint8
	ProcName     string
	LineNr       uint16
}

func (pkg *ErrorPackage) ReadFrom(ch BytesChannel) error {
	var err error
	pkg.Length, err = ch.Uint16()
	if err != nil {
		return ErrNotEnoughBytes
	}

	pkg.ErrorNumber, err = ch.Int32()
	if err != nil {
		return ErrNotEnoughBytes
	}

	pkg.MsgLength, err = ch.Uint16()
	if err != nil {
		return ErrNotEnoughBytes
	}

	pkg.ErrorMsg, err = ch.String(int(pkg.MsgLength))
	if err != nil {
		return ErrNotEnoughBytes
	}

	pkg.ServerLength, err = ch.Uint8()
	if err != nil {
		return ErrNotEnoughBytes
	}

	pkg.ServerName, err = ch.String(int(pkg.ServerLength))
	if err != nil {
		return ErrNotEnoughBytes
	}

	pkg.ProcLength, err = ch.Uint8()
	if err != nil {
		return ErrNotEnoughBytes
	}

	pkg.ProcName, err = ch.String(int(pkg.ProcLength))
	if err != nil {
		return ErrNotEnoughBytes
	}

	pkg.LineNr, err = ch.Uint16()
	if err != nil {
		return ErrNotEnoughBytes
	}

	return nil
}

func (pkg ErrorPackage) WriteTo(ch BytesChannel) error {
	err := ch.WriteByte(byte(TDS_ERROR))
	if err != nil {
		return fmt.Errorf("failed to write TDS Token %s: %w", TDS_ERROR, err)
	}

	err = ch.WriteUint16(pkg.Length)
	if err != nil {
		return fmt.Errorf("failed to write length: %w", err)
	}

	if err := ch.WriteInt32(pkg.ErrorNumber); err != nil {
		return fmt.Errorf("failed to write error number: %w", err)
	}

	if err := ch.WriteUint8(pkg.State); err != nil {
		return fmt.Errorf("failed to write state: %w", err)
	}

	if err := ch.WriteUint8(pkg.Class); err != nil {
		return fmt.Errorf("failed to write class: %w", err)
	}

	if err := ch.WriteUint16(uint16(len(pkg.ErrorMsg))); err != nil {
		return fmt.Errorf("failed to write error message length: %w", err)
	}

	if err := ch.WriteString(pkg.ErrorMsg); err != nil {
		return fmt.Errorf("failed to write error message: %w", err)
	}

	if err := ch.WriteUint8(uint8(len(pkg.ServerName))); err != nil {
		return fmt.Errorf("failed to write servername length: %w", err)
	}

	if err := ch.WriteString(pkg.ServerName); err != nil {
		return fmt.Errorf("failed to write servername: %w", err)
	}

	if err := ch.WriteUint8(uint8(len(pkg.ProcName))); err != nil {
		return fmt.Errorf("failed to write procname length: %w", err)
	}

	if err := ch.WriteString(pkg.ProcName); err != nil {
		return fmt.Errorf("failed to write procname: %w", err)
	}

	if err := ch.WriteUint16(pkg.LineNr); err != nil {
		return fmt.Errorf("failed to write linenr: %w", err)
	}

	return nil
}

func (pkg ErrorPackage) String() string {
	return fmt.Sprintf("%T(%d: %s)", pkg, pkg.ErrorNumber, pkg.ErrorMsg)
}

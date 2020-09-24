// SPDX-FileCopyrightText: 2020 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

package tds

import (
	"bytes"
	"fmt"
)

// TokenlessType defines the type of a TokenlessPackage
type TokenlessType int

const (
	// TokenlessPackage is a data container, e.g. for login buffers
	DataPackage TokenlessType = iota
	// TokenlessPackage signals the end of a payload
	EOMPackage
)

// TokenlessPackage is used to store tokenless payloads from the server
// as well as internal communication in go-ase.
// The type of a TokenlessPackage is defined by the TokenlessType.
type TokenlessPackage struct {
	Type TokenlessType
	// Data is only valid if Type is DataPackage.
	Data *bytes.Buffer
}

func NewTokenlessPackage() *TokenlessPackage {
	return &TokenlessPackage{
		Data: &bytes.Buffer{},
	}
}

func (pkg *TokenlessPackage) ReadFrom(ch BytesChannel) error {
	if pkg.Type != DataPackage {
		return nil
	}
	_, err := pkg.Data.ReadFrom(ch)
	return err
}

func (pkg TokenlessPackage) WriteTo(ch BytesChannel) error {
	if pkg.Type != DataPackage {
		return nil
	}
	return ch.WriteBytes(pkg.Data.Bytes())
}

func (pkg TokenlessPackage) String() string {
	var possibleToken byte = 0
	if pkg.Data != nil && len(pkg.Data.Bytes()) > 0 {
		possibleToken = pkg.Data.Bytes()[0]
	}

	return fmt.Sprintf("%T(%s, possibleToken=%x) %#v", pkg, pkg.Type, possibleToken, pkg)
}

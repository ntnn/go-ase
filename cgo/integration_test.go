// SPDX-FileCopyrightText: 2020 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// +build integration

package cgo

import (
	"fmt"
	"log"
	"testing"

	"github.com/SAP/go-ase/libase/integration"
)

func TestMain(m *testing.M) {
	err := testMain(m)
	if err != nil {
		log.Fatal(err)
	}
}

func testMain(m *testing.M) error {
	GlobalServerMessageBroker.RegisterHandler(genMessageHandler())
	GlobalClientMessageBroker.RegisterHandler(genMessageHandler())

	simpleDSN, simpleTeardown, err := integration.DSN(false)
	if err != nil {
		return fmt.Errorf("error setting up simple DSN: %w", err)
	}
	defer simpleTeardown()

	if err := integration.RegisterDSN("username password", simpleDSN, NewConnector); err != nil {
		return fmt.Errorf("error setting up simple database: %w", err)
	}

	userstoreDSN, userstoreTeardown, err := integration.DSN(true)
	if err != nil {
		return fmt.Errorf("error setting up userstore DSN: %w", err)
	}
	defer userstoreTeardown()

	if err := integration.RegisterDSN("userstorekey", userstoreDSN, NewConnector); err != nil {
		return fmt.Errorf("error setting up userstore database: %w", err)
	}

	if rc := m.Run(); rc != 0 {
		return fmt.Errorf("tests failed with %d", rc)
	}

	return nil
}

// Exact numeric integer
func TestBigInt(t *testing.T)           { integration.DoTestBigInt(t) }
func TestInt(t *testing.T)              { integration.DoTestInt(t) }
func TestSmallInt(t *testing.T)         { integration.DoTestSmallInt(t) }
func TestTinyInt(t *testing.T)          { integration.DoTestTinyInt(t) }
func TestUnsignedBigInt(t *testing.T)   { integration.DoTestUnsignedBigInt(t) }
func TestUnsignedInt(t *testing.T)      { integration.DoTestUnsignedInt(t) }
func TestUnsignedSmallInt(t *testing.T) { integration.DoTestUnsignedSmallInt(t) }

// Exact numeric decimal
func TestDecimal(t *testing.T)     { integration.DoTestDecimal(t) }
func TestDecimal10(t *testing.T)   { integration.DoTestDecimal10(t) }
func TestDecimal380(t *testing.T)  { integration.DoTestDecimal380(t) }
func TestDecimal3838(t *testing.T) { integration.DoTestDecimal3838(t) }

// Approximate numeric
func TestFloat(t *testing.T) { integration.DoTestFloat(t) }
func TestReal(t *testing.T)  { integration.DoTestReal(t) }

// Money
func TestMoney(t *testing.T)  { integration.DoTestMoney(t) }
func TestMoney4(t *testing.T) { integration.DoTestMoney4(t) }

// Date and time
func TestDate(t *testing.T)          { integration.DoTestDate(t) }
func TestTime(t *testing.T)          { integration.DoTestTime(t) }
func TestSmallDateTime(t *testing.T) { integration.DoTestSmallDateTime(t) }
func TestDateTime(t *testing.T)      { integration.DoTestDateTime(t) }
func TestBigDateTime(t *testing.T)   { integration.DoTestBigDateTime(t) }
func TestBigTime(t *testing.T)       { integration.DoTestBigTime(t) }

// Character
func TestVarChar(t *testing.T)  { integration.DoTestVarChar(t) }
func TestChar(t *testing.T)     { integration.DoTestChar(t) }
func TestNChar(t *testing.T)    { integration.DoTestNChar(t) }
func TestNVarChar(t *testing.T) { integration.DoTestNVarChar(t) }
func TestText(t *testing.T)     { integration.DoTestText(t) }
func TestUniChar(t *testing.T)  { integration.DoTestUniChar(t) }
func TestUniText(t *testing.T)  { integration.DoTestUniText(t) }

// Binary
func TestBinary(t *testing.T)    { integration.DoTestBinary(t) }
func TestVarBinary(t *testing.T) { integration.DoTestVarBinary(t) }

// Bit
func TestBit(t *testing.T) { integration.DoTestBit(t) }

// Image
func TestImage(t *testing.T) { integration.DoTestImage(t) }

// Routines
func TestSQLTx(t *testing.T)   { integration.DoTestSQLTx(t) }
func TestSQLExec(t *testing.T) { integration.DoTestSQLExec(t) }

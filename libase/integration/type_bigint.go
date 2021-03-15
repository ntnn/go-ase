// SPDX-FileCopyrightText: 2020 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by "gen_type BigInt int64"; DO NOT EDIT.

package integration

import (
	"database/sql"

	"testing"
)

// DoTestBigInt tests the handling of the BigInt.
func DoTestBigInt(t *testing.T) {
	TestForEachDB("TestBigInt", t, testBigInt)
	//
}

func testBigInt(t *testing.T, db *sql.DB, tableName string) {
	pass := make([]interface{}, len(samplesBigInt))
	mySamples := make([]int64, len(samplesBigInt))

	for i, sample := range samplesBigInt {

		mySample := sample

		pass[i] = mySample
		mySamples[i] = mySample
	}

	rows, teardownFn, err := SetupTableInsert(db, tableName, "bigint", pass...)
	if err != nil {
		t.Errorf("Error preparing table: %v", err)
		return
	}
	defer rows.Close()
	defer teardownFn()

	i := 0
	var recv int64
	for rows.Next() {
		if err := rows.Scan(&recv); err != nil {
			t.Errorf("Scan failed on %dth scan: %v", i, err)
			continue
		}

		if recv != mySamples[i] {

			t.Errorf("Received value does not match passed parameter")
			t.Errorf("Expected: %v", mySamples[i])
			t.Errorf("Received: %v", recv)
		}

		i++
	}

	if err := rows.Err(); err != nil {
		t.Errorf("Error preparing rows: %v", err)
	}

	if i != len(pass) {
		t.Errorf("Only read %d values from database, expected to read %d", i, len(pass))
	}
}
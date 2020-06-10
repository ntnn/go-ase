package libtest

import (
	"database/sql"

	"testing"
)

// DoTestUnsignedBigInt tests the handling of the UnsignedBigInt.
func DoTestUnsignedBigInt(t *testing.T) {
	TestForEachDB("TestUnsignedBigInt", t, testUnsignedBigInt)
	//
}

func testUnsignedBigInt(t *testing.T, db *sql.DB, tableName string) {
	pass := make([]interface{}, len(samplesUnsignedBigInt))
	mySamples := make([]uint64, len(samplesUnsignedBigInt))

	for i, sample := range samplesUnsignedBigInt {

		mySample := sample

		pass[i] = mySample
		mySamples[i] = mySample
	}

	rows, teardownFn, err := SetupTableInsert(db, tableName, "unsigned bigint", pass...)
	if err != nil {
		t.Errorf("Error preparing table: %v", err)
		return
	}
	defer rows.Close()
	defer teardownFn()

	i := 0
	var recv uint64
	for rows.Next() {
		err = rows.Scan(&recv)
		if err != nil {
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
}

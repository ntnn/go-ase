// SPDX-FileCopyrightText: 2020 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	"github.com/SAP/go-ase/libase/libdsn"
	_ "github.com/SAP/go-ase/purego"
)

// This example shows how to use transactions using the database/sql
// interface and the pure go driver.

func main() {
	if err := DoMain(); err != nil {
		log.Fatalf("transaction: %v", err)
	}
}

func DoMain() error {
	dsn, err := libdsn.NewInfoFromEnv("")
	if err != nil {
		return fmt.Errorf("error reading DSN info from env: %w", err)
	}

	db, err := sql.Open("ase", dsn.AsSimple())
	if err != nil {
		return fmt.Errorf("failed to open connection to database: %w", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("error closing database: %v", err)
		}
	}()

	fmt.Println("creating table simple")
	if _, err := db.Exec("if object_id('simple') is not null drop table simple"); err != nil {
		return fmt.Errorf("failed to drop table 'simple': %w", err)
	}

	if _, err := db.Exec("create table simple (a int, b char(30))"); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	defer func() {
		if _, err := db.Exec("drop table simple"); err != nil {
			log.Printf("failed to drop table: %v", err)
		}
	}()

	fmt.Println("opening transaction")
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error creating transaction: %w", err)
	}

	fmt.Println("inserting values into simple")
	if _, err := tx.Exec("insert into simple (a, b) values (?, ?)", math.MaxInt32, "a string"); err != nil {
		return fmt.Errorf("failed to insert values: %w", err)
	}

	fmt.Println("preparing statement")
	stmt, err := tx.Prepare("select * from simple where a=?")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("error closing statement: %v", err)
		}
	}()

	fmt.Println("executing prepared statement")
	rows, err := stmt.Query(math.MaxInt32)
	if err != nil {
		return fmt.Errorf("error querying with prepared statement: %w", err)
	}

	var a int
	var b string
	for rows.Next() {
		if err := rows.Scan(&a, &b); err != nil {
			return fmt.Errorf("error scanning row: %w", err)
		}

		fmt.Printf("a: %d\n", a)
		fmt.Printf("b: %s\n", b)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows reported error: %w", err)
	}

	fmt.Println("committing transaction")
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/SAP/go-ase/cgo"
	"github.com/SAP/go-ase/libase/libdsn"
)

func main() {
	err := doMain()
	if err != nil {
		log.Printf("Failed: %v", err)
		os.Exit(1)
	}
}

type devInfo struct {
	name    string
	phyname string
	crdate  time.Time
	size    int64
}

func doMain() error {
	dsn := libdsn.NewDsnInfoFromEnv("")

	fmt.Println("Opening database")
	db, err := sql.Open("ase", dsn.AsSimple())
	if err != nil {
		return fmt.Errorf("Failed to open connection to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connect to master")
	_, err = db.Exec("use master")
	if err != nil {
		return fmt.Errorf("Failed to switch to master database: %v", err)
	}

	query := "select sysdevices.vdevno, sysdevices.name, sysusages.size, sysusages.unreservedpgs from sysdevices, sysusages where sysdevices.vdevno = sysusages.vdevno"

	fmt.Println("Querying devices")
	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("Failed to query devices")
	}
	defer rows.Close()

	colNames, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("Failed to retrieve column names: %v", err)
	}

	colNamesI := make([]interface{}, len(colNames))
	for i, name := range colNames {
		colNamesI[i] = name
	}

	// format := "| %s | %s | %s | %d | %d | %d | %d | %d |\n"
	format := "| %-20s | %-50s | %-40s | %-7d |\n"

	fmt.Println("Printing device information")
	fmt.Printf("| %-20s | %-50s | %-40s | %-7s |\n", colNamesI...)

	dev := devInfo{}

	for rows.Next() {
		err = rows.Scan(&dev.name, &dev.phyname, &dev.crdate, &dev.size)
		if err != nil {
			return fmt.Errorf("Failed to scan row: %v", err)
		}

		fmt.Printf(format, dev.name, dev.phyname, dev.crdate, dev.size)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("Error reading rows: %v", err)
	}

	return nil
}

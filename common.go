package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

var localDB = map[string]string{
	"service":  "XE",
	"username": "demo",
	"server":   "localhost",
	"port":     "1521",
	"password": "demo",
}

var autonomousDB = map[string]string{
	"walletLocation": "your wallet path",
	"tns_entry":	  "your tns entry: dbxxxxdb2022xxx_low",
	"username":       "your schema/username",
	"password":       "your password",
}

func main() {
	fmt.Println("****************Part 1*******************")
	fmt.Println("*** Using Instant Client & GoDrOr package - Row by row")
	t := time.Now()
	doDBThingsThroughInstantClient(autonomousDB)
	fmt.Println("Time Elapsed", time.Now().Sub(t).Milliseconds())
	fmt.Println(" ")
	fmt.Println("****************Part 2*******************")
	fmt.Println("*** Using Instant Client & GoDrOr package - Array")
	t = time.Now()
	doDBThingsThroughInstantClientArray(autonomousDB)
	fmt.Println("Time Elapsed", time.Now().Sub(t).Milliseconds())
}

func handleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

const createTableStatement = "CREATE TABLE TEMP_TABLE ( NAME VARCHAR2(100), CREATION_TIME TIMESTAMP DEFAULT SYSTIMESTAMP, VALUE  NUMBER(5))"
const dropTableStatement = "DROP TABLE TEMP_TABLE PURGE"
const insertStatement = "INSERT INTO TEMP_TABLE ( NAME , VALUE) VALUES (:name, :value)"

func someAdditionalActions(db *sql.DB) {

	var queryResultColumnOne string
	row := db.QueryRow("SELECT systimestamp FROM dual")
	err := row.Scan(&queryResultColumnOne)
	if err != nil {
		panic(fmt.Errorf("error scanning db: %w", err))
	}
	fmt.Println("The time in the database ", queryResultColumnOne)
	//_, err = db.Exec(createTableStatement)
	//handleError("create table", err)
	//defer db.Exec(dropTableStatement)
	stmt, err := db.Prepare(insertStatement)
	handleError("prepare insert statement", err)
	sqlresult, err := stmt.Exec("John", 42)
	handleError("execute insert statement", err)
	rowCount, _ := sqlresult.RowsAffected()
	fmt.Println("Inserted number of rows = ", rowCount)
	//
	sqlresult, err = stmt.Exec("Jane", 69)
	handleError("execute insert statement", err)
	rowCount, _ = sqlresult.RowsAffected()
	fmt.Println("Inserted number of rows = ", rowCount)
	//
	sqlresult, err = stmt.Exec("Malcolm", 13)
	handleError("execute insert statement", err)
	rowCount, _ = sqlresult.RowsAffected()
	fmt.Println("Inserted number of rows = ", rowCount)

}

func someAdditionalActionsArray(db *sql.DB) {

	var queryResultColumnOne string
	row := db.QueryRow("SELECT systimestamp FROM dual")
	err := row.Scan(&queryResultColumnOne)
	if err != nil {
		panic(fmt.Errorf("error scanning db: %w", err))
	}
	fmt.Println("The time in the database ", queryResultColumnOne)
	//_, err = db.Exec(createTableStatement)
	handleError("create table", err)
	//defer db.Exec(dropTableStatement)
	stmt, err := db.Prepare(insertStatement)
	handleError("prepare insert statement", err)
	sqlresult, err := stmt.Exec([]string{"John", "Jane", "Malcolm"},  []int{42,69,13})
	handleError("execute insert statement", err)
	rowCount, _ := sqlresult.RowsAffected()
	fmt.Println("Inserted number of rows = ", rowCount)

}

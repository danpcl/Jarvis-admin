package main

import (
	"container/list"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func GetRecentSchedule(machineType string) string {
	collection := list.New()
	db, err := sql.Open("sqlite3", "./procschedule.sqlite")
	checkErr(err)

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM jobscheduling WHERE 'Job Name' IN (SELECT 'Job Name' FROM jobscheduling WHERE Machine == '%s')", machineType))
	checkErr(err)

	for rows.Next() {
		var date time.Time
		var cnt int32
		var name string
		var machine string
		var thickness float64
		var material string

		err = rows.Scan(&date, &cnt, &name, &machine, &thickness, &material)
		checkErr(err)
		collection.PushBack(name)
	}
	db.Close()

	return collection.Back().Value.(string)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

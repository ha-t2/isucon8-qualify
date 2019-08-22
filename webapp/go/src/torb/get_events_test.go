package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
)

func TestGetEvents(t *testing.T) {
	event1 := Event{ID: 10, Title: "瞳を澄ませば",
		PublicFg: true, ClosedFg: false,
		Price: 3000}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		"isucon", "isucon", "127.0.0.1", "3306", "torb",
	)

	fmt.Println(dsn)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	events, err := getEvents(false)
	fmt.Println(events[0].ID)
	fmt.Println(events[1])
	fmt.Println(events[1].Sheets["S"])
	if err == nil {
		t.Fatalf("failed test %#v", err)
	}

	if event1.Title != events[1].Title {
		t.Fatalf("failed test %v != %v", event1.Title, events[1].Title)
	}
}

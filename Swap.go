package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Seat struct {
	ID      int
	Student string
}

func SwapSeat(db *sql.DB) ([]Seat, error) {
	var seats []Seat
	rows, err := db.Query("SELECT id, student FROM Seat ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var seat Seat
		err := rows.Scan(&seat.ID, &seat.Student)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}

	if len(seats)%2 == 1 {
		return seats, nil
	}

	for i := 0; i < len(seats)-1; i += 2 {
		seats[i].ID, seats[i+1].ID = seats[i+1].ID, seats[i].ID
	}

	return seats, nil
}

func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	seats, err := SwapSeat(db)
	if err != nil {
		panic(err.Error())
	}

	for _, seat := range seats {
		fmt.Println(seat.ID, seat.Student)
	}
}
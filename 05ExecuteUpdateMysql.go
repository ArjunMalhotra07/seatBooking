package main

import (
	"log"
)

func executeUpdatedSQL(seatID string, bookedStatus int) {

	_, err := db.Exec(`UPDATE bookSeat SET  bookedStatus = ? WHERE seatID = ?`,
		bookedStatus,
		seatID)

	if err != nil {
		log.Fatal(err)
	}
}

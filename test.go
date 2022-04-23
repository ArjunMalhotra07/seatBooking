package main

import (
	"fmt"
	"strings"
)

func bookFlightSeat() {

	p := fmt.Println
	seats := seating()
	seatNAMES(seats)
	p()
	// p("Enter the VACANT Seat ID to book  --- ")
	// var seatidNumber string
	// fmt.Scan(&seatidNumber)

}

func seatNAMES(vacantSeats []seatType) {
	f := fmt.Print
	p := fmt.Println
	var eachSeat []string
	var tempBookedStatus []int
	for i := 0; i < len(vacantSeats); i++ {
		eachSeat = append(eachSeat, vacantSeats[i].SeatID)
		tempBookedStatus = append(tempBookedStatus, vacantSeats[i].BookedStatus)

	}
	f(eachSeat)
	p()
	f(tempBookedStatus)
	p()
	p()
	p("		  SEATS")

	printLoop(0, 6, eachSeat, tempBookedStatus, "A4")
	printLoop(6, 12, eachSeat, tempBookedStatus, "B4")
	printLoop(12, 18, eachSeat, tempBookedStatus, "C4")
}

func printLoop(m, n int, eachSeat []string, tempBookedStatus []int, seatID string) {
	f := fmt.Print
	p := fmt.Println
	p()
	for i := m; i < n; i++ {

		if tempBookedStatus[i] == 1 {
			if strings.Contains(eachSeat[i], seatID) {
				f("         --  ")
			} else {
				f("  --  ")
			}

		} else {
			if strings.Contains(eachSeat[i], seatID) {
				f("         ", eachSeat[i], "  ")
			} else {
				f("  ", eachSeat[i], "  ")
			}
		}

	}
	p()
}

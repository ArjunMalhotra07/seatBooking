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

	p()
	for i := 0; i < 6; i++ {

		if tempBookedStatus[i] == 1 {
			if strings.Contains(eachSeat[i], "A4") {
				f("       --  ")
			} else {
				f("  --  ")
			}

		} else {
			if strings.Contains(eachSeat[i], "A4") {
				f("       ", eachSeat[i], "  ")
			} else {
				f("  ", eachSeat[i], "  ")
			}
		}

	}
	p()
	for i := 6; i < 12; i++ {

		if tempBookedStatus[i] == 1 {
			if strings.Contains(eachSeat[i], "B4") {
				f("       --  ")
			} else {
				f("  --  ")
			}

		} else {
			if strings.Contains(eachSeat[i], "B4") {
				f("       ", eachSeat[i], "  ")
			} else {
				f("  ", eachSeat[i], "  ")
			}
		}

	}
	p()
	for i := 12; i < 18; i++ {

		if tempBookedStatus[i] == 1 {
			if strings.Contains(eachSeat[i], "C4") {
				f("       --  ")
			} else {
				f("  --  ")
			}

		} else {
			if strings.Contains(eachSeat[i], "C4") {
				f("       ", eachSeat[i], "  ")
			} else {
				f("  ", eachSeat[i], "  ")
			}
		}

	}
	p()
}

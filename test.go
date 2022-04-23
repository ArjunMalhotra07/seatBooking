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
	var filledORNot []string
	var tempBookedStatus []int
	for i := 0; i < len(vacantSeats); i++ {
		filledORNot = append(filledORNot, vacantSeats[i].SeatID)
		tempBookedStatus = append(tempBookedStatus, vacantSeats[i].BookedStatus)

	}
	f(filledORNot)
	p()
	f(tempBookedStatus)
	p()
	p()
	p("		  SEATS")

	p()
	for i := 0; i < 6; i++ {

		if tempBookedStatus[i] == 1 {
			if strings.Contains(filledORNot[i], "A4") {
				f("       --  ")
			} else {
				f("  --  ")
			}

		} else {
			if strings.Contains(filledORNot[i], "A4") {
				f("       ", filledORNot[i], "  ")
			} else {
				f("  ", filledORNot[i], "  ")
			}
		}

	}
	p()
	for i := 6; i < 12; i++ {

		if tempBookedStatus[i] == 1 {
			if strings.Contains(filledORNot[i], "B4") {
				f("       --  ")
			} else {
				f("  --  ")
			}

		} else {
			if strings.Contains(filledORNot[i], "B4") {
				f("       ", filledORNot[i], "  ")
			} else {
				f("  ", filledORNot[i], "  ")
			}
		}

	}
	p()
	for i := 12; i < 18; i++ {

		if tempBookedStatus[i] == 1 {
			if strings.Contains(filledORNot[i], "C4") {
				f("       --  ")
			} else {
				f("  --  ")
			}

		} else {
			if strings.Contains(filledORNot[i], "C4") {
				f("       ", filledORNot[i], "  ")
			} else {
				f("  ", filledORNot[i], "  ")
			}
		}

	}
	p()
}

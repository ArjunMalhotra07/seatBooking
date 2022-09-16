package main

import (
	"fmt"
)

func bookFlightSeat() {
	p := fmt.Println
	seats := seating()
	eachSeat, tempBookedStatus := seatNAMES(seats)

	p()
	p()
	p()
	var flag bool
	for !flag {
		p("Enter the VACANT Seat ID to Book  --- ")
		var seatidNumber string
		fmt.Scanln(&seatidNumber)

		for i := 0; i < len(eachSeat); i++ {
			if eachSeat[i] == seatidNumber {
				if tempBookedStatus[i] == 0 {
					flag = true
					executeUpdatedSQL(seatidNumber, 1)
					p("Booked")
					updatesSeats := seating()
					seatNAMES(updatesSeats)
				} else {
					p("Seat is already Booked! ")
				}
			}

		}
		if !flag {
			p("VISIT AGAIN ")
		}
		p()
	}

}

package main

import (
	"fmt"
	"strings"
)

func seatNAMES(totalSeats []seatType) ([]string, []int) {
	p := fmt.Println
	var eachSeat []string
	var tempBookedStatus []int
	for i := 0; i < len(totalSeats); i++ {
		eachSeat = append(eachSeat, totalSeats[i].SeatID)
		tempBookedStatus = append(tempBookedStatus, totalSeats[i].BookedStatus)

	}
	length := len(eachSeat)
	p()
	p()
	p("		  SEATS")
	for i := 0; i < length; i++ {
		if i%6 == 0 {
			printLoop(i, i+6, eachSeat, tempBookedStatus)
		}
	}

	return eachSeat, tempBookedStatus
}

func printLoop(m, n int, eachSeat []string, tempBookedStatus []int) {
	f := fmt.Print
	p := fmt.Println
	p()
	for i := m; i < n; i++ {
		if tempBookedStatus[i] == 1 {

			if strings.Contains(eachSeat[i], "4") {
				f("         --  ")
			} else {
				f("  --  ")
			}

		} else {
			if strings.Contains(eachSeat[i], "4") {
				f("         ", eachSeat[i], "  ")
			} else {
				f("  ", eachSeat[i], "  ")
			}
		}

	}
	p()
	p()
}

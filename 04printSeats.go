package main

import "fmt"

func seatNAMES(totalSeats []seatType) ([]string, []int) {
	p := fmt.Println
	var eachSeat []string
	var tempBookedStatus []int
	for i := 0; i < len(totalSeats); i++ {
		eachSeat = append(eachSeat, totalSeats[i].SeatID)
		tempBookedStatus = append(tempBookedStatus, totalSeats[i].BookedStatus)

	}
	p()
	p()
	p("		  SEATS")

	printLoop(0, 6, eachSeat, tempBookedStatus, "A4")
	printLoop(6, 12, eachSeat, tempBookedStatus, "B4")
	printLoop(12, 18, eachSeat, tempBookedStatus, "C4")
	return eachSeat, tempBookedStatus
}

func printLoop(m, n int, eachSeat []string, tempBookedStatus []int, partitionSeat string) {
	f := fmt.Print
	p := fmt.Println
	p()
	for i := m; i < n; i++ {
		if tempBookedStatus[i] == 1 {
			if eachSeat[i] == partitionSeat {
				f("         --  ")
			} else {
				f("  --  ")
			}

		} else {
			if eachSeat[i] == partitionSeat {
				f("         ", eachSeat[i], "  ")
			} else {
				f("  ", eachSeat[i], "  ")
			}
		}

	}
	p()
}

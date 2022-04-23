package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func totalSeats(c *gin.Context) {
	seats := seating()
	c.IndentedJSON(http.StatusOK, seats)

}

func seating() []seatType {
	var seats []seatType
	rows, err := db.Query("SELECT * FROM bookSeat")
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var seat seatType
		if err := rows.Scan(&seat.ID, &seat.SeatID, &seat.BookedStatus); err != nil {
			return nil
		}
		seats = append(seats, seat)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return seats
}

func vacancy() []seatType {
	var vacantSeats []seatType
	rows, err := db.Query("SELECT * FROM bookSeat WHERE bookedStatus = 0")
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var seat seatType
		if err := rows.Scan(&seat.ID, &seat.SeatID, &seat.BookedStatus); err != nil {
			return nil
		}
		vacantSeats = append(vacantSeats, seat)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return vacantSeats
}

// func vacanc() []seatType {
// 	var vacantSeats []seatType
// 	rows, err := db.Query("SELECT * FROM bookSeat WHERE bookedStatus = 0")
// 	if err != nil {
// 		return nil
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var seat seatType
// 		if err := rows.Scan(&seat.ID, &seat.SeatID, &seat.BookedStatus); err != nil {
// 			return nil
// 		}
// 		vacantSeats = append(vacantSeats, seat)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil
// 	}
// 	return vacantSeats
// }

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func bookSeating(c *gin.Context) {
	var seatData seatType

	if err1 := c.BindJSON(&seatData); err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": "Invalid inputs. Please check your inputs"})
		return
	}

	_, err := db.Exec(`UPDATE bookSeat SET seatID = ? , bookedStatus = ? WHERE seatID = ?`,
		seatData.SeatID,
		seatData.BookedStatus,
		seatData.SeatID)

	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(200, "Successfully Altered Data")
}

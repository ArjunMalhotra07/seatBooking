package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type seatType struct {
	ID           int
	SeatID       string
	BookedStatus int
}

func main() {
	f := fmt.Println

	var err error
	db, err = sql.Open("mysql", "aman:Mysql_Witcher7%@tcp(127.0.0.1:3306)/ticket_App")

	// Get a database handle.

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	f("Connected!")
	bookFlightSeat()

	route := gin.Default()
	route.GET("/getSeats", totalSeats)
	route.POST("/getSeats", bookSeating)
	route.Run("localhost:8080")

}

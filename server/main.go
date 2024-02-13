package main

import (
	"fmt"
	"log"
	"net/http"
	"tryUrl/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")

	// Create a new connection to the database
	// db := DBConn.CreateDBConnection()
	// defer db.Close()
	// CreateTicketServer.CreateTable(db)
}

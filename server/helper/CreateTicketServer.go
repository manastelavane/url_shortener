package helper

import (
	"database/sql"
	"fmt"
	"log"
)

// CREATE TABLE your_table_name (
//     id SERIAL PRIMARY KEY,
//     start_ticket BIGINT NOT NULL,
//     end_ticket BIGINT NOT NULL,
//     current BIGINT NOT NULL
// );

func CreateTable(db *sql.DB) {
	// Insert 1000 rows into the table
	start_ticket := 3844
	for i := 0; i < 1000; i++ {
		end_ticket := start_ticket + int(1.2e9) - 1
		current := 0

		_, err := db.Exec("INSERT INTO ticket_server_1 (start_ticket, end_ticket, current) VALUES ($1, $2, $3)", start_ticket, end_ticket, current)
		if err != nil {
			log.Fatal("Error inserting data into table: ", err)
		}

		// Update start for the next iteration
		start_ticket = end_ticket + 1
	}

	fmt.Println("Data inserted successfully!")
}
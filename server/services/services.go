package services

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
	"tryUrl/DBConn"
)

func GetLongURL(shortURL string) (string, error) {
	// Create a new connection to the database
	db := DBConn.CreateDBConnection()
	defer db.Close()

	// Retrieve the long URL from the database based on the short URL
	var longURL string
	err := db.QueryRow("SELECT long_url FROM url_mapping WHERE short_url = $1", shortURL).Scan(&longURL)
	if err != nil {
		return "", err
	}
	return longURL, nil
}

func SelectRandomRange(db *sql.DB) (int64, int64,int64) {
	var start_ticket, end_ticket,current int64
    // Query the database to get the count of ranges
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM ticket_server_1").Scan(&count)
    if err != nil {
        fmt.Println("Error querying count of ranges:", err)
        return start_ticket, end_ticket,current
    }

    // Select a random row
    randomIndex := rand.Intn(count)+1
    err = db.QueryRow("SELECT start_ticket, end_ticket,current FROM ticket_server_1 where id=$1", randomIndex).Scan(&start_ticket, &end_ticket, &current)
    if err != nil {
        fmt.Println("Error selecting random range:", err)
    }

    return start_ticket, end_ticket,current
}

func IncrementCurrent(start_ticket int64, db *sql.DB) error {
	// Increment the current value in the database
	_, err := db.Exec("UPDATE ticket_server_1 SET current = current + 1 WHERE start_ticket = $1", start_ticket)
	if err != nil {
		return err
	}
	return nil
}

func SaveURLMapping(shortURL, longURL string, db *sql.DB) error  {
	// Save the short URL and long URL in the database
	_, err := db.Exec("INSERT INTO url_mapping (short_url, long_url) VALUES ($1, $2)", shortURL, longURL)
	if err != nil {
		fmt.Println("Error saving URL mapping:", err)
		return err;
	}
	return nil;
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

package DBConn

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateDBConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// Print all the rows from the url_mapping table
func SelectQuery(db *sql.DB) {
	// SQL SELECT statement to retrieve rows from the url_mapping table
	selectStmt := `SELECT * FROM url_mapping`

	// Execute the SELECT statement
	rows, err := db.Query(selectStmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the rows returned by the SELECT statement
	for rows.Next() {
		var id int
		var shortURL, longURL string
		err := rows.Scan(&id, &shortURL, &longURL)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Short URL: %s, Long URL: %s\n", id, shortURL, longURL)
	}

	// Check for errors encountered while iterating over the rows
	if err = rows.Err(); err != nil {
		panic(err)
	}
}
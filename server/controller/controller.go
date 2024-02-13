package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tryUrl/DBConn"
	"tryUrl/helper"
	"tryUrl/services"

	"github.com/gorilla/mux"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	// Create a new connection to the database
	db := DBConn.CreateDBConnection()
	defer db.Close()

	// Parse the JSON request body
	var requestData struct {
		LongURL string `json:"long_url"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	// Select a random range from the database
	start_ticket, _ ,current := services.SelectRandomRange(db)

	// Generate the long URL number
	longURLNum := start_ticket + current // Placeholder calculation

	// Increment the current value in the database
	err = services.IncrementCurrent(start_ticket,db)
	if err != nil {
		json.NewEncoder(w).Encode("Error incrementing current value")
		return
	}

	// Call your encoding function to get the short URL
	shortURL := helper.ConvertToAlphabetic(longURLNum)

	// Save the short URL and long URL in the database
	err = services.SaveURLMapping(shortURL, requestData.LongURL,db)
	if err != nil {
		http.Error(w, "Error saving URL mapping", http.StatusInternalServerError)
		return
	}

	// Respond to the user with the short URL
	response := struct {
		LongURL  string `json:"long_url"`
		ShortURL string `json:"short_url"`
	}{
		LongURL:  requestData.LongURL,
		ShortURL: shortURL,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	// Get the short URL from the request parameters
	params := mux.Vars(r)
	fmt.Println(params["shortURL"])
	// Retrieve the long URL from the database based on the short URL
	longURL, err := services.GetLongURL(params["shortURL"])
	if err != nil {
		http.Error(w, "Error retrieving long URL", http.StatusInternalServerError)
		return
	}
	fmt.Println(longURL)
	// Redirect the user to the long URL
	http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
}

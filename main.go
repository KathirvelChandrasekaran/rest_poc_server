package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type RideData struct {
	Key              string `json:"key"`
	FareAmount       string `json:"fare_amount"`
	PickupDatetime   string `json:"pickup_datetime"`
	PickupLongitude  string `json:"pickup_longitude"`
	PickupLatitude   string `json:"pickup_latitude"`
	DropoffLongitude string `json:"dropoff_longitude"`
	DropoffLatitude  string `json:"dropoff_latitude"`
	PassengerCount   string `json:"passenger_count"`
}

// Handler for fetching ride data from the JSON file
func getRideDataHandler(w http.ResponseWriter, r *http.Request) {
	// Open the JSON file
	startTime := time.Now()
	file, err := os.Open("60MB.json")
	if err != nil {
		http.Error(w, "Unable to open JSON file", http.StatusInternalServerError)
		log.Printf("Failed to open JSON file: %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Failed to close JSON file: %v", err)
		}
	}(file)

	// Decode the JSON file
	decoder := json.NewDecoder(file)
	var rides []RideData
	err = decoder.Decode(&rides)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusInternalServerError)
		log.Printf("Failed to decode JSON: %v", err)
		return
	}

	// Send the JSON response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rides)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
		return
	}
	log.Printf("Server: Time taken to process and send data: %v", time.Since(startTime))
}

func main() {
	// Register the API route
	http.HandleFunc("/rides", getRideDataHandler)

	// Start the HTTP server on port 8080
	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}

package main

import (
	"log"
	"net/http"

	"image-processing-service/api"
	"image-processing-service/store_master"
)

func main() {
	// Load store master data
	log.Println("Loading store master data...")
	err := store_master.LoadStoreData("store_master/store_master.csv")
	if err != nil {
		log.Fatalf("Failed to load store master data: %v", err)
	}
	log.Println("Store master data loaded successfully.")

	// Setup router
	router := api.SetupRouter()

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

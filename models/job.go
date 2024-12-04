package models

import "sync"

type JobPayload struct {
	Count  int     `json:"count"`  // Number of visits
	Visits []Visit `json:"visits"` // List of visits with store and image data
}

type Visit struct {
	StoreID   string   `json:"store_id"`   // Store ID where the visit occurred
	ImageURLs []string `json:"image_url"`  // List of image URLs to process
	VisitTime string   `json:"visit_time"` // Time of the visit
}

type ImageResult struct {
	ImageURL  string `json:"image_url"` // Image URL that was processed
	Perimeter int    `json:"perimeter"` // Perimeter of the image
}

// New struct for errors related to a job
type JobError struct {
	StoreID string `json:"store_id"` // Store where the error occurred
	Message string `json:"message"`  // Error message
}

// Updated Job struct with Results field for storing image processing results
type Job struct {
	ID      int           `json:"id"`
	Status  string        `json:"status"`
	Payload JobPayload    `json:"payload"`
	Errors  []JobError    `json:"errors"`
	Results []ImageResult `json:"results"`
	Mu      sync.Mutex    // Mutex for synchronization
}
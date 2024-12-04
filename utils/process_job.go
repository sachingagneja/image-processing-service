package utils

import (
	"image-processing-service/models"
	"log"
)

func ProcessJob(jobID int, jobs map[int]*models.Job, storeData map[string]models.Store) {
    job, exists := jobs[jobID]
    if !exists {
        log.Printf("Job with ID %d not found", jobID)
        return
    }

    var errors []models.JobError

    // Process each visit
    for _, visit := range job.Payload.Visits {
        // Validate store ID
        store, valid := storeData[visit.StoreID]
        if !valid {
            log.Printf("Invalid store ID: %s", visit.StoreID)
            errors = append(errors, models.JobError{
                StoreID: visit.StoreID,
                Message: "Invalid store ID",
            })
            continue
        }

        log.Printf("Processing store: %s (Area Code: %s)", store.StoreName, store.AreaCode)

        // Process each image URL
        for _, imageURL := range visit.ImageURLs {
            log.Printf("Downloading image from %s...", imageURL)

            img, format, err := DownloadImage(imageURL)
            if err != nil {
                log.Printf("Failed to download image from %s: %v", imageURL, err)
                errors = append(errors, models.JobError{
                    StoreID: visit.StoreID,
                    Message: "Failed to download image: " + err.Error(),
                })
                continue
            }

            // Calculate the perimeter
            perimeter := CalculatePerimeter(img)
            log.Printf("Image processed. URL: %s, Perimeter: %d, Format: %s", imageURL, perimeter, format)

            // Append results
            job.Results = append(job.Results, models.ImageResult{
                ImageURL:  imageURL,
                Perimeter: perimeter,
            })

            log.Printf("Starting download for %s", imageURL)
            log.Printf("Image downloaded successfully. Format: %s", format)
            log.Printf("Calculating perimeter...")
            log.Printf("Calculated perimeter: %d", perimeter)

            // Simulate GPU processing
            RandomSleep()
        }
    }

    // Set job status based on errors
    if len(errors) > 0 {
        job.Status = "failed"
        job.Errors = errors
    } else {
        job.Status = "completed"
    }

    // Log the final job status
    log.Printf("Job %d processing completed with status: %s", jobID, job.Status)
}
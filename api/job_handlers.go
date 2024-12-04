package api

import (
	"encoding/json"
	"image-processing-service/models"
	"image-processing-service/store_master"
	"image-processing-service/utils"
	"net/http"
	"strconv"
)

var jobs = make(map[int]*models.Job) // Use int as key

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
    var payload models.JobPayload

    // Decode the request body into JobPayload
    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil || len(payload.Visits) != payload.Count || payload.Count == 0 {
        // If there is an error in decoding, or if the visit count is incorrect, or count is 0
        http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
        return
    }

    jobID := len(jobs) + 1
    jobs[jobID] = &models.Job{ID: jobID, Status: "ongoing", Payload: payload}

    // Convert store_master.StoreData to map[string]models.Store
    storeData := make(map[string]models.Store)
    for storeID, store := range store_master.StoreData {
        storeData[storeID] = utils.ConvertStore(store)
    }

    // Initialize errors slice for this job
    var errors []models.JobError

    // Validate each visit for empty image_url
    for _, visit := range payload.Visits {
        for _, imageURL := range visit.ImageURLs {
            if imageURL == "" {
                errors = append(errors, models.JobError{
                    StoreID: visit.StoreID,
                    Message: "Empty image URL found",
                })
                break // No need to check other image URLs for this visit
            }
        }
    }

    // If there were any errors, fail the job and include them in the response
    if len(errors) > 0 {
        jobs[jobID].Status = "failed"
        jobs[jobID].Errors = errors
    } else {
        // If no errors, start the job processing in a goroutine
        go utils.ProcessJob(jobID, jobs, storeData)
    }

    // Updated JobResponse with Status
    response := models.JobResponse{
        JobID:  jobID,
        Status: "completed",  // Default to "completed"
    }

    // If errors are present, set status to "failed"
    if len(errors) > 0 {
        response.Status = "failed"
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}


func GetJobStatusHandler(w http.ResponseWriter, r *http.Request) {
    jobID := r.URL.Query().Get("jobid")
    if jobID == "" {
        http.Error(w, `{"error": "jobID is required"}`, http.StatusBadRequest)
        return
    }

    // Convert jobID string to int
    jobIDInt, err := strconv.Atoi(jobID)
    if err != nil {
        http.Error(w, `{"error": "Invalid jobID"}`, http.StatusBadRequest)
        return
    }

    job, exists := jobs[jobIDInt]
    if !exists {
        http.Error(w, `{"error": "job not found"}`, http.StatusBadRequest)
        return
    }

    // Set the response header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Send the status of the job in the desired format
    response := map[string]interface{}{
        "status": job.Status,
        "job_id": jobIDInt,
    }

    // If there are errors, send them in the `error` field
    if job.Status == "failed" {
        errorDetails := []map[string]interface{}{}
        for _, jobError := range job.Errors {
            errorDetails = append(errorDetails, map[string]interface{}{
                "store_id": jobError.StoreID,
                "error":    jobError.Message,
            })
        }
        response["error"] = errorDetails
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}


func GetJobResultsHandler(w http.ResponseWriter, r *http.Request) {
    jobID := r.URL.Query().Get("jobid")
    if jobID == "" {
        http.Error(w, `{"error": "jobID is required"}`, http.StatusBadRequest)
        return
    }

    jobIDInt, err := strconv.Atoi(jobID)
    if err != nil {
        http.Error(w, `{"error": "Invalid jobID"}`, http.StatusBadRequest)
        return
    }

    job, exists := jobs[jobIDInt]
    if !exists {
        http.Error(w, `{"error": "Job not found"}`, http.StatusNotFound)
        return
    }

    response := map[string]interface{}{
        "job_id":   jobIDInt,
        "results":  job.Results,
        "status":   job.Status,
        "errors":   job.Errors,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
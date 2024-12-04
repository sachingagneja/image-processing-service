package models

type JobResponse struct {
    JobID  int    `json:"job_id"`
    Status string `json:"status,omitempty"`  // Add Status field here
}


type ErrorResponse struct {
    Error string `json:"error"`
}

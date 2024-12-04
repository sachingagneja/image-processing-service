package api

import "github.com/gorilla/mux"

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/api/submit", SubmitJobHandler).Methods("POST")
    router.HandleFunc("/api/status", GetJobStatusHandler).Methods("GET")
    router.HandleFunc("/api/results", GetJobResultsHandler).Methods("GET") // New endpoint
    return router
}

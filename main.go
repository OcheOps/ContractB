package main

import (
    "log"
    "net/http"

    "github.com/OcheOps/ContractB/handlers"
)

func main() {
    http.HandleFunc("/report", handlers.ReportHandler)
	http.HandleFunc("/project-details", handlers.CreateProjectDetailsHandler)
	http.HandleFunc("/project-progress", handlers.CreateProjectProgressHandler)

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


package main

import (
    "log"
    "net/http"

    "github.com/your-project/handlers"
)

func main() {
    http.HandleFunc("/report", handlers.ReportHandler)

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
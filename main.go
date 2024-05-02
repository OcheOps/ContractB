package main

import (
    "log"
    "net/http"

    "github.com/OcheOps/ContractB/handlers"
    "github.com/rs/cors"
)

func main() {

    corsHandler := cors.New(cors.Options{
        AllowedOrigins:   []string{"https://contract-f.vercel.app"}, // Replace with your allowed origins
        AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
        AllowCredentials: false, // Allow cookies for cross-origin requests (optional)
    })

    wrappedHandler := corsHandler.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/report":
            handlers.ReportHandler(w, r)
        case "/project-details":
            handlers.CreateProjectDetailsHandler(w, r)
        case "/project-progress":
            handlers.CreateProjectProgressHandler(w, r)
        default:
        }
    }))

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", wrappedHandler))

    // http.HandleFunc("/report", handlers.ReportHandler)
	// http.HandleFunc("/project-details", handlers.CreateProjectDetailsHandler)
	// http.HandleFunc("/project-progress", handlers.CreateProjectProgressHandler)

    // log.Println("Server started on :8080")
    // log.Fatal(http.ListenAndServe(":8080", nil))
}


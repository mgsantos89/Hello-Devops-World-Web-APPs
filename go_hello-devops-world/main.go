package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        hostname = "unknown"
    }
    currentTime := time.Now().Format("2006-01-02 15:04:05")
    fmt.Fprintf(w, "<h1>Hello Devops World</h1>")
    fmt.Fprintf(w, "<p>Date and Time: %s</p>", currentTime)
    fmt.Fprintf(w, "<p>Hostname: %s</p>", hostname)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}

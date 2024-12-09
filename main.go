package main

import (
    "log"
    "net/http"

    "groupie-tracker/api"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/artist", api.GetArtistsHandler)
    mux.HandleFunc("/api/location", api.GetLocationsHandler)
    mux.HandleFunc("/api/date", api.GetDatesHandler)
    mux.HandleFunc("/api/relation", api.GetRelationsHandler)

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
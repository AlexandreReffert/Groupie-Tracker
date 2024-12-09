package main

import (
    "log"
    "net/http"

    "groupie-tracker/api"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/artists", api.GetArtistsHandler)
    mux.HandleFunc("/api/locations", api.GetLocationsHandler)
    mux.HandleFunc("/api/dates", api.GetDatesHandler)
    mux.HandleFunc("/api/relations", api.GetRelationsHandler)

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
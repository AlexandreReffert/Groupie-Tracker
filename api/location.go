package api

import (
    "encoding/json"
    "net/http"
    "groupie-tracker/Basededonnées"
)

func GetLocationsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Basededonnées.GetLocations())
}
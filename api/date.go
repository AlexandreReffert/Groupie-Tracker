package api

import (
    "encoding/json"
    "net/http"
    "groupie-tracker/Basededonnées"
)

func GetDatesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Basededonnées.GetDates())
}
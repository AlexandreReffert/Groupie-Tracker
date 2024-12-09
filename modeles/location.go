package models

type Location struct {
    LastLocation   string   `json:"last_location"`
    Upcoming       []string `json:"upcoming"`
}
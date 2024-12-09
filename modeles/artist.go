package models

type Artist struct {
    Name         string   `json:"name"`
    Image        string   `json:"image"`
    StartYear    int      `json:"start_year"`
    FirstAlbum   int      `json:"first_album"`
    Members      []string `json:"members"`
}
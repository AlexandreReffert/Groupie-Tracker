package modeles

type Relation struct {
    ArtistName string `json:"artist_name"` // Nom de l'artiste
    City       string `json:"city"`       // Ville
    Date       string `json:"date"`       // Date du concert
}
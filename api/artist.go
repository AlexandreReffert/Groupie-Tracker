package api

import (
	"encoding/json" // Pour encoder/décoder les données JSON
	"net/http"      // Pour gérer les requêtes HTTP
	"groupie-tracker/Basededonnées"     // Import du package de base de données
	"groupie-tracker/models" // Import des modèles pour les artistes
)

// Fonction pour récupérer tous les artistes
func GetArtistsHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour retourner les artistes (à implémenter)
}

// Fonction pour ajouter un nouvel artiste
func AddArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Logique pour ajouter un nouvel artiste (à implémenter)
}

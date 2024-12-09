package api

import (
	"encoding/json"
	"net/http"
	"groupie-tracker/Basededonnées"
	"groupie-tracker/models"
)

// Endpoint pour récupérer toutes les relations
func GetRelationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.Relations)
}

// Endpoint pour ajouter une nouvelle relation
func AddRelationHandler(w http.ResponseWriter, r *http.Request) {
	var relation models.Relation

	// Décoder le JSON envoyé dans la requête
	if err := json.NewDecoder(r.Body).Decode(&relation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ajouter la relation dans la base de données
	db.AddRelation(relation)

	// Répondre avec succès
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(relation)
}
package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"login/src/models" // Remplace par le chemin correct vers tes modèles
	"net/http"
)

// Supprimer le compte de l'utilisateur connecté
func DeleteAccountHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Vérifier si l'utilisateur est connecté
	session, _ := GetSessionUsername(r)
	username := session

	fmt.Println("fzffzf", session)

	// Supprimer l'utilisateur de la base de données
	err := models.DeleteUserByUsername(db, username)
	if err != nil {
		log.Println("Erreur lors de la suppression du compte :", err)
		http.Error(w, "Impossible de supprimer le compte", http.StatusInternalServerError)
		return
	}

	// Supprimer la session de l'utilisateur
	ClearSession(w, r)

	// Rediriger vers la page de confirmation ou la page d'accueil
	http.Redirect(w, r, "/account-deleted", http.StatusSeeOther)
}

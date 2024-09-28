package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"login/src/models"
)

// Affiche la page de connexion
func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/template/login.html"))
		tmpl.Execute(w, nil)
	}
}

// Gère la soumission du formulaire de connexion
func LoginHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Récupérer l'utilisateur depuis la base de données
	user, err := models.GetUserByUsername(db, username)
	if err != nil || user.Password != password {
		http.Error(w, "Nom d'utilisateur ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Bienvenue, %s!", user.Username)
}

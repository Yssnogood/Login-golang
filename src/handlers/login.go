package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"login/src/models"
)

// Affiche la page de connexion
func LoginPage(w http.ResponseWriter, r *http.Request) {

	if IsLoggedIn(r) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/template/login.html"))
		tmpl.Execute(w, nil)
	}
}

// Gère la soumission du formulaire de connexion
func LoginHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Récupérer l'utilisateur
		user, err := models.GetUserByUsername(db, username)
		if err != nil || models.CheckPassword(user.Password, password) != nil {
			http.Error(w, "Nom d'utilisateur ou mot de passe incorrect", http.StatusUnauthorized)
			return
		}

		// Créer une session si la connexion est valide
		SetSession(w, r, username)

		// Rediriger vers le dashboard après une connexion réussie
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
}

// Déconnexion de l'utilisateur et suppression de la session
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	ClearSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

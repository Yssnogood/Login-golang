package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"login/src/models"
)

// Affiche la page d'inscription
func SignupPage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/template/signup.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Vérifier si l'utilisateur existe déjà
		_, err := models.GetUserByUsername(db, username)
		if err == nil {
			http.Error(w, "Cet utilisateur existe déjà", http.StatusConflict)
			return
		}

		// Ajouter l'utilisateur dans la base de données
		err = models.CreateUser(db, username, password)
		if err != nil {
			http.Error(w, "Erreur lors de la création du compte", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Compte créé avec succès pour %s!", username)
	}
}

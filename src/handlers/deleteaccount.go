package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"login/src/models"
)

// Page du dashboard avec le nom d'utilisateur
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	if !IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	username, exists := GetSessionUsername(r)
	if !exists {
		http.Redirect(w, r, "/", http.StatusSeeOther) // If no session, redirect to login
		return
	}

	tmpl := template.Must(template.ParseFiles("./web/template/deleteaccount.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Username": username,
	})
}

func DeleteAccountHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if !IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	username, exists := GetSessionUsername(r)
	if !exists {
		http.Redirect(w, r, "/", http.StatusSeeOther) // If no session, redirect to login
		return
	}

	tmpl := template.Must(template.ParseFiles("./web/template/deleteaccount.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Username": username,
	})

	// Retrieve the password from the form
	password := r.FormValue("password1")
	passwordConfirm := r.FormValue("password-confirm")

	// Validate the password (you should implement this function)
	if password != passwordConfirm {
		http.Error(w, "Les mots de passe ne correspondent pas.", http.StatusBadRequest)
		return
	}
	user, _ := models.GetUserByUsername(db, username)

	if models.CheckPassword(user.Password, passwordConfirm) != nil {
		http.Error(w, "Mot de passe incorrect.", http.StatusBadRequest)
		return

	}

	// Proceed to delete the account (implement your own logic)
	models.DeleteUser(db, username) // Implement this function

	// Clear the session if account deletion is successful
	ClearSession(w, r)

	// Redirect to a confirmation page or the login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

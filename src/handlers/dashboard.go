package handlers

import (
	"html/template"
	"net/http"
)

// Page du dashboard avec le nom d'utilisateur
func DashboardPage(w http.ResponseWriter, r *http.Request) {
	if !IsLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	session, _ := store.Get(r, "session-name")
	username := session.Values["username"]

	tmpl := template.Must(template.ParseFiles("./web/template/dashboard.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Username": username,
	})
}

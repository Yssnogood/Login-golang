package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

// Obtenir la session actuelle
func GetSession(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "session-name")
	return session
}

// Créer une session pour l'utilisateur
func SetSession(w http.ResponseWriter, r *http.Request, username string) {
	session := GetSession(r)
	session.Values["username"] = username
	session.Save(r, w)
}

// Effacer la session lors de la déconnexion
func ClearSession(w http.ResponseWriter, r *http.Request) {
	session := GetSession(r)
	delete(session.Values, "username")
	session.Save(r, w)
}

// Vérifier si l'utilisateur est connecté
func IsLoggedIn(r *http.Request) bool {
	session := GetSession(r)
	_, ok := session.Values["username"]
	return ok
}

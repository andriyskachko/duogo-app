package main

import (
	"fmt"
	"log"
	"net/http"
)

func UserServer(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodPost:
		HandlePOSTUser(w, r)
	default:
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
	}
}

func HandlePOSTUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, emailErr := IsValidEmail(email)
	log.Println(emailErr)
	if emailErr != nil {
		writeError(w, emailErr)
		return
	}

	_, passwordErr := IsValidPassword(password)
	log.Println(passwordErr)
	if passwordErr != nil {
		writeError(w, passwordErr)
		return
	}

	// TODO: refactor set cookies to generate token
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "session_token",
	// 	Value:    "123",
	// 	Expires:  time.Now(),
	// 	HttpOnly: true,
	// })

	http.Redirect(w, r, "/home", http.StatusFound)
}

func writeError(w http.ResponseWriter, err error) {
	errorMessage := err.Error()
	w.Header().Set("Content-Type", ContentTypeHTML)
	http.Error(w, errorMessage, http.StatusUnprocessableEntity)
	fmt.Fprintf(w, "<div class=\"error\">%s</div>", errorMessage)
}

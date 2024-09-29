package server

import (
	"log"
	"net/http"

	"github.com/8ideaz/manna/views"
	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func indexHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request for %s", "/")
		v.Render(w, nil)
	}
}

func contactHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request for %s", "/contact")
		v.Render(w, nil)
	}
}

func aboutHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request for %s", "/about")
		v.Render(w, nil)
	}
}

func showLoginPage(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request for %s", "/login")
		v.Render(w, nil)
	}
}

func showBibleHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request for %s", "/bible")
		v.Render(w, nil)
	}
}

func signInWithProvider() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")
		// provider := c.Param("provider")
		q := r.URL.Query()
		// q := c.Request.URL.Query()
		q.Add("provider", provider)
		r.URL.RawQuery = q.Encode()

		gothic.BeginAuthHandler(w, r)
	}
}

func callbackHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")
		q := r.URL.Query()
		q.Add("provider", provider)
		r.URL.RawQuery = q.Encode()

		_, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin", http.StatusTemporaryRedirect)
		// c.Redirect(http.StatusTemporaryRedirect, "/success")
	}
}

func logoutHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gothic.Logout(w, r)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
		v.Render(w, nil)
	}
}
func adminHomeHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request for %s", "/success")
		provider := chi.URLParam(r, "provider")
		q := r.URL.Query()
		q.Add("provider", provider)
		r.URL.RawQuery = q.Encode()

		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// w.Header().Set("Location", "/admin")
		// w.WriteHeader(http.StatusTemporaryRedirect)
		v.Render(w, user)
	}
}

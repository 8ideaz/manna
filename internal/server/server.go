package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/8ideaz/manna/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

// var index *views.View
// var contact *views.View

func Run() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	port := os.Getenv("MANNA_PORT")
	if port == "" {
		port = "3000"
	}
	addr := fmt.Sprintf(":%s", port)
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	clientCallbackURL := os.Getenv("CLIENT_CALLBACK_URL")
	if clientID == "" || clientSecret == "" || clientCallbackURL == "" {
		err = fmt.Errorf("error: Environment variables (CLIENT_ID, CLIENT_SECRET, CLIENT_CALLBACK_URL) are required")
		return err
	}
	goth.UseProviders(
		google.New(clientID, clientSecret, clientCallbackURL),
	)

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	// index := views.NewView("bootstrap", "web/layouts", "web/pages/index.html")
	index := views.NewView(&views.ViewOpts{
		Layout:      "bootstrap",
		LayoutDir:   "web/layouts",
		PartialsDir: "web/partials",
		Files:       []string{"web/pages/index.html"},
	})
	// about := views.NewView("default_layout", "web/layouts", "web/pages/about.html")
	about := views.NewView(&views.ViewOpts{
		Layout:      "default_layout",
		LayoutDir:   "web/layouts",
		PartialsDir: "web/partials",
		Files:       []string{"web/pages/about.html"},
	})
	contact := views.NewView(&views.ViewOpts{
		Layout:      "default_layout",
		LayoutDir:   "web/layouts",
		PartialsDir: "web/partials",
		Files:       []string{"web/pages/contact.html"},
	})
	login := views.NewView(&views.ViewOpts{
		Layout:      "default_layout",
		LayoutDir:   "web/layouts",
		PartialsDir: "web/partials",
		Files:       []string{"web/pages/login.html"},
	})
	admin := views.NewView(&views.ViewOpts{
		Layout:      "admin_layout",
		LayoutDir:   "web/layouts",
		PartialsDir: "web/partials",
		Files:       []string{"web/pages/admin_home.html"},
	})
	logout := views.NewView(&views.ViewOpts{
		Layout:      "default_layout",
		LayoutDir:   "web/layouts",
		PartialsDir: "web/partials",
		Files:       []string{"web/pages/index.html"},
	})

	r.Get("/", indexHandler(index))
	r.Get("/about", aboutHandler(about))
	r.Get("/contact", contactHandler(contact))
	r.Get("/auth/{provider}", signInWithProvider())
	r.Get("/auth/{provider}/callback", adminHomeHandler(admin))
	r.Get("/admin", adminHomeHandler(admin))
	r.Get("/login", showLoginPage(login))
	r.Get("/logout/{provider}", logoutHandler(logout))
	log.Printf("Listening on %s", addr)
	return http.ListenAndServe(addr, r)
}

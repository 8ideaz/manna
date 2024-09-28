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
)

// var index *views.View
// var contact *views.View

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

	r.Get("/", indexHandler(index))
	r.Get("/about", aboutHandler(about))
	r.Get("/contact", contactHandler(contact))
	log.Printf("Listening on %s", addr)
	return http.ListenAndServe(addr, r)
}

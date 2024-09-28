package server

import (
	"net/http"

	"github.com/8ideaz/manna/views"
)

// var index *views.View
// var contact *views.View

func indexHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v.Render(w, nil)
	}
}

func contactHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v.Render(w, nil)
	}
}

func aboutHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v.Render(w, nil)
	}
}

func Run() error {
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

	http.HandleFunc("/", indexHandler(index))
	http.HandleFunc("/about", aboutHandler(about))
	http.HandleFunc("/contact", contactHandler(contact))
	return http.ListenAndServe(":3000", nil)
}

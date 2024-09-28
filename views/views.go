package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type ViewOpts struct {
	Layout      string
	LayoutDir   string
	PartialsDir string
	Files       []string
}

// func NewView(layout string, partialDir string, layoutDir string, files ...string) *View {
func NewView(opts *ViewOpts) *View {
	files := append(layoutFiles(opts.LayoutDir), opts.Files...)
	files = append(files, partialFiles(opts.PartialsDir)...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   opts.Layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func layoutFiles(dir string) []string {
	filesPath := filepath.Join(dir, "*.html")
	files, err := filepath.Glob(filesPath)
	if err != nil {
		panic(err)
	}
	return files
}

func partialFiles(dir string) []string {
	filesPath := filepath.Join(dir, "*.html")
	files, err := filepath.Glob(filesPath)
	if err != nil {
		panic(err)
	}
	return files
}

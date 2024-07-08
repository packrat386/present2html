package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"os"

	"golang.org/x/tools/present"
)

//go:embed templates static
var embeddedFS embed.FS

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Lshortfile)

	tmpl := present.Template()

	_, err := tmpl.ParseFS(embeddedFS, "templates/*.tmpl")
	if err != nil {
		log.Fatalf("cannot parse embedded templates: %s", err)
	}

	jsdata, err := fs.ReadFile(embeddedFS, "static/slides.js")
	if err != nil {
		log.Fatalf("cannot read embedded file 'static/slides.js': %s", err)
	}

	cssdata, err := fs.ReadFile(embeddedFS, "static/styles.css")
	if err != nil {
		log.Fatalf("cannot read embedded file 'static/styles.css': %s", err)
	}

	if len(os.Args) != 2 {
		log.Fatalf("wrong number of arguments: %d\nusage: %s input.slide > output.html", len(os.Args), os.Args[0])
	}

	fname := os.Args[1]
	fdata, err := os.Open(fname)
	if err != nil {
		log.Fatalf("cannot open file '%s': %s", fname, err)
	}

	pctx := present.Context{
		ReadFile: os.ReadFile,
	}

	doc, err := pctx.Parse(fdata, fname, 0)
	if err != nil {
		log.Fatalf("cannot parse file '%s': %s'", fname, err)
	}

	data := struct {
		*present.Doc
		Template     *template.Template
		PlayEnabled  bool
		NotesEnabled bool
		SlidesCSS    template.CSS
		SlidesJS     template.JS
	}{
		doc,
		tmpl,
		false,
		false,
		template.CSS(string(cssdata)),
		template.JS(string(jsdata)),
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "root", data)
	if err != nil {
		log.Fatalf("cannot render templated data: %s", err)
	}
}

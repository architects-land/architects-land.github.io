package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

//go:embed dist/.vite/manifest.json
var manifest []byte

var dev bool

var templates = []string{
	"resources/base.gohtml",
	"resources/footer.gohtml",
	"resources/navbar.gohtml",
}

type ViteManifestEntry struct {
	File    string   `json:"file"`
	Name    string   `json:"name"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	CSS     []string `json:"css"`
}

type TemplateData struct {
	Resources struct {
		JS  []string
		CSS []string
	}
	Title     string
	Dev       bool
	HasFooter bool
	HasNav    bool
	Data      interface{}
}

func init() {
	flag.BoolVar(&dev, "dev", false, "development mode")
}

func main() {
	r := mux.NewRouter()
	var v map[string]ViteManifestEntry
	err := json.Unmarshal(manifest, &v)
	if err != nil {
		panic(err)
	}

	tmpl := make(map[string]*template.Template)
	tmpl["index"] = template.Must(template.ParseFiles(
		"templates/index.gohtml",
		"templates/footer.gohtml",
		"templates/navbar.gohtml",
		"templates/base.gohtml",
	))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl["index"].ExecuteTemplate(w, "base", TemplateData{
			Resources: struct {
				JS  []string
				CSS []string
			}{},
			Title:     "Architects Land",
			Dev:       dev,
			HasFooter: true,
			HasNav:    true,
			Data:      nil,
		})
		if err != nil {
			panic(err)
		}
	})
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

package main

import (
	"embed"
	_ "embed"
	"github.com/anhgelus/golatt"
	"github.com/gorilla/mux"
)

type TemplateData struct {
	Title     string
	HasFooter bool
	HasNav    bool
	SEO       SEOData
	Data      interface{}
}

type HeroData struct {
	Title       string
	Description string
	Image       string
	Dark        bool
	Min         bool
}

type SeasonData struct {
	Left        bool
	Title       string
	Description []string
	Image       string
	ImageAlt    string
	Href        string
}

type PersonData struct {
	Name        string
	Image       string
	Description string
	Link        string
}

type SEOData struct {
	Title       string
	URL         string
	Image       string
	Description string
	Domain      string
}

//go:embed templates
var templates embed.FS

var g *golatt.Golatt

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome)
	r.HandleFunc("/rules", handleRules)
	r.NotFoundHandler = &NotFound{}
	r.HandleFunc("/team", handleTeam)
	r.HandleFunc("/season/{id:[a-z-]+}", handleSeason)
	r.HandleFunc("/season/{id:[a-z-]+}/player/{player}", handlePlayer)

	g = golatt.New(templates)
	//g.NotFoundHandler = &NotFound{}
	g.DefaultSeoData = &golatt.SeoData{
		Image:       "",
		Description: "",
		Domain:      "architects-land.anhgelus.world",
	}
	g.FormatTitle = func(t string) string {
		if t == "Architects Land" {
			return t
		}
		return t + " - Architects Land"
	}
	g.Templates = append(g.Templates,
		"templates/organisms/*.gohtml",
		"templates/molecules/*.gohtml",
		"templates/atoms/*.gohtml",
		"templates/base/*.gohtml",
	)

	g.HandleFunc("/", handleHome)

	g.StartServer(":8000")
}

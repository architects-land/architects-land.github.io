package main

import (
	"embed"
	_ "embed"
	"github.com/anhgelus/golatt"
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
	g = golatt.New(templates)
	g.NotFoundHandler = &NotFound{}
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

	rules := golatt.Template{
		Golatt:      g,
		Name:        "rules",
		Title:       "Règles",
		Image:       "purgatory.webp",
		Description: "Les règles d'Architects Land",
		URL:         "/rules",
		Data: struct {
			HasFooter bool
			HasNav    bool
			Hero      *HeroData
		}{
			HasFooter: true,
			HasNav:    true,
			Hero: &HeroData{
				Title:       "Règles",
				Description: "",
				Image:       "purgatory.webp",
				Dark:        false,
				Min:         true,
			},
		},
	}
	teamPage := golatt.Template{
		Golatt:      g,
		Name:        "team",
		Title:       "Équipe",
		Image:       "village-night.webp",
		Description: "L'équipe derrière d'Architects Land",
		URL:         "/team",
		Data: struct {
			HasFooter bool
			HasNav    bool
			Hero      *HeroData
			Team      []*PersonData
		}{
			HasFooter: true,
			HasNav:    true,
			Hero: &HeroData{
				Title:       "Équipe",
				Description: "",
				Image:       "village-night.webp",
				Dark:        false,
				Min:         true,
			},
			Team: team,
		},
	}

	g.HandleFunc("/", handleHome)
	g.HandleFunc("/rules", rules.Handle())
	g.HandleFunc("/team", teamPage.Handle())
	g.HandleFunc("/season/{id:[a-z-]+}", handleSeason)
	g.HandleFunc("/season/{id:[a-z-]+}/player/{player}", handlePlayer)

	g.StartServer(":80")
}

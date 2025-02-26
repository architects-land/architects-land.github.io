package main

import (
	"embed"
	"flag"
	"github.com/anhgelus/golatt"
)

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

type CommonData struct {
	HasFooter bool
	HasNav    bool
	Hero      *HeroData
	Team      []*PersonData
}

//go:embed templates
var templates embed.FS

var g *golatt.Golatt

var dev bool

func init() {
	flag.BoolVar(&dev, "dev", false, "development mode")
}

func main() {
	flag.Parse()
	g = golatt.New(templates)
	g.NotFoundHandler = handleNotFound
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
	g.NewTemplate("rules",
		"/rules",
		"Règles",
		"purgatory.webp",
		"Les Règles d'Architects Land",
		CommonData{
			HasFooter: true,
			HasNav:    true,
			Hero: &HeroData{
				Title:       "Règles",
				Description: "",
				Image:       "purgatory.webp",
				Dark:        false,
				Min:         true,
			},
		}).
		Handle()
	g.NewTemplate("team",
		"/team",
		"Équipe",
		"supernoob-field.jpg",
		"L'équipe derrière Architects Land",
		CommonData{
			HasFooter: true,
			HasNav:    true,
			Hero: &HeroData{
				Title:       "Équipe",
				Description: "",
				Image:       "supernoob-field.jpg",
				Dark:        true,
				Min:         true,
			},
			Team: team,
		}).
		Handle()
	g.HandleFunc("/season/{id:[a-z-]+}", handleSeason)
	g.HandleFunc("/season/{id:[a-z-]+}/player/{player}", handlePlayer)

	if dev {
		g.StartServer(":8000")
	} else {
		g.StartServer(":80")
	}
}

package main

import (
	_ "embed"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"log/slog"
	"net/http"
	"time"
)

var dev bool

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

func init() {
	flag.BoolVar(&dev, "dev", false, "development mode")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome)
	r.HandleFunc("/rules", handleRules)
	r.NotFoundHandler = &NotFound{}
	r.HandleFunc("/team", handleTeam)
	r.HandleFunc("/season/{id:[a-z-]+}", handleSeason)
	r.HandleFunc("/season/{id:[a-z-]+}/player/{player}", handlePlayer)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./dist/assets"))))

	srv := &http.Server{
		Handler:      r,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	slog.Info("Starting...")
	log.Fatal(srv.ListenAndServe())
}

package main

import (
	_ "embed"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

//go:embed dist/.vite/manifest.json
var manifest []byte

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
}

type SeasonData struct {
	Left        bool
	Title       string
	Description []string
	Image       string
	ImageAlt    string
}

func init() {
	flag.BoolVar(&dev, "dev", false, "development mode")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

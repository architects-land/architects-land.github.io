package main

import (
	"github.com/anhgelus/golatt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleHome(w http.ResponseWriter, _ *http.Request) {
	var seasonsData []*SeasonData
	i := 0
	for _, v := range seasons {
		seasonsData = append(seasonsData, &SeasonData{
			Left:        i%2 == 0,
			Title:       v.Name,
			Description: v.Information.Description,
			Image:       v.Logo,
			ImageAlt:    "Logo de " + v.Name,
			Href:        "/season/" + v.ID,
		})
		i++
	}

	img := "ruine-des-civilisations/background.jpg"

	g.Render(w, "index", &golatt.TemplateData{
		Title: "Architects Land",
		SEO: &golatt.SeoData{
			URL:         "",
			Image:       img,
			Description: "Famille de SMP Minecraft privé",
		},
		Data: struct {
			HasFooter bool
			HasNav    bool
			Hero      *HeroData
			Seasons   []*SeasonData
		}{
			HasFooter: true,
			HasNav:    true,
			Hero: &HeroData{
				Title:       "Architects Land",
				Description: "Famille de SMP Minecraft privé",
				Image:       img,
				Dark:        false,
				Min:         false,
			},
			Seasons: seasonsData,
		},
	})
}

func handleSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		handleNotFound(w, r)
		return
	}
	season, ok := GetSeason(id)
	if !ok {
		handleNotFound(w, r)
		return
	}

	g.Render(w, "season/index", &golatt.TemplateData{
		Title: season.Name,
		SEO: &golatt.SeoData{
			URL:         "/season/" + season.ID,
			Image:       season.Image,
			Description: season.Description,
		},
		Data: struct {
			HasFooter bool
			HasNav    bool
			Hero      *HeroData
			Season    *FullSeasonData
		}{
			HasFooter: true,
			HasNav:    true,
			Hero: &HeroData{
				Title:       season.Name,
				Description: season.Description,
				Image:       season.Image,
				Dark:        false,
				Min:         true,
			},
			Season: season.toFullSeasonData(),
		},
	})
}

func handlePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		handleNotFound(w, r)
		return
	}
	season, ok := GetSeason(id)
	if !ok {
		handleNotFound(w, r)
		return
	}
	pseudo, ok := vars["player"]
	if !ok {
		handleNotFound(w, r)
		return
	}
	var player *SeasonPlayer
	for _, p := range season.Players {
		if p.Pseudo == pseudo {
			player = p
		}
	}
	if player == nil {
		handleNotFound(w, r)
		return
	}

	g.Render(w, "season/player", &golatt.TemplateData{
		Title: player.Name + " - " + season.Name,
		SEO: &golatt.SeoData{
			URL:         "/season/" + season.Name + "/" + player.Pseudo,
			Image:       "skins/" + player.Pseudo + ".png",
			Description: player.Description,
		},
		Data: struct {
			HasFooter bool
			HasNav    bool
			Season    *Season
			Player    *SeasonPlayer
		}{
			HasFooter: false,
			HasNav:    false,
			Season:    season,
			Player:    player,
		},
	})
}

func handleNotFound(w http.ResponseWriter, _ *http.Request) {
	g.Render(w, "lost", &golatt.TemplateData{
		Title: "404",
		SEO: &golatt.SeoData{
			URL:         "",
			Image:       "nether.jpg",
			Description: "Il semblerait que vous vous êtes perdu·es dans le nether. (Erreur 404)",
		},
		Data: struct {
			HasFooter bool
			HasNav    bool
			Hero      *HeroData
		}{
			HasFooter: false,
			HasNav:    false,
			Hero: &HeroData{
				Title:       "Perdu ?",
				Description: "Il semblerait que vous vous êtes perdu·es dans le nether. Vous allez être redirigés dans l'overworld.",
				Image:       "nether.jpg",
				Dark:        false,
				Min:         false,
			},
		},
	})
}

package main

import (
	"fmt"
	"github.com/anhgelus/golatt"
	"github.com/gorilla/mux"
	"image"
	"image/draw"
	"image/png"
	"net/http"
)

func handleHome(w http.ResponseWriter, _ *http.Request) {
	var seasonsData []*SeasonData
	for i, v := range seasons {
		seasonsData = append(seasonsData, &SeasonData{
			Left:        i%2 == 0,
			Title:       v.Name,
			Description: v.Information.Description,
			Image:       v.Logo,
			ImageAlt:    "Logo de " + v.Name,
			Href:        "/season/" + string(v.ID),
		})
	}

	img := "ruine-des-civilisations/background.jpg"

	g.Render(w, "index", &golatt.TemplateData{
		Title: "Architects Land",
		SEO: &golatt.SeoData{
			URL:         "/",
			Image:       golatt.GetStaticPath(img),
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
	season, ok := GetSeason(SeasonID(id))
	if !ok {
		handleNotFound(w, r)
		return
	}

	g.Render(w, "season/index", &golatt.TemplateData{
		Title: season.Name,
		SEO: &golatt.SeoData{
			URL:         "/season/" + string(season.ID),
			Image:       golatt.GetStaticPath(season.Image),
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
	season, player, ok := getInfoFromURI(r)
	if !ok {
		handleNotFound(w, r)
		return
	}

	img := GetSkin(season.ID, player)

	g.Render(w, "season/player", &golatt.TemplateData{
		Title: player.Name + " - " + season.Name,
		SEO: &golatt.SeoData{
			URL:         "/season/" + season.Name + "/" + player.Pseudo,
			Image:       img,
			Description: player.Description,
		},
		Data: struct {
			HasFooter bool
			HasNav    bool
			Season    *Season
			Player    *SeasonPlayer
			Image     string
		}{
			HasFooter: false,
			HasNav:    false,
			Season:    season,
			Player:    player,
			Image:     img,
		},
	})
}

func handleSkin(w http.ResponseWriter, r *http.Request) {
	season, player, ok := getInfoFromURI(r)
	if !ok {
		handleNotFound(w, r)
		return
	}
	f, err := g.StaticFS.Open(fmt.Sprintf("%s/skins/%s.png", season.ID, player.Pseudo))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	output := image.NewRGBA(img.Bounds())
	draw.Draw(output, output.Bounds(), img, image.Pt(0, 0), draw.Src)
	err = png.Encode(w, output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getInfoFromURI(r *http.Request) (*Season, *SeasonPlayer, bool) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, nil, false
	}
	season, ok := GetSeason(SeasonID(id))
	if !ok {
		return nil, nil, false
	}
	pseudo, ok := vars["player"]
	if !ok {
		return nil, nil, false
	}
	var player *SeasonPlayer
	for _, p := range season.Players {
		if p.Pseudo == pseudo {
			player = p
		}
	}
	if player == nil {
		return nil, nil, false
	}
	return season, player, true
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

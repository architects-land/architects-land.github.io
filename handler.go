package main

import (
	"github.com/anhgelus/golatt"
	"github.com/gorilla/mux"
	"html/template"
	"log/slog"
	"net/http"
)

var components = []string{
	"src/organisms/footer.gohtml",
	"src/organisms/navbar.gohtml",
	"src/molecules/hero.gohtml",
	"src/molecules/season.gohtml",
	"src/molecules/person.gohtml",
	"src/atoms/button.gohtml",
	"src/base/base.gohtml",
	"src/base/opengraph.gohtml",
}

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

	g.Render(w, "index", &golatt.TemplateData{
		Title: "Architects Land",
		SEO: &golatt.SeoData{
			Title:       "Architects Land",
			URL:         "/",
			Image:       "terre-des-civilisations/background.webp",
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
				Image:       "terre-des-civilisations/background.webp",
				Dark:        false,
				Min:         false,
			},
			Seasons: seasonsData,
		},
	})
}

func handleTeam(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "team", &TemplateData{
		Title:     "Équipe - Architects Land",
		HasFooter: true,
		HasNav:    true,
		SEO: SEOData{
			Title:       "Équipe - Architects Land",
			URL:         "team",
			Image:       "village-night.webp",
			Description: "L'équipe derrière Architects Land",
		},
		Data: struct {
			Hero *HeroData
			Team []*PersonData
		}{
			Hero: &HeroData{
				Title:       "Équipe",
				Description: "",
				Image:       "village-night.webp",
				Dark:        false,
				Min:         true,
			},
			Team: team,
		},
	})
}

func handleSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		(&NotFound{}).ServeHTTP(w, r)
		return
	}
	season, ok := GetSeason(id)
	if !ok {
		(&NotFound{}).ServeHTTP(w, r)
		return
	}

	executeTemplate(w, "season/index", &TemplateData{
		Title:     season.Name + " - Architects Land",
		HasFooter: true,
		HasNav:    true,
		SEO: SEOData{
			Title:       season.Name + " - Architects Land",
			URL:         "season/" + season.ID,
			Image:       season.Image,
			Description: season.Description,
		},
		Data: struct {
			Hero   *HeroData
			Season *FullSeasonData
		}{
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
		(&NotFound{}).ServeHTTP(w, r)
		return
	}
	season, ok := GetSeason(id)
	if !ok {
		(&NotFound{}).ServeHTTP(w, r)
		return
	}
	pseudo, ok := vars["player"]
	if !ok {
		(&NotFound{}).ServeHTTP(w, r)
		return
	}
	var player *SeasonPlayer
	for _, p := range season.Players {
		if p.Pseudo == pseudo {
			player = p
		}
	}
	if player == nil {
		(&NotFound{}).ServeHTTP(w, r)
		return
	}

	executeTemplate(w, "season/player", &TemplateData{
		Title:     player.Name + " - " + season.Name + " - Architects Land",
		HasFooter: false,
		HasNav:    false,
		SEO: SEOData{
			Title:       player.Name + " - " + season.Name + " - Architects Land",
			URL:         "season/" + season.Name + "/" + player.Pseudo,
			Image:       "skins/" + player.Pseudo + ".png",
			Description: player.Description,
		},
		Data: struct {
			Season *Season
			Player *SeasonPlayer
		}{
			Season: season,
			Player: player,
		},
	})
}

type NotFound struct{}

func (nf *NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "lost", &TemplateData{
		Title:     "404 - Architects Land",
		HasFooter: false,
		HasNav:    false,
		SEO: SEOData{
			Title:       "404 - Architects Land",
			URL:         "",
			Image:       "nether.webp",
			Description: "Il semblerait que vous vous êtes perdu·es dans le nether. (Erreur 404)",
		},
		Data: struct {
			Hero *HeroData
		}{
			Hero: &HeroData{
				Title:       "Perdu ?",
				Description: "Il semblerait que vous vous êtes perdu·es dans le nether. Vous allez être redirigés dans l'overworld.",
				Image:       "nether.webp",
				Dark:        false,
				Min:         false,
			},
		},
	})
}

func executeTemplate(w http.ResponseWriter, page string, data *TemplateData) {
	data.SEO.Domain = "architects-land.anhgelus.world"
	slog.Info("Loading page", "page", page)
	err := template.Must(template.ParseFiles(
		append(components, "src/page/"+page+".gohtml")...,
	)).ExecuteTemplate(w, "base", data)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

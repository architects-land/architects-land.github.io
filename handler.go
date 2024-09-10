package main

import (
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
	"src/templates/base.gohtml",
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var seasonsData []*SeasonData
	i := 0
	for k, v := range seasons {
		seasonsData = append(seasonsData, &SeasonData{
			Left:        i%2 == 0,
			Title:       v.Name,
			Description: v.Information.Description,
			Image:       v.Logo,
			ImageAlt:    "Logo de " + v.Name,
			Href:        "/season/" + k,
		})
		i++
	}

	executeTemplate(w, "index", TemplateData{
		Resources: struct {
			JS  []string
			CSS []string
		}{},
		Title:     "Architects Land",
		Dev:       dev,
		HasFooter: true,
		HasNav:    true,
		Data: struct {
			Hero    *HeroData
			Seasons []*SeasonData
		}{
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

func handleRules(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "rules", TemplateData{
		Resources: struct {
			JS  []string
			CSS []string
		}{},
		Title:     "Règles - Architects Land",
		Dev:       dev,
		HasFooter: true,
		HasNav:    true,
		Data: struct {
			Hero *HeroData
		}{
			Hero: &HeroData{
				Title:       "Règles",
				Description: "",
				Image:       "purgatory.webp",
				Dark:        false,
				Min:         true,
			},
		},
	})
}

func handleTeam(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "team", TemplateData{
		Resources: struct {
			JS  []string
			CSS []string
		}{},
		Title:     "Équipe - Architects Land",
		Dev:       dev,
		HasFooter: true,
		HasNav:    true,
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
	season, ok := seasons[id]
	if !ok {
		(&NotFound{}).ServeHTTP(w, r)
		return
	}

	executeTemplate(w, "season/index", TemplateData{
		Resources: struct {
			JS  []string
			CSS []string
		}{},
		Title:     season.Name + " - Architects Land",
		Dev:       dev,
		HasFooter: true,
		HasNav:    true,
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
	season, ok := seasons[id]
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

	executeTemplate(w, "season/player", TemplateData{
		Resources: struct {
			JS  []string
			CSS []string
		}{},
		Title:     player.Name + " - " + season.Name + " - Architects Land",
		Dev:       dev,
		HasFooter: true,
		HasNav:    true,
		Data: struct {
			Season *Season
			Player *SeasonPlayer
		}{
			Season: &season,
			Player: player,
		},
	})
}

type NotFound struct{}

func (nf *NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "lost", TemplateData{
		Resources: struct {
			JS  []string
			CSS []string
		}{},
		Title:     "404 - Architects Land",
		Dev:       dev,
		HasFooter: true,
		HasNav:    true,
		Data: struct {
			Hero *HeroData
		}{
			Hero: &HeroData{
				Title:       "Perdu ?",
				Description: "Il semblerait que vous vous êtes perdu·es dans le nether. Vous allez être redirigés dans l'overworld.",
				Image:       "nether.png",
				Dark:        false,
				Min:         false,
			},
		},
	})
}

func executeTemplate(w http.ResponseWriter, page string, data TemplateData) {
	slog.Info("Loading page", "page", page)
	err := template.Must(template.ParseFiles(
		append(components, "src/pages/"+page+".gohtml")...,
	)).ExecuteTemplate(w, "base", data)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

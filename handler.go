package main

import (
	"html/template"
	"net/http"
)

var components = []string{
	"src/organisms/footer.gohtml",
	"src/organisms/navbar.gohtml",
	"src/molecules/hero.gohtml",
	"src/molecules/season.gohtml",
	"src/molecules/person.gohtml",
	"src/templates/base.gohtml",
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var seasonsData []SeasonData
	i := 0
	for k, v := range seasons {
		seasonsData = append(seasonsData, SeasonData{
			Left:        i%2 == 0,
			Title:       v.Name,
			Description: v.Information.Description,
			Image:       v.Image,
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
			Hero    HeroData
			Seasons []SeasonData
		}{
			Hero: HeroData{
				Title:       "Architects Land",
				Description: "Famille de SMP Minecraft privé",
				Image:       "/static/terre-des-civilisations/background.webp",
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
			Hero HeroData
		}{
			Hero: HeroData{
				Title:       "Règles",
				Description: "",
				Image:       "/static/purgatory.webp",
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
			Hero HeroData
			Team []PersonData
		}{
			Hero: HeroData{
				Title:       "Règles",
				Description: "",
				Image:       "/static/purgatory.webp",
			},
			Team: team,
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
			Hero HeroData
		}{
			Hero: HeroData{
				Title:       "Perdu ?",
				Description: "Il semblerait que vous vous êtes perdu·es dans le nether. Vous allez être redirigés dans l'overworld.",
				Image:       "/static/nether.png",
			},
		},
	})
}

func executeTemplate(w http.ResponseWriter, page string, data TemplateData) {
	err := template.Must(template.ParseFiles(
		append(components, "src/pages/"+page+".gohtml")...,
	)).ExecuteTemplate(w, "base", data)
	if err != nil {
		panic(err)
	}
}

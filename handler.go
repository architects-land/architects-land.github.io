package main

import (
	"html/template"
	"net/http"
)

var tmpl map[string]*template.Template

var components = []string{
	"src/organisms/footer.gohtml",
	"src/organisms/navbar.gohtml",
	"src/molecules/hero.gohtml",
	"src/molecules/season.gohtml",
	"src/templates/base.gohtml",
}

func init() {
	tmpl = make(map[string]*template.Template)
	tmpl["index"] = template.Must(template.ParseFiles(
		append(components, "src/pages/index.gohtml")...,
	))
	tmpl["rules"] = template.Must(template.ParseFiles(
		append(components, "src/pages/rules.gohtml")...,
	))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	err := tmpl["index"].ExecuteTemplate(w, "base", TemplateData{
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
				Image:       "/terre-des-civilisations/background.webp",
			},
			Seasons: nil,
		},
	})
	if err != nil {
		panic(err)
	}
}

func handleRules(w http.ResponseWriter, r *http.Request) {
	err := tmpl["rules"].ExecuteTemplate(w, "base", TemplateData{
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
				Image:       "/purgatory.webp",
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

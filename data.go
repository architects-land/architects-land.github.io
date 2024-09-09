package main

import (
	_ "embed"
	"encoding/json"
)

//go:embed seasons.json
var seasonsRaw []byte

//go:embed team.json
var teamRaw []byte

var seasons map[string]Season

type Season struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Logo        string `json:"logo"`
	ID          string `json:"id"`
	RP          bool   `json:"rp"`
	Information struct {
		Description    []string `json:"description"`
		Version        string   `json:"version"`
		Mods           []string `json:"mods"`
		Video          string   `json:"video"`
		VideoThumbnail string   `json:"video_thumbnail"`
		Seed           string   `json:"seed"`
		Map            string   `json:"map"`
	} `json:"information"`
	Players []SeasonPlayer `json:"players"`
}

type SeasonPlayer struct {
	Name        string   `json:"name"`
	Pseudo      string   `json:"pseudo"`
	Description string   `json:"description"`
	RP          []string `json:"rp"`
	Links       []struct {
		Link string `json:"link"`
		Name string `json:"name"`
	} `json:"links"`
	Link string `json:"link"`
}

var team []PersonData

func init() {
	err := json.Unmarshal(seasonsRaw, &seasons)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(teamRaw, &team)
	if err != nil {
		panic(err)
	}
}

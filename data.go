package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
)

//go:embed seasons.json
var seasonsRaw []byte

//go:embed team.json
var teamRaw []byte

var seasons []Season

type CommonSeason struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Logo        string `json:"logo"`
	ID          string `json:"id"`
	RP          bool   `json:"rp"`
	Information struct {
		Description    []string `json:"description"`
		Particularity  []string `json:"particularity"`
		Version        string   `json:"version"`
		Mods           []string `json:"mods"`
		Video          string   `json:"video"`
		VideoThumbnail string   `json:"video_thumbnail"`
		Seed           string   `json:"seed"`
		Map            string   `json:"map"`
	} `json:"information"`
}

type Season struct {
	CommonSeason
	Players []*SeasonPlayer `json:"players"`
}

type FullSeasonData struct {
	CommonSeason
	Players []*PersonData
}

func (s *Season) toFullSeasonData() *FullSeasonData {
	fsg := FullSeasonData{CommonSeason: s.CommonSeason}
	for _, player := range s.Players {
		data := PersonData{
			Name:        player.Name,
			Image:       fmt.Sprintf("season/%s/skins/3d/%s.png", s.ID, player.Pseudo),
			Description: player.Description,
		}
		if s.RP {
			data.Link = "/season/" + s.ID + "/player/" + player.Pseudo
		} else {
			data.Link = player.Link
		}
		fsg.Players = append(fsg.Players, &data)
	}
	return &fsg
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

var team []*PersonData

func init() {
	err := parse()
	if err != nil {
		panic(err)
	}
}

func parse() error {
	var err error
	if dev {
		seasonsRaw, err = os.ReadFile("seasons.json")
		if err != nil {
			return err
		}
		teamRaw, err = os.ReadFile("team.json")
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(seasonsRaw, &seasons)
	if err != nil {
		return err
	}
	return json.Unmarshal(teamRaw, &team)
}

func GetSeason(id string) (*Season, bool) {
	if dev {
		err := parse()
		if err != nil {
			panic(err)
		}
	}
	for _, s := range seasons {
		if s.ID == id {
			return &s, true
		}
	}
	return nil, false
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Valorant struct {
	Status int `json:"status"`
	Data   []struct {
		UUID                      string   `json:"uuid"`
		DisplayName               string   `json:"displayName"`
		Description               string   `json:"description"`
		DeveloperName             string   `json:"developerName"`
		CharacterTags             any      `json:"characterTags"`
		DisplayIcon               string   `json:"displayIcon"`
		DisplayIconSmall          string   `json:"displayIconSmall"`
		BustPortrait              string   `json:"bustPortrait"`
		FullPortrait              string   `json:"fullPortrait"`
		FullPortraitV2            string   `json:"fullPortraitV2"`
		KillfeedPortrait          string   `json:"killfeedPortrait"`
		Background                string   `json:"background"`
		BackgroundGradientColors  []string `json:"backgroundGradientColors"`
		AssetPath                 string   `json:"assetPath"`
		IsFullPortraitRightFacing bool     `json:"isFullPortraitRightFacing"`
		IsPlayableCharacter       bool     `json:"isPlayableCharacter"`
		IsAvailableForTest        bool     `json:"isAvailableForTest"`
		IsBaseContent             bool     `json:"isBaseContent"`
		Role                      struct {
			UUID        string `json:"uuid"`
			DisplayName string `json:"displayName"`
			Description string `json:"description"`
			DisplayIcon string `json:"displayIcon"`
			AssetPath   string `json:"assetPath"`
		} `json:"role"`
		Abilities []struct {
			Slot        string `json:"slot"`
			DisplayName string `json:"displayName"`
			Description string `json:"description"`
			DisplayIcon string `json:"displayIcon"`
		} `json:"abilities"`
		VoiceLine struct {
			MinDuration float64 `json:"minDuration"`
			MaxDuration float64 `json:"maxDuration"`
			MediaList   []struct {
				ID    int    `json:"id"`
				Wwise string `json:"wwise"`
				Wave  string `json:"wave"`
			} `json:"mediaList"`
		} `json:"voiceLine"`
	} `json:"data"`
}

func main() {
	url := "https://valorant-api.com/v1/agents?isPlayableCharacter=true"

	timeClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-agent", "spacecount-total")
	res, geterr := timeClient.Do(req)
	if geterr != nil {
		fmt.Println(geterr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readerr := io.ReadAll(res.Body)
	if readerr != nil {
		fmt.Println(readerr)
	}
	var temp Valorant
	total := json.Unmarshal(body, &temp)
	if total != nil {
		fmt.Println(total)
	}
	for i := 0; i < len(temp.Data); i++ {
		fmt.Println(temp.Data[i].BustPortrait, temp.Data[i].DisplayName)
	}
}

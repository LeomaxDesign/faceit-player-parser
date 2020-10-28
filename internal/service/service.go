package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type FacietPlayerResponse struct {
	Avatar string `json:"avatar"`
	Bans   []struct {
	} `json:"bans"`
	Country            string   `json:"country"`
	CoverFeaturedImage string   `json:"cover_featured_image"`
	CoverImage         string   `json:"cover_image"`
	FaceitURL          string   `json:"faceit_url"`
	FriendsIds         []string `json:"friends_ids"`
	Games              struct {
		AdditionalProp1 struct {
			FaceitElo      int    `json:"faceit_elo"`
			GamePlayerID   string `json:"game_player_id"`
			GamePlayerName string `json:"game_player_name"`
			GameProfileID  string `json:"game_profile_id"`
			Region         string `json:"region"`
			Regions        struct {
			} `json:"regions"`
			SkillLevel      int    `json:"skill_level"`
			SkillLevelLabel string `json:"skill_level_label"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			FaceitElo      int    `json:"faceit_elo"`
			GamePlayerID   string `json:"game_player_id"`
			GamePlayerName string `json:"game_player_name"`
			GameProfileID  string `json:"game_profile_id"`
			Region         string `json:"region"`
			Regions        struct {
			} `json:"regions"`
			SkillLevel      int    `json:"skill_level"`
			SkillLevelLabel string `json:"skill_level_label"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			FaceitElo      int    `json:"faceit_elo"`
			GamePlayerID   string `json:"game_player_id"`
			GamePlayerName string `json:"game_player_name"`
			GameProfileID  string `json:"game_profile_id"`
			Region         string `json:"region"`
			Regions        struct {
			} `json:"regions"`
			SkillLevel      int    `json:"skill_level"`
			SkillLevelLabel string `json:"skill_level_label"`
		} `json:"additionalProp3"`
	} `json:"games"`
	Infractions struct {
	} `json:"infractions"`
	MembershipType string   `json:"membership_type"`
	Memberships    []string `json:"memberships"`
	NewSteamID     string   `json:"new_steam_id"`
	Nickname       string   `json:"nickname"`
	Platforms      struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"platforms"`
	PlayerID string `json:"player_id"`
	Settings struct {
		Language string `json:"language"`
	} `json:"settings"`
	SteamID64     string `json:"steam_id_64"`
	SteamNickname string `json:"steam_nickname"`
}

func GetFaceitByLink(username string) (bool, error) {
	var (
		err      error
		response FacietPlayerResponse
	)

	req, err := http.NewRequest("GET", "https://open.faceit.com/data/v4/players/"+username, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")
	req.Header.Set("Authorization", "Bearer 17081e1b-6cfb-4eb4-86b7-cd4a4b4e30f1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err

	}

	req.Close = true
	if resp != nil {
		defer resp.Body.Close()
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err

	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return false, err

	}

	if resp.StatusCode != 200 {
		return false, errors.New("Invalid Faceit: " + username)

	}

	if response.PlayerID == "" {
		return false, errors.New("Invalid Faceit: " + username)
	}

	return true, nil
}

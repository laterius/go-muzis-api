package go_muzis_api

import (
	"github.com/go-resty/resty"
	"encoding/json"
	"strings"
	"fmt"
	"strconv"
)
const (
	SEACH_API = "http://muzis.ru/api/search.api"
	SEACH_PERFORMER_API = "http://muzis.ru/api/get_songs_by_performer.api"
)

type SearchResponse struct {
	Songs 		*[]Song        `json:"songs"`
	Performer	*[]Performer	`json:"performers"`
}

type SearchPerformersResponse struct {
	Performer	*[]Performer	`json:"performers"`
}

type SearchSongsResponse struct {
	Songs 		*[]Song        `json:"songs"`
}

type Song struct {
	Type          int        `json:"type"`
	TrackId       int           `json:"id"`
	TrackName     string        `json:"track_name"`
	Lyrics        string        `json:"lyrics"`
	PerformerName string        `json:"performer"`
	PerformerId   int           `json:"performer_id"`
	Poster        string        `json:"poster"`
	Url          string        `json:"file_mp3"`
	Duration      int        `json:"timestudy"`
}


type Performer struct{
	PerformerId 	int	`json:"id"`
	PerformerName	string	`json:"title"`
	Poster		string	`json:"poster"`
}

type MuzisApi struct {

}




func (api *MuzisApi) GetSongsByPerformerName(name string) []Song {
	performerId := getPerformerIdByName(name)
	if performerId !=0 {
		return api.GetSongsByPerformerId(performerId)
	}

	return []Song{}
}

func (api *MuzisApi) GetSongsByPerformerId(performerId int) []Song{
		response, err := resty.R().
		SetFormData(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}).
		SetFormData(map[string]string{
		"performer_id": strconv.Itoa(performerId),
	}).
		Post(SEACH_PERFORMER_API)

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	r := response.String()


	responseAsBytes := []byte(r)


	searchSongsResponse := &SearchSongsResponse{}
	if err := json.Unmarshal(responseAsBytes, searchSongsResponse); err != nil {
		panic(err)
	}

	return *searchSongsResponse.Songs
}


func getPerformerIdByName(name string) int{
	response, err := resty.R().
		SetFormData(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}).
		SetFormData(map[string]string{
		"q_track": name,
	}).
		Post(SEACH_API)

	if err != nil {
		panic(err)
	}

	responseAsBytes := []byte(response.String())
	searchResponse := &SearchResponse{}
	if err := json.Unmarshal(responseAsBytes, searchResponse); err != nil {
		panic(err)
	}

	for _ ,song := range *searchResponse.Songs{
		if strings.ToLower(song.PerformerName) == strings.ToLower(name){
			return song.PerformerId
		}
	}
	return 0
}

//func (api MuzisApi) GetArtist(string name)

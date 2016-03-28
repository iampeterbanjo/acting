package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// APIRoot url for movie info
var APIRoot = "https://api.themoviedb.org/3"

// APIKey for movie info
var APIKey = ""

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	APIKey = os.Getenv("TMDB_API_KEY")
}

// Actor from TMDB
type Actor struct {
	Popularity  float64 `json:"popularity"`
	Name        string  `json:"name"`
	ID          int     `json:"id"`
	ProfilePath string  `json:"profile_path"`
	Credits     []Credit
}

// ActorSearchResults from API
type ActorSearchResults struct {
	Page         int     `json:"page"`
	Results      []Actor `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

// FetchActor from TMDB or error
func FetchActor(name string) (Actor, error) {

	u := fmt.Sprintf("%s/search/person?api_key=%s&query=%s", APIRoot, APIKey, url.QueryEscape(name))
	results := ActorSearchResults{}

	res, err := http.Get(u)
	a := Actor{}
	if err != nil {
		return a, err
	}

	err = json.NewDecoder(res.Body).Decode(&results)
	if err != nil {
		return a, err
	}

	if results.TotalResults == 0 {
		return a, fmt.Errorf("There are no search results for: %s", name)
	}

	return results.Results[0], nil
}

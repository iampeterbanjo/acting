package main

import "os"

// APIRoot url for movie info
var APIRoot = "https://api.themoviedb.org/3"

// APIKey for movie info
var APIKey = os.Getenv("TMDB_API_KEY")

// Actor from IMDB
type Actor struct {
	Popularity  float64 `json:"popularity"`
	Name        string  `json:"name"`
	ID          int     `json:"id"`
	ProfilePath float64 `json:"profile_path"`
}

// ActorSearchResults from API
type ActorSearchResults struct {
	Page         int     `json:"page"`
	Results      []Actor `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

func fetch(name string) (Actor, error) {

}

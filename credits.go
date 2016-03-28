package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Credit for Actor
type Credit struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

// NameOrTitle for TV or movie respectively
func (c Credit) NameOrTitle() string {
	if c.Name != "" {
		return c.Name
	}
	return c.Title
}

// CreditSearchResults returns Cast credit
type CreditSearchResults struct {
	Cast []Credit `json:"cast"`
}

// FetchCredits for Actor
func FetchCredits(actor *Actor) error {
	u := fmt.Sprintf("%s/person/%d/combined_credits?api_key=%s", APIRoot, actor.ID, APIKey)
	results := CreditSearchResults{}

	res, err := http.Get(u)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&results)
	if err != nil {
		return err
	}

	actor.Credits = results.Cast
	return nil
}

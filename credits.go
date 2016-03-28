package main

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
	return nil
}

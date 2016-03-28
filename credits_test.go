package main_test

import (
	"acting_buddies"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchCredits_WithResults(t *testing.T) {
	a := assert.New(t)
	body := `{
    "cast": [
      {
        "id": 5966,
        "title": "Along Came Polly"
      },
      {
        "id": 1668,
        "name": "Friends"
      }
    ]
  }`

	actor := &main.Actor{}
	FakeServer(body, func() {
		err := main.FetchCredits(actor)
		a.NoError(err)
		credits := actor.Credits
		a.Equal(2, len(credits))
		a.Equal("Along Came Polly", credits[0].NameOrTitle())
		a.Equal("Friends", credits[1].NameOrTitle())
	})
}

func Test_FetchCredits_WithoutResults(t *testing.T) {
	a := assert.New(t)
	body := `{
    "cast": []
  }`

	actor := &main.Actor{}
	FakeServer(body, func() {
		err := main.FetchCredits(actor)
		a.NoError(err)
		credits := actor.Credits
		a.Equal(0, len(credits))
	})
}

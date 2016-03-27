package main_test

import (
	"acting_buddies"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchActor_WithResults(t *testing.T) {
	a := assert.New(t)

	body := `{
    "page":1,
    "results":[
      {
        "id":287,
        "name":"Brad Pitt",
        "popularity":12,
        "profile_path":"/brad.jpg"
      }
    ],
    "total_pages":1,
    "total_results":1
  }`
	FakeServer(body, func() {
		actor, err := main.FetchActor("Brad Pitt")
		a.NoError(err)
		a.Equal("Brad Pitt", actor.Name)
	})
}

func Test_FetchActor_WithoutResults(t *testing.T) {
	a := assert.New(t)

	body := `{
    "page":1,
    "results":[],
    "total_pages":1,
    "total_results":0
  }`
	FakeServer(body, func() {
		_, err := main.FetchActor("Brad Pitt")
		a.Error(err)
		a.Equal("There are no search results for: Brad Pitt", err.Error())
	})
}

func FakeServer(b string, f func()) {
	root := main.APIRoot

	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, b)
	}))
	defer ts.Close()

	main.APIRoot = ts.URL

	main.APIRoot = root
}

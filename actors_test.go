package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchActor_WithResults(t *testing.T) {
	a := assert.New(t)

	actor, err = main.FetchActor("Brad Pitt")
	a.NoError(err)
	a.Equal("Brad Pitt", actor.Name)
}

func Test_FetchActor_WithoutResults(t *testing.T) {
	a := assert.New(t)

	_, err := main.FetchActor("Brad Pitt")
	a.Error(err)
	a.Equal("There are no search results for: Brad Pitt", err.Error())
}

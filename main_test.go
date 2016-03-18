package main_test

import (
	"acting-buddies"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AskForName(t *testing.T) {
	setup()

	b := []byte("Mark\n")

	r := bytes.NewBuffer(b)
	main.AskForName(r)

	a := assert.New(t)
	a.Equal(len(main.ActorNames), 1)
	a.Equal(main.ActorNames[0], "Mark")
}

func Test_AskForNames(t *testing.T) {
	setup()

	b := []byte("Mark\nBates\n")

	r := bytes.NewBuffer(b)
	main.AskForNames(r)

	a := assert.New(t)
	a.Equal(len(main.ActorNames), 2)
	a.Equal(main.ActorNames[0], "Mark")
	a.Equal(main.ActorNames[1], "Bates")
}

// tear down
func setup() {
	main.ActorNames = []string{}
}

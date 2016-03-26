package main_test

import (
	"acting_buddies"
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

	b := []byte("Mark\nBates\nn\n")

	r := bytes.NewBuffer(b)
	main.AskForNames(r)

	a := assert.New(t)
	a.Equal(len(main.ActorNames), 2)
	a.Equal(main.ActorNames[0], "Mark")
	a.Equal(main.ActorNames[1], "Bates")
}

func Test_AskForNames_FourNames(t *testing.T) {
	setup()

	b := []byte("Mark\nBates\ny\nAngelina\ny\nJolie\nn\n")

	r := bytes.NewBuffer(b)
	main.AskForNames(r)

	a := assert.New(t)
	a.Equal(len(main.ActorNames), 4)
	a.Equal(main.ActorNames[0], "Mark")
	a.Equal(main.ActorNames[1], "Bates")
	a.Equal(main.ActorNames[2], "Angelina")
	a.Equal(main.ActorNames[3], "Jolie")
}

// tear down
func setup() {
	main.ActorNames = []string{}
}

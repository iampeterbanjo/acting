package main_test

import (
	"acting-buddies"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AskForName(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Mark\n")

	r := bytes.NewBuffer(b)
	main.AskForName(r)

	a.Equal(len(main.ActorNames), 1)
	a.Equal(main.ActorNames[0], "Mark")
}

// tear down
func setup() {
	main.ActorNames = []string{}
}

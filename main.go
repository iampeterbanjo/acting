package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ActorNames from terminal
var ActorNames = []string{}

type stringReader interface {
	ReadString(byte) (string, error)
}

// AskForName gets a name from terminal
func AskForName(r stringReader) {
	fmt.Println("Please enter an actor's name:")
	name, _ := r.ReadString('\n')
	name = strings.TrimSpace(name)
	ActorNames = append(ActorNames, name)
}

// AskForNames will AskForName until we have 2 names
func AskForNames(r stringReader) {
	asking := true
	for asking {
		if len(ActorNames) < 2 {
			AskForName(r)
		} else {
			asking = false
		}
	}
}

func main() {
	AskForNames(bufio.NewReader(os.Stdin))
}

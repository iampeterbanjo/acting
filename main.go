package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/markbates/going/wait"
)

// ActorNames from terminal
var ActorNames = []string{}

// Run user interaction
func Run(in stringReader, out io.Writer) {
	AskForNames(in)

	actors := []Actor{}
	m := sync.Mutex{}

	fmt.Fprintf(out, "\nYou selected the following %d actors: \n", len(ActorNames))
	wait.Wait(len(ActorNames), func(i int) {
		actor, err := FetchActor(ActorNames[i])
		if err != nil {
			log.Fatal(err)
		}
		m.Lock()
		actors = append(actors, actor)
		m.Unlock()
		fmt.Fprintln(out, actor.Name)
	})

	credits := FilterCredits(actors)
	if len(credits) > 0 {
		fmt.Fprintln(out, "\nThey have appeared in the following movies and TV shows together:")
		for _, c := range credits {
			fmt.Fprintln(out, c.NameOrTitle())
		}
	} else {
		fmt.Fprintln(out, "\nHave not appeared in anything together.")
	}
}

func main() {
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/markbates/going/wait"
)

// ActorNames from terminal
var ActorNames = []string{}

// Run user interaction
func Run(in stringReader, out io.Writer) {
	AskForNames(in)

	fmt.Fprintf(out, "You selected the following %d actors: \n", len(ActorNames))
	wait.Wait(len(ActorNames), func(i int) {
		actor, err := FetchActor(ActorNames[i])
		if err != nil {
			log.Panic(err)
		}
		fmt.Fprintln(out, actor.Name)
	})
}

func main() {
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}

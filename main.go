package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Run user interaction
func Run(in stringReader, out io.Writer) {
	ActorNames = []string{}
	AskForNames(in)

	fmt.Fprintf(out, "You selected the following %d actors: \n", len(ActorNames))
	for _, v := range ActorNames {
		fmt.Fprintln(out, v)
	}
}

func main() {
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Run user interaction
func Run(in stringReader, out io.Writer) {
	AskForNames(in)

	fmt.Printf("You selected the following %d actors: \n", len(ActorNames))
	for _, v := range ActorNames {
		fmt.Println(out, v)
	}
}

func main() {
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}

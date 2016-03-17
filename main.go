package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ActorsNames a string array
var ActorsNames = []string{}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter an actor's name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println(name)
}

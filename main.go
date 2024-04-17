package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Printf("Hey! This is the Monkey repl!\n")
	fmt.Printf("Feel free to type any Monkey commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

package main

import (
	"os"

	"github.com/shozawa/monkey2/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}

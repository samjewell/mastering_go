package main

import (
	"fmt"
	"github.com/mactsouk/go/simpleGitHub"
	// requires me to run
	// `go get -v github.com/mactsouk/go/simpleGitHub`
)

func main() {
	fmt.Println(simpleGitHub.AddTwo(5, 6))
}
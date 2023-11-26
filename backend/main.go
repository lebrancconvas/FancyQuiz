package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {

}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()

	// Start Database

	// Server Start 
}

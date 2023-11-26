package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lebrancconvas/FancyQuiz/db"
	"github.com/lebrancconvas/FancyQuiz/server"
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
	db.Init()

	// Server Start
	server.Start()
}

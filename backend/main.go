package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lebrancconvas/FancyQuiz/db"
	"github.com/lebrancconvas/FancyQuiz/server"
)

func init() {
	// Localhost Environment.
	os.Setenv("PORT", "8081")
	os.Setenv("DBHOST", "localhost")
	os.Setenv("DBPORT", "5432")
	os.Setenv("DBUSER", "postgres")
	os.Setenv("DBPASSWORD", "P@ssw0rd")
	os.Setenv("DBNAME", "postgres")
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

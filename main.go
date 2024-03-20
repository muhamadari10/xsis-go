package main

import (
	"log"
	"movie-app/cmd"
	env "movie-app/config"
	"movie-app/database"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	//Setup env
	di, err := env.NewENV("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Database
	if di.DB, di.Err = database.NewPostGresDB(di.Params); di.Err != nil {
		// Handle with middleware here upon error
		log.Fatal(di.Err)
	}

	// Run service
	cli := cmd.NewCLI(di, os.Args)
	if cli.Start(); cli.Error() != nil {
		log.Fatal(cli.Error())
	}
}

package database

import (
	"database/sql"
	"errors"
	"log"
	env "movie-app/config"
)

type Env struct {
	Params *env.EnvironmentParameters
	Error  error
}

func NewPostGresDB(config env.EnvironmentParameters) (*sql.DB, error) {
	conn, err := sql.Open(config.Database.PGDB.DBType, config.SetupPostGresDBConnection().Database.PGDB.DBConfig)
	if err != nil {
		return conn, errors.New("failed connecting to database : " + err.Error())
	}
	conn.SetMaxOpenConns(100)

	if err = conn.Ping(); err != nil {
		return conn, errors.New("failed pinging to database : " + err.Error())
	}
	log.Printf("Connected to %s database on port : %s ", config.Database.PGDB.DBName, config.Database.PGDB.DBPort)

	return conn, err
}

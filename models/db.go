package models

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	// import postgres driver
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    username text,
    email text,
		passwordhash text,

		admin boolean,
		passwordreset boolean,
		confirmed boolean
);

CREATE TABLE IF NOT EXISTS series (
    id serial PRIMARY KEY,
    title text UNIQUE,
    description text,
		posterurl text,
		seasons int
);

CREATE TABLE IF NOT EXISTS episodes (
    id serial PRIMARY KEY,
		title text,
    releasedate date,
		season int,
		number int,
		seriesid int references series(id)
)`

// OpenDB gets a new handle to the database or returns the existing one
func OpenDB() *sqlx.DB {
	log.Println("pq: connecting to database")

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("pq: connected to database")

	db.MustExec(schema)

	return db
}

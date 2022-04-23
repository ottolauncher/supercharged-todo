package db

import (
	"github.com/joho/godotenv"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"log"
	"os"
)

func Init() *r.Session {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	uri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")

	session, err := r.Connect(r.ConnectOpts{
		Address:  uri,
		Database: dbName,
	})
	if err != nil {
		log.Fatal(err)
	}
	return session
}

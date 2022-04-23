package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
	db "github.com/ottolauncher/supercharged-todo/graph/db/rethinkdb"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ottolauncher/supercharged-todo/graph"
	"github.com/ottolauncher/supercharged-todo/graph/generated"
)

const defaultPort = "8080"

func main() {
	var (
		once    sync.Once
		session *r.Session
	)
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	once.Do(func() {
		session = db.Init()
	})

	um := db.NewUserManager(session, "users")
	rm := db.NewRoleManager(session, "roles")
	tm := db.NewTodoManager(session, "todos")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UM: um, RM: rm, TM: tm, Session: session,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	db "github.com/ottolauncher/supercharged-todo/graph/db/rethinkdb"
	"github.com/ottolauncher/supercharged-todo/graph/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ottolauncher/supercharged-todo/graph"
	"github.com/ottolauncher/supercharged-todo/graph/generated"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
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
		UM: um, RM: rm, TM: tm, Session: session, TodoChannels: make(chan *model.Todo, 1),
	}}))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// check against your desired domains here
				// return r.Host == "example.org"
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	// srv.Use(extension.Introspection{}) this is include by default with NewDefaultServer

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown: ", err)
	}
	log.Println("server exiting")

}

//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	db "github.com/ottolauncher/supercharged-todo/graph/db/rethinkdb"
	"github.com/ottolauncher/supercharged-todo/graph/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RM           *db.RoleManager
	UM           *db.UserManager
	TM           *db.TodoManager
	Session      *r.Session
	TodoChannels chan *model.Todo
}

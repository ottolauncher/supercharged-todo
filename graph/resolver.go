package graph

import db "github.com/ottolauncher/supercharged-todo/graph/db/rethinkdb"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RM db.RoleManager
	UM db.UserManager
	TM db.TodoManager
}

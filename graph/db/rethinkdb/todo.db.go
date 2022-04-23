package db

import (
	"context"
	"github.com/ottolauncher/supercharged-todo/graph/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type ITodo interface {
	Create(ctx context.Context, args model.Todo) (*model.Todo, error)
	Update(ctx context.Context, args model.Todo) (*model.Todo, error)
	Delete(ctx context.Context, filter map[string]interface{}) error
	Get(ctx context.Context, filter map[string]interface{}) (*model.Todo, error)
	All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.Todo, error)
	Search(ctx context.Context, query string) ([]*model.Todo, error)
}

type TodoManager struct {
	Session *r.Session
	tblName string
}

func NewTodoManager(session *r.Session, tblName string) *TodoManager {
	return &TodoManager{Session: session, tblName: tblName}
}

func (t TodoManager) Create(ctx context.Context, args model.Todo) (*model.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoManager) Update(ctx context.Context, args model.Todo) (*model.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoManager) Delete(ctx context.Context, filter map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoManager) Get(ctx context.Context, filter map[string]interface{}) (*model.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoManager) All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoManager) Search(ctx context.Context, query string) ([]*model.Todo, error) {
	//TODO implement me
	panic("implement me")
}

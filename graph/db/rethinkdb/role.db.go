package db

import (
	"context"
	"github.com/ottolauncher/supercharged-todo/graph/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type IRole interface {
	Create(ctx context.Context, args model.Role) (*model.Role, error)
	Update(ctx context.Context, args model.UpdateRole) (*model.Role, error)
	Delete(ctx context.Context, filter map[string]interface{}) error
	Get(ctx context.Context, filter map[string]interface{}) (*model.Role, error)
	All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.Role, error)
	Search(ctx context.Context, query string) ([]*model.Role, error)
}

type RoleManager struct {
	Session *r.Session
	tblName string
}

func NewRoleManager(session *r.Session, tblName string) *RoleManager {
	return &RoleManager{
		Session: session,
		tblName: tblName,
	}
}

func (r RoleManager) Create(ctx context.Context, args model.Role) (*model.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleManager) Update(ctx context.Context, args model.UpdateRole) (*model.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleManager) Delete(ctx context.Context, filter map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (r RoleManager) Get(ctx context.Context, filter map[string]interface{}) (*model.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleManager) All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleManager) Search(ctx context.Context, query string) ([]*model.Role, error) {
	//TODO implement me
	panic("implement me")
}

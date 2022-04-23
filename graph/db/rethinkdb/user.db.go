package db

import (
	"context"
	"github.com/ottolauncher/supercharged-todo/graph/model"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type IUser interface {
	Create(ctx context.Context, args model.User) (*model.User, error)
	Update(ctx context.Context, args model.User) (*model.User, error)
	Delete(ctx context.Context, filter map[string]interface{}) error
	Get(ctx context.Context, filter map[string]interface{}) (*model.User, error)
	All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.User, error)
	Search(ctx context.Context, query string) ([]*model.User, error)
}

type UserManager struct {
	Session *r.Session
	tblName string
}

func NewUserManager(session *r.Session, tblName string) *UserManager {
	return &UserManager{Session: session, tblName: tblName}
}

func (u UserManager) Create(ctx context.Context, args model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserManager) Update(ctx context.Context, args model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserManager) Delete(ctx context.Context, filter map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (u UserManager) Get(ctx context.Context, filter map[string]interface{}) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserManager) All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserManager) Search(ctx context.Context, query string) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

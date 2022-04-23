package db

import (
	"context"
	"time"

	"github.com/ottolauncher/supercharged-todo/graph/model"
	"github.com/ottolauncher/supercharged-todo/helpers/text"
	rd "gopkg.in/rethinkdb/rethinkdb-go.v6"
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
	Session *rd.Session
	tblName string
}

func NewRoleManager(session *rd.Session, tblName string) *RoleManager {
	return &RoleManager{
		Session: session,
		tblName: tblName,
	}
}

func (r RoleManager) Create(ctx context.Context, args model.Role) (*model.Role, error) {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	tmpStr := text.Slugify(args.Name)
	args.Slug = &tmpStr
	res, err := rd.Table(r.tblName).Insert(args).RunWrite(r.Session)
	if err != nil {
		return nil, err
	}
	args.ID = res.GeneratedKeys[0]

	return &args, nil
}

func (r RoleManager) Update(ctx context.Context, args model.UpdateRole) (*model.Role, error) {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	tmpStr := text.Slugify(args.Name)
	now := time.Now()
	var role model.Role
	cursor, err := rd.Table(r.tblName).Get(args.ID).Run(r.Session)
	defer func(cursor *rd.Cursor) {
		err := cursor.Close()
		if err != nil {

		}
	}(cursor)
	if err != nil {
		return nil, err
	}
	err = cursor.One(&role)
	if len(*args.Description) > 0 {
		role.Description = args.Description
	}
	role.Slug = &tmpStr
	role.UpdatedAt = &now
	role.Name = args.Name
	_, updateErr := rd.Table(r.tblName).Get(args.ID).Update(role).RunWrite(r.Session)
	if updateErr != nil {
		return nil, updateErr
	}
	return &role, nil
}

func (r RoleManager) Delete(ctx context.Context, filter map[string]interface{}) error {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()

	_, err := rd.Table(r.tblName).Filter(filter).Delete().RunWrite(r.Session)
	if err != nil {
		return err
	}
	return nil
}

func (r RoleManager) Get(ctx context.Context, filter map[string]interface{}) (*model.Role, error) {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	cursor, err := rd.Table(r.tblName).Filter(filter).Run(r.Session)
	defer func(cursor *rd.Cursor) {
		err := cursor.Close()
		if err != nil {

		}
	}(cursor)
	if err != nil {
		return nil, err
	}
	var role model.Role
	err = cursor.One(&role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r RoleManager) All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.Role, error) {
	_, cancel := context.WithTimeout(ctx, 650*time.Millisecond)
	defer cancel()
	var roles []*model.Role
	cursor, err := rd.Table(r.tblName).Skip(page).Limit(limit).Filter(filter).Run(r.Session)
	defer func(cursor *rd.Cursor) {
		err := cursor.Close()
		if err != nil {

		}
	}(cursor)
	if err != nil {
		return nil, err
	}
	err = cursor.All(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r RoleManager) Search(ctx context.Context, query string) ([]*model.Role, error) {
	_, cancel := context.WithTimeout(ctx, 450*time.Millisecond)
	defer cancel()
	var roles []*model.Role
	cursor, err := rd.Table(r.tblName).Filter(func(row rd.Term) interface{} {
		return rd.Expr([]string{"description", "name"}).Contains(func(key rd.Term) interface{} {
			return row.Field(key).CoerceTo("string").Match("(?i)" + query)
		})
	}).Run(r.Session)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	err = cursor.All(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

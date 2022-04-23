package db

import (
	"context"
	"strings"
	"time"

	"github.com/ottolauncher/supercharged-todo/graph/model"
	"github.com/ottolauncher/supercharged-todo/helpers/text"
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

func mergeUsers(todo r.Term) r.Term {
	return r.Table("users").GetAll(r.Args(todo.Field("assigned_ids"))).CoerceTo("array").
		Merge(func(user r.Term) interface{} {
			return map[string]interface{}{
				"roles": mergeRole(user),
			}
		})
}

func mergeUser(todo r.Term) r.Term {
	return r.Table("users").Get(todo.Field("user_id")).Merge(func(user r.Term) interface{} {
		return map[string]interface{}{
			"roles": mergeRole(user),
		}
	})
}

func NewTodoManager(session *r.Session, tblName string) *TodoManager {
	return &TodoManager{Session: session, tblName: tblName}
}

func (t TodoManager) Create(ctx context.Context, args model.Todo) (*model.Todo, error) {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()

	tmpStr := strings.Split(args.Text, " ")
	slug := text.Slugify(strings.Join(tmpStr[:2], " "))
	args.Slug = &slug
	res, err := r.Table(t.tblName).Insert(args).RunWrite(t.Session)
	if err != nil {
		return nil, err
	}
	args.ID = res.GeneratedKeys[0]
	args.User = nil
	args.Assigned = nil
	return &args, nil
}

func (t TodoManager) Update(ctx context.Context, args model.Todo) (*model.Todo, error) {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()

	tmpStr := strings.Split(args.Text, " ")
	slug := text.Slugify(strings.Join(tmpStr[:2], " "))
	args.Slug = &slug
	res, err := r.Table(t.tblName).Get(args.ID).Update(args).RunWrite(t.Session)
	if err != nil {
		return nil, err
	}
	args.ID = res.GeneratedKeys[0]
	args.User = nil
	args.Assigned = nil
	return &args, nil
}

func (t TodoManager) Delete(ctx context.Context, filter map[string]interface{}) error {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()

	_, err := r.Table(t.tblName).Filter(filter).Delete().RunWrite(t.Session)
	if err != nil {
		return err
	}
	return nil
}

func (t TodoManager) Get(ctx context.Context, filter map[string]interface{}) (*model.Todo, error) {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()

	res, err := r.Table(t.tblName).Filter(filter).Merge(func(todo r.Term) interface{} {
		return map[string]interface{}{
			"assigned": mergeUsers(todo),
			"user":     mergeUser(todo),
		}
	}).Run(t.Session)
	if err != nil {
		return nil, err
	}
	var todo model.Todo
	err = res.One(&todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t TodoManager) All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.Todo, error) {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()

	res, err := r.Table(t.tblName).Skip(page).Limit(limit).
		Filter(filter).Merge(func(todo r.Term) interface{} {
		return map[string]interface{}{
			"assigned": mergeUsers(todo),
			"user":     mergeUser(todo),
		}
	}).Run(t.Session)
	if err != nil {
		return nil, err
	}
	var todos []*model.Todo
	err = res.All(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t TodoManager) Search(ctx context.Context, query string) ([]*model.Todo, error) {
	_, cancel := context.WithTimeout(ctx, 650*time.Millisecond)
	defer cancel()
	var todos []*model.Todo
	cursor, err := r.Table(t.tblName).Filter(func(row r.Term) interface{} {
		return r.Expr([]string{"text", "description"}).Contains(func(key r.Term) interface{} {
			return row.Field(key).CoerceTo("string").Match("(?i)" + query)
		})
	}).Merge(func(todo r.Term) interface{} {
		return map[string]interface{}{
			"assigned": mergeUsers(todo),
			"user":     mergeUser(todo),
		}
	}).Run(t.Session)
	if err != nil {
		return nil, err
	}
	err = cursor.All(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

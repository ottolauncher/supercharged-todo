package db

import (
	"context"
	"time"

	"github.com/ottolauncher/supercharged-todo/graph/model"
	"github.com/ottolauncher/supercharged-todo/helpers/text"
	"golang.org/x/crypto/bcrypt"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type IUser interface {
	Create(ctx context.Context, args model.User) (*model.User, error)
	Update(ctx context.Context, args model.User) (*model.User, error)
	Delete(ctx context.Context, filter map[string]interface{}) error
	Get(ctx context.Context, filter map[string]interface{}) (*model.User, error)
	All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.User, error)
	Search(ctx context.Context, query string) ([]*model.User, error)
	CheckPassword(ctx context.Context, password, hash string) bool
}

type UserManager struct {
	Session *r.Session
	tblName string
}

func NewUserManager(session *r.Session, tblName string) *UserManager {
	return &UserManager{Session: session, tblName: tblName}
}

func mergeRole(user r.Term) r.Term {
	return r.Table("roles").GetAll(r.Args(user.Field("role_ids"))).CoerceTo("array")
}

// hashPassword hash the user password before save it
func hashPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compare the stored and the incoming password tp match a specific user account
func (u *UserManager) CheckPassword(ctx context.Context, password string, hash string) bool {
	_, cancel := context.WithTimeout(ctx, 250*time.Millisecond)
	defer cancel()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *UserManager) Create(ctx context.Context, args model.User) (*model.User, error) {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	pwd, err := hashPassword([]byte(args.Password))
	if err != nil {
		return nil, err
	}
	args.Password = pwd
	slug := text.Slugify(args.Username)
	args.Slug = &slug
	res, err := r.Table(u.tblName).Insert(args).RunWrite(u.Session)
	if err != nil {
		return nil, err
	}
	args.Email = "-"
	args.ID = res.GeneratedKeys[0]
	return &args, nil
}

func (u *UserManager) Update(ctx context.Context, args model.User) (*model.User, error) {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	now := time.Now()
	args.UpdatedAt = &now
	_, err := r.Table(u.tblName).Get(args.ID).Update(args).RunWrite(u.Session)
	if err != nil {
		return nil, err
	}
	return &args, nil
}

func (u *UserManager) Delete(ctx context.Context, filter map[string]interface{}) error {
	_, err := r.Table(u.tblName).Filter(filter).RunWrite(u.Session)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserManager) Get(ctx context.Context, filter map[string]interface{}) (*model.User, error) {
	_, cancel := context.WithTimeout(ctx, 350*time.Millisecond)
	defer cancel()
	cursor, err := r.Table(u.tblName).Filter(filter).Merge(func(p r.Term) interface{} {
		return map[string]interface{}{"roles": mergeRole(p)}
	}).Run(u.Session)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	var user model.User
	err = cursor.One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserManager) All(ctx context.Context, filter map[string]interface{}, limit int, page int) ([]*model.User, error) {
	_, cancel := context.WithTimeout(ctx, 650*time.Millisecond)
	defer cancel()
	cursor, err := r.Table(u.tblName).Filter(filter).Merge(func(p r.Term) interface{} {
		return map[string]interface{}{"roles": mergeRole(p)}
	}).Run(u.Session)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	var users []*model.User
	err = cursor.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserManager) Search(ctx context.Context, query string) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

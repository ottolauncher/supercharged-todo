package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/ottolauncher/supercharged-todo/graph/generated"
	"github.com/ottolauncher/supercharged-todo/graph/model"
	rd "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	now := time.Now()
	todo := model.Todo{
		Text:        input.Text,
		Description: input.Description,
		CreatedAt:   &now,
		UserID:      input.UserID,
		Start:       input.Start,
		End:         input.End,
		AssignedIDs: input.UserIds,
	}
	res, err := r.TM.Create(ctx, todo)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	todo := model.Todo{
		ID:          input.ID,
		Text:        input.Text,
		Description: input.Description,
		UserID:      input.UserID,
		Start:       input.Start,
		End:         input.End,
		AssignedIDs: input.UserIds,
	}
	res, err := r.TM.Update(ctx, todo)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input map[string]interface{}) (bool, error) {
	if len(input) == 0 {
		return false, errors.New("you need to setup a filter like {'field': 'something'} first")
	}
	if err := r.TM.Delete(ctx, input); err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	now := time.Now()
	var user model.User
	if len(input.Password1) > 0 && len(input.Password2) > 0 {
		if input.Password1 != input.Password2 {
			return nil, errors.New("password missmatch")
		}
		user.Password = input.Password2
	}
	if len(input.Email) > 0 {
		addr, err := mail.ParseAddress(input.Email)
		if err != nil {
			return nil, err
		}
		user.Email = addr.Address
	}
	user.Biography = input.Biography
	user.RoleIDs = input.Roles
	user.CreatedAt = &now
	user.Username = input.Username

	res, err := r.UM.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	var user model.User
	cursor, err := rd.Table("users").Get(input.ID).Run(r.Session)
	if err != nil {
		return nil, err
	}
	err = cursor.One(&user)
	if err != nil {
		return nil, err
	}

	if !r.UM.CheckPassword(ctx, input.OldPassword, user.Password) {
		return nil, errors.New("user does not exists")
	}
	if len(input.Password1) > 0 && len(input.Password2) > 0 {
		if input.Password1 != input.Password2 {
			return nil, errors.New("password missmatch")
		}
		user.Password = input.Password2
	}
	if len(input.Email) > 0 {
		addr, err := mail.ParseAddress(input.Email)
		if err != nil {
			return nil, err
		}
		user.Email = addr.Address
	}
	if len(*input.Biography) > 0 {
		user.Biography = input.Biography
	}
	if len(input.Roles) > 0 {
		user.RoleIDs = input.Roles
	}
	if len(input.Username) > 0 {
		user.Username = input.Username
	}
	res, err := r.UM.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input map[string]interface{}) (bool, error) {
	if err := r.UM.Delete(ctx, input); err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) CreateRole(ctx context.Context, input model.NewRole) (*model.Role, error) {
	now := time.Now()
	role := model.Role{
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}
	res, err := r.RM.Create(ctx, role)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) UpdateRole(ctx context.Context, input model.UpdateRole) (*model.Role, error) {
	res, err := r.RM.Update(ctx, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *mutationResolver) DeleteRole(ctx context.Context, input map[string]interface{}) (bool, error) {
	if err := r.RM.Delete(ctx, input); err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Todos(ctx context.Context, filter map[string]interface{}, limit *int, page *int) ([]*model.Todo, error) {
	res, err := r.TM.All(ctx, filter, *limit, *page)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Todo(ctx context.Context, filter map[string]interface{}) (*model.Todo, error) {
	res, err := r.TM.Get(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Users(ctx context.Context, filter map[string]interface{}, limit *int, page *int) ([]*model.User, error) {
	res, err := r.UM.All(ctx, filter, *limit, *page)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) User(ctx context.Context, filter map[string]interface{}) (*model.User, error) {
	res, err := r.UM.Get(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Roles(ctx context.Context, filter map[string]interface{}, limit *int, page *int) ([]*model.Role, error) {
	res, err := r.RM.All(ctx, filter, *limit, *page)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Role(ctx context.Context, filter map[string]interface{}) (*model.Role, error) {
	res, err := r.RM.Get(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) Search(ctx context.Context, query string, limit *int, page *int) ([]model.SearchResult, error) {
	var (
		res         []model.SearchResult
		searchError []error
	)
	// TODO: add user and todos also
	roles, rErr := r.RM.Search(ctx, query)
	if rErr != nil {
		searchError = append(searchError, rErr)
	}
	users, uErr := r.UM.Search(ctx, query)
	if uErr != nil {
		searchError = append(searchError, uErr)
	}
	todos, tErr := r.TM.Search(ctx, query)
	if tErr != nil {
		searchError = append(searchError, tErr)
	}
	if len(searchError) > 0 {
		return nil, fmt.Errorf("%s", searchError)
	}
	if len(roles) > 0 {
		for _, r := range roles {
			res = append(res, r)
		}
	}
	if len(todos) > 0 {
		for _, todo := range todos {
			res = append(res, todo)
		}
	}

	if len(users) > 0 {
		for _, u := range users {
			res = append(res, u)
		}
	}
	return res, nil
}

func (r *subscriptionResolver) TodoAdded(ctx context.Context) (<-chan *model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

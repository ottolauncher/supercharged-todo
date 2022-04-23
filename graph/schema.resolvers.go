package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/ottolauncher/supercharged-todo/graph/generated"
	"github.com/ottolauncher/supercharged-todo/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input map[string]interface{}) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input map[string]interface{}) (bool, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todo(ctx context.Context, filter map[string]interface{}) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, filter map[string]interface{}, limit *int, page *int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, filter map[string]interface{}) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
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
	if len(searchError) > 0 {
		return nil, fmt.Errorf("%s", searchError)
	}
	if len(roles) > 0 {
		for _, r := range roles {
			res = append(res, r)
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

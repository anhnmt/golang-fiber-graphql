package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/xdorro/golang-fiber-project/ent"
	"github.com/xdorro/golang-fiber-project/graph/generated"
	"github.com/xdorro/golang-fiber-project/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserIn) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUserIn) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.DeleteUserResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, input *model.UsersIn) (*ent.UserConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Status(ctx context.Context, obj *ent.User) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

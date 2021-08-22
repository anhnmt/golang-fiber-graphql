package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/xdorro/golang-fiber-base-project/ent"
	"github.com/xdorro/golang-fiber-base-project/graph/generated"
	"github.com/xdorro/golang-fiber-base-project/graph/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/jwt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	hashPassword, _ := jwt.HashPassword(input.Password)

	return ent.FromContext(ctx).User.Create().
		SetName(input.Name).
		SetEmail(input.Email).
		SetPassword(hashPassword).
		SetStatus(input.Status.String()).
		Save(ctx)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.DeleteUserResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, input *model.UsersIn) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, input.PageIn.After, input.PageIn.First, input.PageIn.Before, input.PageIn.Last,
			ent.WithUserOrder(input.OrderBy),
			ent.WithUserFilter(input.Where.Filter),
		)
}

func (r *userResolver) Status(ctx context.Context, obj *ent.User) (model.UserStatus, error) {
	return model.UserStatus(obj.Status), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

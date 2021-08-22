package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/xdorro/golang-fiber-base-project/graph/handler/auth"

	"github.com/xdorro/golang-fiber-base-project/ent"
	"github.com/xdorro/golang-fiber-base-project/graph/model"
)

func (r *mutationResolver) Token(ctx context.Context, input model.LoginInput) (*model.Token, error) {
	return auth.Authenticate(ctx, &input)
}

func (r *mutationResolver) Refresh(ctx context.Context) (*model.Token, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

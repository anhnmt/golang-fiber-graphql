package auth

import (
	"context"
	"errors"
	"github.com/xdorro/golang-fiber-base-project/ent"
	"github.com/xdorro/golang-fiber-base-project/ent/user"
	"github.com/xdorro/golang-fiber-base-project/graph/model"
	"github.com/xdorro/golang-fiber-base-project/pkg/jwt"
)

func Authenticate(ctx context.Context, input *model.LoginInput) (*model.Token, error) {
	currentUser, err := ent.FromContext(ctx).User.Query().
		Where(user.EmailEQ(input.Email)).
		Where(user.StatusNotIn(model.UserStatusDeleted.String())).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("User is not found")
		}

		return nil, errors.New("User is exist")
	}

	if currentUser.Status == model.UserStatusDisabled.String() {
		return nil, errors.New("User is block")
	}

	if ok := jwt.VerifyPassword(input.Password, currentUser.Password); ok {
		return jwt.GenerateToken(currentUser)
	}

	return nil, errors.New("Password wrong")
}

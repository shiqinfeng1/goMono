package auth

import (
	"context"

	"github.com/pkg/errors"
)

type User struct {
	UUID  string
	Email string
	Role  string

	DisplayName string
}

type ctxKey int

const (
	userContextKey ctxKey = iota
)

var (
	// if we expect that the user of the function may be interested with concrete error,
	// it's a good idea to provide variable with this error
	ErrNoUserInContext = errors.New("no user in context")
)

func UserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(userContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, ErrNoUserInContext
}

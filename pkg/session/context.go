package session

import (
	"context"
	"errors"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func CreateSessionContext(ctx context.Context, token jwt.Token) (context.Context, error) {
	userID, username, err := ExtractTokenInformation(token)
	if err != nil {
		return nil, errors.New("failed to extract token information")
	}
	ctx = context.WithValue(ctx, "userID", userID)
	ctx = context.WithValue(ctx, "username", username)
	return ctx, nil
}

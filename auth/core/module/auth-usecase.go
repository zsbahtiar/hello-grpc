package module

import (
	"context"
	"fmt"
	"github.com/zsbahtiar/hello-grpc/auth/core/entity"
)

type auth struct {
}

type AuthImpl interface {
	ValidateToken(ctx context.Context, req entity.ValidateTokenRequest) (entity.ValidationTokenResponse, error)
}

func NewAuthUsecase() AuthImpl {
	return &auth{}
}

func (a *auth) ValidateToken(ctx context.Context, req entity.ValidateTokenRequest) (entity.ValidationTokenResponse, error) {
	if len(req.Token) < 1 {
		return entity.ValidationTokenResponse{}, fmt.Errorf("authorization is required")
	}

	return entity.ValidationTokenResponse{
		Name:  "bob",
		Email: "bob@mail.com",
	}, nil
}

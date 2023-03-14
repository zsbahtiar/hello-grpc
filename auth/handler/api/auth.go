package api

import (
	"context"
	"github.com/zsbahtiar/hello-grpc/auth/core/entity"
	"github.com/zsbahtiar/hello-grpc/auth/core/module"
	pbAuth "github.com/zsbahtiar/hello-grpc/auth/handler/api/grpc/pb/auth"
	"github.com/zsbahtiar/hello-grpc/auth/pkg/helper"
	"google.golang.org/protobuf/types/known/emptypb"
)

type auth struct {
	authUsecase module.AuthImpl
}

func NewAuthHandler(authUsecase module.AuthImpl) pbAuth.AuthServer {
	return &auth{authUsecase: authUsecase}
}

func (a *auth) ValidateToken(ctx context.Context, empty *emptypb.Empty) (*pbAuth.ValidateTokenResponse, error) {
	myCtx, _ := helper.FromContext(ctx)

	validated, err := a.authUsecase.ValidateToken(ctx, entity.ValidateTokenRequest{Token: myCtx.Token})
	if err != nil {
		return nil, err
	}
	return &pbAuth.ValidateTokenResponse{
		Name:  validated.Name,
		Email: validated.Email,
	}, nil
}

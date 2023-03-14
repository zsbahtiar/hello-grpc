package mycontext

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type MyContext struct {
	Authorization string
	UserID        string
}

func FromContext(ctx context.Context) (MyContext, bool) {
	myContext := MyContext{}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return MyContext{}, false
	}
	authorization := md.Get("authorization")
	if len(authorization) > 0 {
		myContext.Authorization = authorization[0]
	}

	userID := md.Get("userid")
	if len(userID) > 0 {
		myContext.UserID = userID[0]
	}

	return myContext, ok
}

package helper

import (
	"context"
	md "google.golang.org/grpc/metadata"
)

type Context struct {
	Token string
}

func FromContext(ctx context.Context) (Context, bool) {
	metadata := Context{}
	meta, ok := md.FromIncomingContext(ctx)
	if !ok {
		return metadata, false
	}
	token := meta.Get("authorization")
	if len(token) > 0 {
		metadata.Token = token[0]
	}
	return metadata, ok
}

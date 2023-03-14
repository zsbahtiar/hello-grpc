package mycontext

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"testing"
)

func Test_FromContext(t *testing.T) {
	type want struct {
		myContext MyContext
		ok        bool
	}
	tests := []struct {
		name string
		want want
		req  context.Context
	}{
		{
			name: "should return not ok when from metadata is not ok",
			want: want{},
			req:  context.Background(),
		},
		{
			name: "should return ok when from metadata is ok but my context is empty",
			want: want{
				ok: true,
			},
			req: metadata.NewIncomingContext(context.Background(), metadata.New(nil)),
		},
		{
			name: "should return my context and ok",
			want: want{
				ok: true,
				myContext: MyContext{
					Authorization: "Bearer uuid",
					UserID:        "uuid",
				},
			},
			req: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
				"authorization": "Bearer uuid",
				"userid":        "uuid",
			})),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, ok := FromContext(test.req)
			assert.Equal(t, test.want.myContext, got)
			assert.Equal(t, test.want.ok, ok)
		})
	}
}

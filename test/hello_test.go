package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yukpiz/grpc-e2e-example/pb"
)

func Test_Hello(t *testing.T) {

	t.Run("test name", func(t *testing.T) {
		res, err := svc.Hello(context.Background(), &pb.HelloRequest{
			Id:   1,
			Name: "yukpiz",
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}

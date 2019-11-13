package test

import (
	"os"
	"testing"

	"github.com/yukpiz/grpc-e2e-example/service"
)

var svc *service.Service

func TestMain(m *testing.M) {

	svc = service.NewService()
	os.Exit(m.Run())
}

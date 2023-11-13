package hello

import (
	"context"
	"testing"

	"github.com/luyasr/hello-world/pkg/ioc"
)

var (
	svc Service
	ctx context.Context
)

func init() {
	_ = ioc.Controller().Init()
	svc = ioc.Controller().Get(Name).(Service)
}

func TestSayHello(t *testing.T) {
	in := &HelloRequest{
		Name: "luya",
	}
	msg, err := svc.SayHello(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(msg)
}

package registry

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"
	kotobakov1 "github.com/kotonohako/all-in/backend/generated/buf/kotobako/v1"
	"github.com/kotonohako/all-in/backend/generated/buf/kotobako/v1/kotobakov1connect"
)

var _ kotobakov1connect.KotobakoServiceHandler = &KotobakoRegistry{}

type KotobakoRegistry struct{}

func (s *KotobakoRegistry) Health(context.Context, *connect_go.Request[kotobakov1.HealthRequest]) (*connect_go.Response[kotobakov1.HealthResponse], error) {
	return connect_go.NewResponse(&kotobakov1.HealthResponse{
		Status: "Hello, Mr Kotobako",
	}), nil
}

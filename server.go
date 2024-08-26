package main

import (
	"context"

	"connectrpc.com/connect"
	apiv1beta1 "github.com/lemon-mint/llm-translation/api/v1beta1"
	"github.com/lemon-mint/llm-translation/api/v1beta1/apiv1beta1connect"
)

type Server struct {
	apiv1beta1connect.UnimplementedTranslationServiceHandler
}

func (s *Server) Healthz(
	ctx context.Context,
	req *connect.Request[apiv1beta1.HealthzRequest],
) (*connect.Response[apiv1beta1.HealthzResponse], error) {
	return connect.NewResponse(&apiv1beta1.HealthzResponse{
		Healthy: true,
		Message: "ok",
		Version: "v1beta1",
	}), nil
}

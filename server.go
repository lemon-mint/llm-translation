package main

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/lemon-mint/coord"
	"github.com/lemon-mint/coord/llm"
	"github.com/lemon-mint/coord/pconf"
	apiv1beta1 "github.com/lemon-mint/llm-translation/api/v1beta1"
	"github.com/lemon-mint/llm-translation/api/v1beta1/apiv1beta1connect"

	_ "github.com/lemon-mint/coord/provider/aistudio"
	_ "github.com/lemon-mint/coord/provider/anthropic"
	_ "github.com/lemon-mint/coord/provider/openai"
	_ "github.com/lemon-mint/coord/provider/vertexai"
)

type Server struct {
	Model llm.Model

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

var (
	ErrNoModelAvailable   = errors.New("no model available")
	ErrUnknownProvider    = errors.New("unknown provider")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrNoInputText        = errors.New("no input text")
	ErrInvalidLanguage    = errors.New("invalid language")
)

func (s *Server) Translate(
	ctx context.Context,
	req *connect.Request[apiv1beta1.TranslateRequest],
) (*connect.Response[apiv1beta1.TranslateResponse], error) {
	if req.Msg.Options == nil && s.Model == nil {
		return nil, ErrNoModelAvailable
	}

	var model llm.Model
	if req.Msg.Options != nil {
		if req.Msg.Options.Credentials == nil {
			return nil, ErrInvalidCredentials
		}

		provider := ""
		var options []pconf.Config

		switch req.Msg.Options.Provider {
		case apiv1beta1.LLMOptions_OPENAI:
			provider = "openai"
			options = append(options, pconf.WithAPIKey(req.Msg.Options.Credentials.ApiKey))
		case apiv1beta1.LLMOptions_ANTHROPIC:
			provider = "anthropic"
			options = append(options, pconf.WithAPIKey(req.Msg.Options.Credentials.ApiKey))
		case apiv1beta1.LLMOptions_AISTUDIO:
			provider = "aistudio"
			options = append(options, pconf.WithAPIKey(req.Msg.Options.Credentials.ApiKey))
		default:
			return nil, ErrUnknownProvider
		}

		client, err := coord.NewLLMClient(ctx, provider, options...)
		if err != nil {
			return nil, err
		}
		defer client.Close()

		model, err = client.NewLLM(req.Msg.Options.Model, nil)
		if err != nil {
			return nil, err
		}
		defer model.Close()
	} else {
		model = s.Model
	}

	text := strings.TrimSpace(req.Msg.Text)
	if text == "" {
		return nil, ErrNoInputText
	}

	from := apiv1beta1.Language_ENGLISH
	to := req.Msg.TargetLanguage
	if to == apiv1beta1.Language_UNSPECIFIED {
		return nil, ErrInvalidLanguage
	}

	translated, err := translate(ctx, model, text, from, to)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&apiv1beta1.TranslateResponse{
		Success:        true,
		TranslatedText: translated,
		TargetLanguage: to,
	}), nil
}

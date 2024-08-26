package main

import (
	"context"
	"embed"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lemon-mint/coord"
	"github.com/lemon-mint/coord/pconf"
	"github.com/lemon-mint/coord/provider"
	"github.com/lemon-mint/envaddr"
	"github.com/lemon-mint/llm-translation/api/v1beta1/apiv1beta1connect"
	"github.com/rs/zerolog/log"
	"gopkg.eu.org/envloader"

	_ "github.com/lemon-mint/coord/provider/aistudio"
	_ "github.com/lemon-mint/coord/provider/anthropic"
	_ "github.com/lemon-mint/coord/provider/openai"
	_ "github.com/lemon-mint/coord/provider/vertexai"
)

//go:generate buf generate

//go:embed web/dist/**
var webFS embed.FS

var staticFS fs.FS

type svelteFS struct {
	fs.FS
}

func (g *svelteFS) Open(name string) (fs.File, error) {
	f, err := g.FS.Open(name)
	if err != nil {
		return g.FS.Open("index.html")
	}
	return f, nil
}

func main() {
	envloader.LoadEnvFile(".env")

	ln, err := net.Listen("tcp", envaddr.Get(":14402"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	defer ln.Close()

	staticFS, err = fs.Sub(webFS, "web/dist")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get static fs, maybe web/dist is not a directory")
	}

	var client provider.LLMClient
	rpcServer := &Server{}

	switch os.Getenv("SERVER_PROVIDER") {
	case "aistudio", "anthropic", "openai":
		client, err = coord.NewLLMClient(context.Background(), os.Getenv("SERVER_PROVIDER"), pconf.WithAPIKey(os.Getenv("SERVER_API_KEY")))
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create llm client")
		}
		defer client.Close()

		model, err := client.NewLLM(os.Getenv("SERVER_MODEL"), nil)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create llm model")
		}
		defer model.Close()
		rpcServer.Model = model

	case "openrouter":
		client, err = coord.NewLLMClient(context.Background(), "openai", pconf.WithAPIKey(os.Getenv("SERVER_API_KEY")), pconf.WithBaseURL("https://openrouter.ai/api/v1"))
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create llm client")
		}
		defer client.Close()

		model, err := client.NewLLM(os.Getenv("SERVER_MODEL"), nil)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create llm model")
		}
		defer model.Close()
		rpcServer.Model = model

	case "vertexai":
		client, err = coord.NewLLMClient(context.Background(), "vertexai")
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create llm client")
		}
		defer client.Close()

		model, err := client.NewLLM(os.Getenv("SERVER_MODEL"), nil)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create llm model")
		}
		defer model.Close()
		rpcServer.Model = model
	}

	path, handler := apiv1beta1connect.NewTranslationServiceHandler(
		rpcServer,
	)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(&svelteFS{staticFS})))
	mux.Handle(path, handler)

	srv := &http.Server{
		Handler: mux,
	}
	go func() {
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to serve")
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-sigChan

	err = srv.Close()
	if err != nil {
		log.Error().Err(err).Msg("failed to close server")
	}
}

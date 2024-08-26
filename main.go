package main

import (
	"embed"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/lemon-mint/envaddr"
	"github.com/rs/zerolog/log"
	"gopkg.eu.org/envloader"
)

//go:generate buf generate

//go:embed web/dist
var webFS embed.FS

var staticFS fs.FS

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

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(staticFS)))

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

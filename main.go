package main

import (
	"net"

	"github.com/lemon-mint/envaddr"
	"github.com/rs/zerolog/log"
	"gopkg.eu.org/envloader"
)

//go:generate buf generate

func main() {
	envloader.LoadEnvFile(".env")

	ln, err := net.Listen("tcp", envaddr.Get(":14402"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	defer ln.Close()
}

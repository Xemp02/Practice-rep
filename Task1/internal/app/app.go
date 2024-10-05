package app

import (
	"github.com/Xemp02/Practice-rep/internal/clients"
	"github.com/Xemp02/Practice-rep/internal/services"
	"github.com/Xemp02/Practice-rep/internal/transport"
	"github.com/Xemp02/Practice-rep/repository"
)

const httpPort = ":8080"

func Run() {
	// Start BD
	BDFilm := repository.New()
	// Start clients
	client := clients.New()
	// Start services
	service := services.New(client, BDFilm)
	// Start http server
	httpServer := transport.New(httpPort, service)

	httpServer.AddRoute()
	httpServer.StartHttpServer()
}

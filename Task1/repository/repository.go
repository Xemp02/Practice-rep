package repository

import (
	"github.com/Xemp02/Practice-rep/internal/models"
)

type BDFilm struct {
	db []models.Film
}

func New() *BDFilm {
	return &BDFilm{
		db: []models.Film{},
	}
}

/*func New(address string, service Service) *HttpServer {
	echoServer := echo.New()

	return &HttpServer{
		address: address,
		Echo:    echoServer,
		service: service,
	}
}
*/

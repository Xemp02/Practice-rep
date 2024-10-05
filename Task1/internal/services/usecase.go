package services

import (
	"encoding/json"
	"fmt"
	"github.com/Xemp02/Practice-rep/internal/models"
	"log"
	rand2 "math/rand/v2"
	"net/http"
)

type Client interface {
	GetRequest(url string) (*http.Response, error)
}
type BDFilm interface {
	GetFilm(number int) (*models.Film, error)
	CreateFilm(filmBD models.Film) error
}

func (s *Service) GetRandomFilm() (*models.Film, error) {
	id := rand2.IntN(100) + 300

	response, err := s.client.GetRequest(fmt.Sprintf("https://api.kinopoisk.dev/v1.4/movie/%v", id))
	if err != nil {
		log.Println("GetRandomFilm: " + err.Error())

		return nil, err
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	// Обработать ответ
	var film models.Film
	err = json.NewDecoder(response.Body).Decode(&film)
	if err != nil {
		log.Println(" json.NewDecoder: " + err.Error())

		return nil, err
	}

	return &film, nil
}

package convertor

import "github.com/kenedyCO/Practice/internal/models"

func ToRandomFilmResponseFromFilm(film *models.Film) *models.RandomFilmResponse {
	return &models.RandomFilmResponse{
		Id:     film.Id,
		Name:   film.Name,
		Poster: film.Poster,
	}
}

func ToTrailerResponseFromFilm(film *models.Film) *models.Trailer {
	return &models.Trailer{
		Name: film.Name,
		URL:  film.Videos.Trailers[0].URL,
	}
}

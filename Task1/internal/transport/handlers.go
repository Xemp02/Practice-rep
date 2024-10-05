package transport

import (
	"fmt"
	"github.com/Xemp02/Practice-rep/internal/convertor"
	"github.com/Xemp02/Practice-rep/internal/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetRandomFilm() (*models.Film, error)
}

func (h *HttpServer) randFilm(c echo.Context) error {
	film, err := h.service.GetRandomFilm()
	if err != nil {
		log.Println("h.service.GetRandomFilm: " + err.Error())

		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}

	resp := convertor.ToRandomFilmResponseFromFilm(film)
	// Backend end. Start Frontend.
	response := fmt.Sprintf(`
		<html>
			<head>
				<title> Рандомный фильм из Кинопоиска </title>
			</head>
			<body>
				<h1>%v</h1>
				<p>
					Ссылка на фильм: 
					<a href="https://www.kinopoisk.ru/film/%v">%v</a>
				</p>
				<img src="%v" alt="%v" />
			</body>
		</html>
	`, resp.Name, resp.Id, resp.Name, resp.Poster.URL, resp.Name)

	return c.HTML(http.StatusOK, response)
}

func (h *HttpServer) getTrailer(c echo.Context) error {
	film, err := h.service.GetRandomFilm()

	if err != nil {
		log.Println("h.service.getTrailer: " + err.Error())

		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}

	resp := convertor.ToTrailerResponseFromFilm(film)
	var trailerHTML string
	if len(film.Videos.Trailers) > 0 {
		trailerHTML = resp.URL
	}
	response := fmt.Sprintf(`
		<h3></h3>
      <iframe width="560" height="315" src="%v" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
	`, trailerHTML)

	return c.HTML(http.StatusOK, response)
}

/* <html>
	<head>
		<title> Рандомный фильм из Кинопоиска </title>
	</head>
	<body>
	<video>
		<source src="%v">
	<video>
	</body>
</html>
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
)

func main() {
	httpServer := echo.New()

	httpServer.GET("/", randFilm)

	httpServer.Logger.Fatal(httpServer.Start(":8080"))
}

type Rating struct {
	KP          float64 `json:"kp"`
	Filmcritics float64 `json:"filmCritics"`
	// RusFilmcritics float64 `json:"russianFilmCritics"`
}

type persons struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Photo       string `json:"photo"`
	EnName      string `json:"enName"`
	Description string `json:"description"`
}

type Film struct {
	Id      int        `json:"id"`
	Name    string     `json:"name"`
	Year    int        `json:"year"`
	Rating  Rating     `json:"rating"`
	Persons [1]persons `json:"persons"`
}

func randFilm(c echo.Context) error {
	// Создать случайный id от 300 до 400
	id := rand.IntN(400-300) + 300
	// Создать запрос к апи вместе с токеном
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.kinopoisk.dev/v1.4/movie/%v", id), nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-API-KEY", "VN41HVS-SHKMMZK-NZJ9TXW-20TBET7")

	// Создать http клиента и выполнить запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err - ", err)
		return err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	// Обработать ответ
	var film Film

	err = json.NewDecoder(resp.Body).Decode(&film)

	if err != nil {
		return err
	}

	fmt.Println(film)

	response := fmt.Sprintf("RandFilm - %v\nСсылка - https://www.kinopoisk.ru/film/%v\nГод выпуска - %v\nРейтинг Кинопоиска %v\nРейтинг критиков - %v\n Актёры %v", film.Name, film.Id, film.Year, RoundB(film.Rating.KP, 2), film.Rating.Filmcritics, film.Persons[0])

	return c.String(http.StatusOK, response)
}

func RoundB(a float64, b float64) float64 {
	// func math.Pow = ^ ; b is number power , a is the basic power
	a = math.Round(a*math.Pow(10, b)) / math.Pow(10, b)
	return a
}

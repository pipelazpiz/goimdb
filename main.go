package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/pipelazpiz/goimdb/getmovies"
)

func main() {
	getmovies.Data()
	e := echo.New()

	e.GET("/movies", getmovies.GetAllMoviesHandler)
	e.GET("/movies/:imdbID", getmovies.GetMoviesByIdHandler)

	e.POST("/movies", getmovies.CreateMoviesHandler)

	e.PUT("/movies/:imdbID", getmovies.UpdateMoviesHandlerID)

	port := "8888"
	log.Println("Starting Port:", port)
	err := e.Start("localhost:" + port)
	log.Fatal(err)
}

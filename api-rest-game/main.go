package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Jogador struct {
	Id     int    `json:"id"`
	Nome   string `json:"nome"`
	Online bool   `json:"online"`
}

type Jogadores []Jogador

var jogadores = Jogadores{Jogador{gerarID(), "GabiCigana", true}, Jogador{gerarID(), "Therafs", true}, Jogador{gerarID(), "Shoi", true}}

func gerarID() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10000)
}

func getJogadores(c echo.Context) error {
	return c.JSON(http.StatusOK, jogadores)
}

func postJogador(c echo.Context) error {
	jogador := Jogador{}
	err := c.Bind(&jogador)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	jogador.Id = gerarID()
	jogadores = append(jogadores, jogador)
	return c.JSON(http.StatusCreated, jogadores)
}
func getJogador(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, jogador := range jogadores {
		if jogador.Id == id {
			return c.JSON(http.StatusOK, jogador)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}
func putJogador(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range jogadores {
		if jogadores[i].Id == id {
			jogadores[i].Online = false
			return c.JSON(http.StatusOK, jogadores)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func deleteJogador(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range jogadores {
		if jogadores[i].Id == id {
			jogadores = append(jogadores[:i], jogadores[i+1:]...)
			return c.JSON(http.StatusOK, jogadores)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func main() {
	fmt.Println("Rodando...")
	e := echo.New()
	e.GET("/jogadores", getJogadores)
	e.POST("/jogadores", postJogador)
	e.GET("jogadores/:id", getJogador)
	e.PUT("jogadores/:id", putJogador)
	e.DELETE("jogadores/:id", deleteJogador)
	e.Start(":5555")

}

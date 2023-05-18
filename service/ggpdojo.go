package service

import (
	"net/http"

	"github.com/gin-gonic/gin"	
)

type GamesListResponse struct {
	Name string
	Roles []string
}

func listGames(ctx *gin.Context) {
    ctx.IndentedJSON(http.StatusOK,
			GamesListResponse{"tic-tac-toe", []string{"red", "blue"}})
}

func ServiceEventLoop() {
	router := gin.Default()
	router.GET("/games", listGames)

	router.Run("localhost:8080")
}
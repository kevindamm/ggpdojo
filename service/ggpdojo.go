package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GamesList struct {
	Games []GameMetadata
}

type GameMetadata struct {
	GameId string
	Name   string
	Tags   []string
	Roles  []string
}

func listGames(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK,
		GamesListResponse{
			"Tic-Tac-Toe",
			[]string{"Connection", "Simple"},
			[]string{"red", "blue"},
		})
}

type GameMeasures struct {
	Name     string
	Elegance float32
	Shibui   float32
}

func getGame(ctx *gin.Context) {
	var metadata GameMetadata
	
	ctx.IndentedJSON(http.StatusOK, metadata)
}

func ServiceEventLoop() {
	router := gin.Default()
	router.GET("/games", listGames)
	router.GET("/game/:gamename", getGame)

	// TODO read port from environment, config or flag
	router.Run(":80")
}

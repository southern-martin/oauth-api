package app

import (
	"github.com/gin-gonic/gin"
	"github.com/southern-martin/oauth-api/src/http"
	"github.com/southern-martin/oauth-api/src/repository/db"
	"github.com/southern-martin/oauth-api/src/repository/rest"
	"github.com/southern-martin/oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(
		access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}

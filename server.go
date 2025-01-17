package main

import (
	"postergist-api/init/database"
	postHttp "postergist-api/src/delivery/http/post"
	postRepo "postergist-api/src/repository/post"
	postUc "postergist-api/src/usecase/post"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status   int         `json:"status"`
	Response interface{} `json:"response"`
}

func main() {
	e := echo.New()

	db := database.SetupDatabase()

	postRepo := postRepo.NewPostRepository(db)
	postUC := postUc.NewPostUsecase(postRepo)
	postHttp.NewPostHTTPHandler(e, postUC)

	e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	postHttp "postergist-api/src/delivery/http"
	"postergist-api/src/repository"
	"postergist-api/src/usecase"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status   int         `json:"status"`
	Response interface{} `json:"response"`
}

func main() {
	e := echo.New()

	postRepo := repository.NewPostRepository()
	postUC := usecase.NewPostUsecase(postRepo)
	postHttp.NewPostHTTPHandler(e, postUC)

	e.Logger.Fatal(e.Start(":1323"))
}

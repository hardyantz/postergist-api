package main

import (
	"postergist-api/init/database"
	postHttp "postergist-api/src/post/delivery/http"
	postRepos "postergist-api/src/post/repository"
	postUc "postergist-api/src/post/usecase"

	catHttp "postergist-api/src/category/delivery/http"
	catRepos "postergist-api/src/category/repository"
	catUc "postergist-api/src/category/usecase"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status   int         `json:"status"`
	Response interface{} `json:"response"`
}

func main() {
	e := echo.New()

	db := database.SetupDatabase()

	postRepo := postRepos.NewPostRepository(db)
	postUC := postUc.NewPostUsecase(postRepo)
	postHttp.NewPostHTTPHandler(e, postUC)

	catRepo := catRepos.NewCategoryRepository(db)
	catUC := catUc.NewPostUsecase(catRepo)
	catHttp.NewCategoryHTTPHandler(e, catUC)

	e.Logger.Fatal(e.Start(":1323"))
}

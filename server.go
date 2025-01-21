package main

import (
	"postergist-api/init/database"
	postHttp "postergist-api/src/delivery/http/post"
	postRepos "postergist-api/src/repository/post"
	postUc "postergist-api/src/usecase/post"

	catHttp "postergist-api/src/delivery/http/category"
	catRepos "postergist-api/src/repository/category"
	catUc "postergist-api/src/usecase/category"

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
	catHttp.NewPostHTTPHandler(e, catUC)

	e.Logger.Fatal(e.Start(":1323"))
}

package post

import (
	"net/http"
	usecase "postergist-api/src/usecase/post"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	PostUsecase usecase.PostUc
}

func NewPostHTTPHandler(e *echo.Echo, p usecase.PostUc) {
	handler := &PostHandler{
		PostUsecase: p,
	}
	e.GET("/posts", handler.GetPost)
}

func (p *PostHandler) GetPost(c echo.Context) error {
	results, err := p.PostUsecase.GetPosts()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"response": "data not found"})
	}
	return c.JSON(http.StatusOK, results)
}

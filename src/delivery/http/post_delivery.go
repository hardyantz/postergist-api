package http

import (
	"net/http"
	"postergist-api/src/usecase"

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
	return c.JSON(http.StatusOK, []string{"response OK"})
}

package post

import (
	"net/http"
	domain "postergist-api/src/post/domain"
	usecase "postergist-api/src/post/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	PostUsecase usecase.PostUc
}

func NewPostHTTPHandler(e *echo.Echo, p usecase.PostUc) {
	handler := &PostHandler{
		PostUsecase: p,
	}
	e.GET("/posts", handler.GetPosts)
	e.POST("/post", handler.CreatePost)
	e.GET("/post/:id", handler.GetPost)
}

func (p *PostHandler) GetPosts(c echo.Context) error {
	results, err := p.PostUsecase.GetPosts()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"response": "data not found"})
	}
	return c.JSON(http.StatusOK, results)
}

func (p *PostHandler) CreatePost(c echo.Context) error {
	postData := new(domain.Post)
	if err := c.Bind(postData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"response": "invalid payload"})
	}
	if err := p.PostUsecase.CreatePost(*postData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"response": "failed to save data, with error: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"response": "data successfully created"})
}

func (p *PostHandler) GetPost(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"response": "id not found"})
	}

	results, err := p.PostUsecase.GetPost(postId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"response": "data not found"})
	}
	return c.JSON(http.StatusOK, results)
}

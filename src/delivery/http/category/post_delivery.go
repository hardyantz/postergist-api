package post

import (
	"net/http"
	domain "postergist-api/src/domain/category"
	usecase "postergist-api/src/usecase/category"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CategoryUsecase usecase.CategoryUc
}

func NewPostHTTPHandler(e *echo.Echo, cu usecase.CategoryUc) {
	handler := &CategoryHandler{
		CategoryUsecase: cu,
	}
	e.GET("/categories", handler.GetPost)
	e.POST("/category", handler.CreatePost)
}

func (ch *CategoryHandler) GetPost(c echo.Context) error {
	results, err := ch.CategoryUsecase.GetCategories()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"response": "data not found"})
	}
	return c.JSON(http.StatusOK, results)
}

func (ch *CategoryHandler) CreatePost(c echo.Context) error {
	categoryData := new(domain.Category)
	if err := c.Bind(categoryData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"response": "invalid payload"})
	}
	if err := ch.CategoryUsecase.CreateCategory(*categoryData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"response": "failed to save data, with error: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"response": "data successfully created"})
}

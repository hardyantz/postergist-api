package post

import (
	"net/http"
	domain "postergist-api/src/category/domain"
	usecase "postergist-api/src/category/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CategoryUsecase usecase.CategoryUc
}

func NewCategoryHTTPHandler(e *echo.Echo, cu usecase.CategoryUc) {
	handler := &CategoryHandler{
		CategoryUsecase: cu,
	}
	e.GET("/categories", handler.GetCategories)
	e.POST("/category", handler.CreateCategory)
	e.GET("/category/:id", handler.GetCategory)
}

func (ch *CategoryHandler) GetCategories(c echo.Context) error {
	results, err := ch.CategoryUsecase.GetCategories()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"response": "data not found"})
	}
	return c.JSON(http.StatusOK, results)
}

func (ch *CategoryHandler) CreateCategory(c echo.Context) error {
	categoryData := new(domain.Category)
	if err := c.Bind(categoryData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"response": "invalid payload"})
	}
	if err := ch.CategoryUsecase.CreateCategory(*categoryData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"response": "failed to save data, with error: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"response": "data successfully created"})
}

func (p *CategoryHandler) GetCategory(c echo.Context) error {
	catId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"response": "id not found"})
	}

	results, err := p.CategoryUsecase.GetCategory(catId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"response": "data not found"})
	}
	return c.JSON(http.StatusOK, results)
}

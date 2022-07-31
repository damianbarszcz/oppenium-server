package controllers

import (
	"net/http"

	"go-server/models"

	"github.com/labstack/echo/v4"
)

func GetRecommendProducts(c echo.Context) error {
	var products []models.Product
	Db.Find(&products, "availability = ?", "true")

	return c.JSON(http.StatusOK, products)
}

func GetComputersProducts(c echo.Context) error {
	var products []models.Product
	Db.Find(&products, "category = ?", "computers")

	return c.JSON(http.StatusOK, products)
}

func GetComputerComponentsProducts(c echo.Context) error {
	var products []models.Product
	Db.Find(&products, "category = ?", "computer-components")

	return c.JSON(http.StatusOK, products)
}

func GetSmartphonesProducts(c echo.Context) error {
	var products []models.Product
	Db.Find(&products, "category = ?", "smartphones")

	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	product := new(models.Product)
	Db.Where("id = ?", id).First(&product)
	return c.JSON(http.StatusOK, product)
}

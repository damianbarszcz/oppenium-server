package controllers

import (
	"go-server/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddProductToCart(c echo.Context) error {
	id := c.Param("user_id")

	product_id := c.FormValue("product_id")

	var product models.Product

	Db.Where("id = ?", product_id).First(&product)

	var user models.User

	Db.Where("id = ?", id).First(&user)

	cart := models.Cart{
		User_id:       user.Id,
		Product_price: product.Product_price,
		Product_title: product.Product_title,
		Product_image: product.Product_image,
	}

	Db.Create(&cart)

	return c.JSON(http.StatusCreated, "The product "+product.Product_title+" has been added to the cart.")
}

func GetShoppingCart(c echo.Context) error {
	id := c.Param("user_id")

	var cart []models.Cart

	Db.Where("user_id = ?", id).Find(&cart)

	return c.JSON(http.StatusCreated, cart)
}

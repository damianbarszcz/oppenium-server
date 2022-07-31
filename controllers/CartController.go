package controllers

import (
	"go-server/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddProductToCart(c echo.Context) error {
	id := c.Param("id")

	product_id := c.Param("product-id")

	var product models.Product

	Db.Where("id = ?", product_id).First(&product)

	var user models.User

	Db.Where("id = ?", id).First(&user)

	cart := models.Cart{
		User_id:       user.Id,
		Product_id:    product.Id,
		Product_price: product.Product_price,
		Product_count: 1,
		Product_title: product.Product_title,
		Product_image: product.Product_image,
	}

	Db.Create(&cart)

	return c.JSON(http.StatusCreated, "Produkt został dodany do koszyka.")
}

func GetShoppingCart(c echo.Context) error {
	id := c.Param("id")

	var cart []models.Cart

	Db.Joins("JOIN products ON products.id = carts.product_id").
		Where("carts.user_id=?", id).
		Group("carts.id").
		Find(&cart)

	return c.JSON(http.StatusCreated, cart)
}

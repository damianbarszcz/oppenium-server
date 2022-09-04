package controllers

import (
	"go-server/models"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

const USER_ID = "user_id = ?"

func makeOrderId(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func MakePayment(c echo.Context) error {
	userId := c.Param("user_id")

	var cart []models.Cart

	Db.Where(USER_ID, userId).Find(&cart)

	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(cart[0].Product_title),
					},
					UnitAmount: stripe.Int64(cart[0].Product_price * 100),
				},
				Quantity: stripe.Int64(1),
			},

			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(cart[1].Product_title),
					},
					UnitAmount: stripe.Int64(cart[1].Product_price * 100),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String("http://localhost:3000/member/order/success"),
		CancelURL:  stripe.String("http://localhost:3000/member/order/cancel"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	orderId := makeOrderId(10)

	currentTime := time.Now()

	for _, single_cart := range cart {
		order := models.Order{
			User_id:        single_cart.User_id,
			Order_id:       orderId,
			Product_title:  single_cart.Product_title,
			Product_image:  single_cart.Product_image,
			Product_price:  single_cart.Product_price,
			Date:           currentTime,
			Payment_status: true,
		}

		Db.Create(&order)
	}

	Db.Where(USER_ID, userId).Delete(&cart)

	return c.JSON(http.StatusOK, s)
}

func GetProductsOrder(c echo.Context) error {
	userId := c.Param("user_id")

	var order []models.Order

	Db.Where(USER_ID, userId).Find(&order)

	return c.JSON(http.StatusCreated, order)
}

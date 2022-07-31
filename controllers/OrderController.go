package controllers

import (
	"go-server/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func MakePayment(c echo.Context) error {
	id := c.Param("id")

	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("http://localhost:3000/member/order/success"),
		CancelURL:  stripe.String("http://localhost:3000/member/order/cancel"),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String("price_1LRZZAFtCHoZJI1YwTWwyGWR"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
	}
	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	var sum uint

	Db.Table("carts").Select("sum(product_price)").Where("user_id = ?", id).Row().Scan(&sum)

	var cart []models.Cart

	Db.Where("user_id = ?", id).Find(&cart)

	order := models.Order{
		User_id:     cart[0].User_id,
		Total_price: sum,
		Is_paid:     true,
	}

	Db.Create(&order)

	Db.Where("user_id = ?", id).Delete(&cart)

	return c.JSON(http.StatusOK, s)
}

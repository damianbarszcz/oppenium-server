package route

import (
	"fmt"
	"go-server/controllers"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	g.Static("/public", "public")

	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("%s Backend 0.0.1", os.Getenv("APP")))
	})

	g.GET("/product/recommend", controllers.GetRecommendProducts)
	g.GET("/product/computers", controllers.GetComputersProducts)
	g.GET("/product/computer-components", controllers.GetComputerComponentsProducts)
	g.GET("/product/smartphones", controllers.GetSmartphonesProducts)
	g.GET("/product/single/:id", controllers.GetProduct)

	g.POST("/user/register", controllers.UserRegister)
	g.POST("/user/login", controllers.UserLogin)
	g.GET("/user/data", controllers.GetUser)
	g.POST("/user/logout", controllers.GetLogout)

	g.POST("/cart/:id/add-cart/:product-id", controllers.AddProductToCart)
	g.GET("/cart/:id/products", controllers.GetShoppingCart)

	g.GET("/order/make-peyment/:id", controllers.MakePayment)
}

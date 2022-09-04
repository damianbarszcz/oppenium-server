package main

import (
	"go-server/db"
	"go-server/route"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stripe/stripe-go/v72"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"http://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
)

func main() {
	stripe.Key = "sk_test_51LPYVbFtCHoZJI1YsshBLAVv1S4CyqnDpA0IVZlTUWWbsIlSaMISuu56qNifiYyeg3QiZK70DDMRrCXf6GoD8Aen00qDWndSKT"

	godotenv.Load()

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	db.Init()

	route.Init(app.Group("/api"))

	app.Logger.Fatal(app.Start(":" + os.Getenv("PORT")))
}

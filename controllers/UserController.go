package controllers

import (
	"go-server/models"
	"go-server/models/dto"
	"net/http"

	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func UserRegister(c echo.Context) error {

	var u = dto.UserDTO{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	var ud = dto.UserDataDTO{
		First_name:   c.FormValue("first_name"),
		Last_name:    c.FormValue("last_name"),
		Phone_number: c.FormValue("phone_number"),
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		panic(err)
	}

	var user models.User

	Db.Where("email = ?", u.Email).First(&user)

	if user.Id > 0 {
		return c.JSON(http.StatusConflict, "The user with this email already exists.")
	}

	user_register := models.User{
		Email:    u.Email,
		Password: string(hashPassword),
	}

	Db.Create(&user_register)

	Db.Where("email = ?", u.Email).First(&user)

	user_data_register := models.UserData{
		User_id:      user.Id,
		First_name:   ud.First_name,
		Last_name:    ud.Last_name,
		Phone_number: ud.Phone_number,
	}

	Db.Create(&user_data_register)

	return c.JSON(http.StatusCreated, "The user was created.")
}

func UserLogin(c echo.Context) error {

	var u = dto.UserDTO{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	var user models.User

	Db.Where("email = ?", u.Email).First(&user)

	if user.Id == 0 {
		return c.JSON(http.StatusUnauthorized, "Incorrect email address or password.")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, "Incorrect email address or password.")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day

	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.SetCookie(cookie)

	return c.String(http.StatusOK, token)
}

func GetUser(c echo.Context) error {
	cookie, err := c.Cookie("jwt")

	if err != nil {
		return err
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return echo.ErrUnauthorized
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	Db.Where("id = ?", claims.Issuer).Preload("UserData").First(&user)

	return c.JSON(http.StatusOK, user)
}

func GetLogout(c echo.Context) error {

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour)

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "success")
}

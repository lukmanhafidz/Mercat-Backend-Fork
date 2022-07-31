package common

import (
	"Estore/config"
	"Estore/domain"
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(userdata domain.User) string {
	info := jwt.MapClaims{}
	info["ID"] = userdata.ID
	info["Role"] = userdata.Role
	auth := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	token, err := auth.SignedString([]byte(config.SECRET))
	if err != nil {
		log.Fatal("cannot generate key")
		return ""
	}

	return token
}

func ExtractData(c echo.Context) domain.User {
	var userdata domain.User

	head := c.Request().Header
	token := strings.Split(head.Get("Authorization"), " ")

	res, _ := jwt.Parse(token[len(token)-1], func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})

	if res.Valid {
		resClaim := res.Claims.(jwt.MapClaims)
		userdata.ID = resClaim["ID"].(int)
		userdata.Role = resClaim["Role"].(string)
		return userdata
	}

	return domain.User{}
}

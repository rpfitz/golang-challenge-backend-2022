package controller

import (
	"backendmod/config"
	"backendmod/model"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var googleUser model.GoogleUser = model.NewGoogleUser()
var userModel model.User = model.NewUser()

func CreateToken(w http.ResponseWriter, email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    email,
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
	})

	token, err := claims.SignedString([]byte(config.SECRET))

	if err != nil {
		log.Println("CreateToken - SignedString Error:", err.Error())
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 30),
		HttpOnly: true,
	})

	return token, nil
}

func GetCookie(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		log.Println("GetCookie Error:", err.Error())
		return "", err
	}
	return cookie.Value, nil
}

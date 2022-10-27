package controller

import (
	"backendmod/config"
	"backendmod/types"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthUser(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	jwtCookie, _ := GetCookie(r, "Authorization")

	token, err := jwt.ParseWithClaims(jwtCookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(types.AuthUserFailedResponseType{
			Success: false,
			Message: "Unauthorized.",
		})
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, userExists := userModel.GetByEmail(claims.Issuer)

	if userExists {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(types.AuthUserResponseType{
			Success: true,
			User: types.AuthUserType{
				ID:        user.ID_User,
				Email:     user.Email,
				Full_Name: user.Full_Name,
				Telephone: user.Telephone,
				Password:  user.Password,
				Google:    user.Google,
			},
		})
		return
	}
}

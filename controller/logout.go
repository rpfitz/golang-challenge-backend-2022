package controller

import (
	"backendmod/types"
	"encoding/json"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})
	json.NewEncoder(w).Encode(types.LogoutResponseType{
		Success: true,
		Message: "You have expired your session.",
	})
	return
}

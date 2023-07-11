package controller

import (
	"encoding/json"
	"net/http"
	//"services/cookies"

	// "github.com/dgrijalva/jwt-go"

	"login/services"
)

var users = map[string]string{
	// key value pair
	// we can use db as well
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// type Claims struct {
// 	Username string `json:"username`
// 	jwt.StandardClaims
// }

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials                         // declaring variable name of type credential
	err := json.NewDecoder(r.Body).Decode(&credentials) // r.Body is basically request body , decode will decode the data and store in credential var
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.Username]  //matching username & password
	if !ok || expectedPassword != credentials.Password { // if data not available or password not match
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	services.SetCookie(w, r, credentials.Username)
	// if data match then set cookie
	// expirationTime := time.Now().Add(time.Minute * 5) // adding 5 minutes

	// claims := &Claims{
	// 	Username: credentials.Username,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: expirationTime.Unix(),
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //new jwt token is being created with signingmethodhs256
	// tokenString, err := token.SignedString(jwtKey)             //signed using jwt key
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })

}

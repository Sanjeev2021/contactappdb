package services

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "login/controller"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

// creating jwt key
var jwtKey = []byte("secret_key")

func SetCookie(w http.ResponseWriter, r *http.Request, username string) {

	expirationTime := time.Now().Add(time.Minute * 5) // adding 5 minutes
	// cookie := &Claims{
	// 	Username: credentials.Username,
	// }
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //new jwt token is being created with signingmethodhs256
	tokenString, err := token.SignedString(jwtKey)             //signed using jwt key
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

func isValidCoockie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}

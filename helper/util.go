package helper

import "github.com/dgrijalva/jwt-go"

type Error struct {
	ResponseCode      int    `json:"rc"`
	Message           string `json:"message"`
	Detail            string `json:"detail"`
	ExternalReference string `json:"ext_ref"`
}

func JwtGenerator(username, key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}
	return tokenString
}

func ErrorLog(rc int, detail, ext_ref string) Error {
	var error Error
	error.ResponseCode = rc
	error.Message = "Failed"
	error.Detail = detail
	error.ExternalReference = ext_ref

	return error
}

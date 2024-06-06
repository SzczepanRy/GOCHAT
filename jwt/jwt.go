package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func CreateToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims{
        "username": username,
        "exp": time.Now().Add(time.Hour /2).Unix(),
        })

    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenString string) (error){
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return []byte(os.Getenv("SECRET")), nil
   })

   if err != nil {
      return err
   }

   if !token.Valid {
      return fmt.Errorf("invalid token")
   }
   return nil
}

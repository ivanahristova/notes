package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Valid(c *gin.Context) error {
	tokenString := ExtractToken(c)

	_, err := jwt.Parse(tokenString, keyFunc)

	return err
}

func Generate(user_id uint, admin bool) (string, error) {
	lifespan, err := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ExtractToken(c *gin.Context) string {
	if token := c.Query("token"); token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	tokens := strings.Split(bearerToken, " ")

	if len(tokens) != 2 {
		return ""
	}

	return tokens[1]
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	token, err := jwt.Parse(ExtractToken(c), keyFunc)

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return 0, nil
	}

	var uid uint64

	uid, err = strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)

	if err != nil {
		return 0, err
	}

	return uint(uid), nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(os.Getenv("TOKEN_SECRET")), nil
}

func ExtractAdminRights(c *gin.Context) (bool, error) {
	tokenString := ExtractToken(c)

	tkn, err := jwt.Parse(tokenString, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tkn.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := tkn.Claims.(jwt.MapClaims)
	if ok && tkn.Valid {
		AdminRights, err := strconv.ParseBool(fmt.Sprint(claims["user_admin_rights"]))
		if err != nil {
			return false, err
		}
		return AdminRights, nil
	}

	return false, nil
}

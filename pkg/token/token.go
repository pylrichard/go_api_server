package token

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero")
)

// Context is the context of the jwt
type Context struct {
	ID       uint64
	UserName string
}

// secretFunc validates the secret format
func secretFunc(secret string) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}
	// Parse the token
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Read the token if it's valid
		ctx.ID = uint64(claims["id"].(float64))
		ctx.UserName = claims["user_name"].(string)

		return ctx, nil
	} else {
		return ctx, err
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")
	secret := viper.GetString("jwt_secret")
	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}
	var t string
	fmt.Sscanf(header, "Bearer %s", &t)

	return Parse(t, secret)
}

// Sign signs the context with the specified secret
func Sign(gctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	// Load the jwt secret from the gin config if the secret isn't specified
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	// The token content
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"id":        c.ID,
		"user_name": c.UserName,
		"nbf":       time.Now().Unix(),
		"iat":       time.Now().Unix(),
	})
	// Sign the token with the specified secret
	tokenString, err = token.SignedString([]byte(secret))

	return
}
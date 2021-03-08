package user

import (
	"go/tiny_http_server/handler"
	"go/tiny_http_server/model"
	"go/tiny_http_server/pkg/auth"
	"go/tiny_http_server/pkg/errno"
	"go/tiny_http_server/pkg/token"

	"github.com/gin-gonic/gin"
)

// Login generates the authentication token
// if the password was matched with the specified account
func Login(c *gin.Context)  {
	// Binding the data with the user struct
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Get the user information by the login user name
	d, err := model.GetUser(u.UserName)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user passowrd
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the jwt
	t, err := token.Sign(c, token.Context{ID: d.Id, UserName: d.UserName}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
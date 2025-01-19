package auth

import (
	"go-todo-app/base"
	"go-todo-app/utils/dtoutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *service
}

func (c *controller) login(ctx *gin.Context) {
	var request dtoutil.HasuraRequest[LoginRequest]

	if err := ctx.BindJSON(&request); err != nil {
		ctx.Error(&Errors.InvalidLoginRequest)
		return
	}

	res, err := c.service.login(request.Input.Params)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(base.NewApiMessage(http.StatusOK, res))
}

func (c *controller) verifyToken(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	splitedStrings := strings.Split(authHeader, " ")

	if len(splitedStrings) != 2 {
		ctx.Error(&Errors.UnAuthorized)
		return
	}

	authToken := splitedStrings[1]

	payload, err := c.service.verifyToken(authToken)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, VerifyTokenResponse{
		XHasuraUserId: payload.Id,
		XHasuraRole:   "user",
	})
}

func newController() *controller {
	service := NewService()
	controller := controller{service}

	return &controller
}

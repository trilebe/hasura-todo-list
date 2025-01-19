package user

import (
	"go-todo-app/base"
	"go-todo-app/utils/dtoutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *Service
}

func (c *controller) register(ctx *gin.Context) {
	var request dtoutil.HasuraRequest[RegisterUserRequest]

	if err := ctx.BindJSON(&request); err != nil {
		ctx.Error(&Errors.InvalidRegisterPayload)
		return
	}

	res, err := c.service.Register(request.Input.Params)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(base.NewApiMessage(http.StatusCreated, res))
}

func newController() *controller {
	service := NewService()
	controller := controller{service}

	return &controller
}

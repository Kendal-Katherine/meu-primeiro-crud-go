package controller

import (
	"net/http"

	"github.com/Kendal-Katherine/meu-primeiro-crud-go/src/configurantion/logger"
	"github.com/Kendal-Katherine/meu-primeiro-crud-go/src/configurantion/validation"
	"github.com/Kendal-Katherine/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/Kendal-Katherine/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zapcore.Field{
			Key:    "Journey",
			String: "createUser",
		},
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate info", err)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User created sucessfully",
		zapcore.Field{
			Key:    "Journey",
			String: "createUser",
		})
	c.JSON(http.StatusOK, response)
}

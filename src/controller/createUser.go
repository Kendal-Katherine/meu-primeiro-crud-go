package controller

import (
	"fmt"

	"github.com/Kendal-Katherine/meu-primeiro-crud-go/src/configurantion/rest_err"
	"github.com/Kendal-Katherine/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, erro=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}

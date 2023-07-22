package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) Profile(ginContext *gin.Context) {
	response := make(map[string]string)
	response["message"] = "You are logged in"
	ginContext.Header("Access-Control-Allow-Origin", "*")
	ginContext.Header("Access-Control-Allow-Headers", "*")
	ginContext.JSON(http.StatusOK, response)
	return
}

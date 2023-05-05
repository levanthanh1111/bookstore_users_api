package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vanthanh.com/bookstore_users_api/domain/users"
	"vanthanh.com/bookstore_users_api/services"
	"vanthanh.com/bookstore_users_api/utils/errors"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusOK, "get user")
}

func GetUserById(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errRes := errors.ResponseError("bad request")
		c.JSON(errRes.Status, errRes)
		return
	}
	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := errors.ResponseError("invalid")
		c.JSON(errRes.Status, errRes)
		return
	}
	result, getErr := services.CreateUser(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

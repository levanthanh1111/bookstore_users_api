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

func UpdateUser(c *gin.Context) {
	var user users.User
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errRes := errors.ResponseError("bad request")
		c.JSON(errRes.Status, errRes)
		return
	}
	user.Id = userId
	if err := c.ShouldBindJSON(&user); err != nil {
		errRes := errors.ResponseError("invalid data")
		c.JSON(errRes.Status, errRes)
		return
	}
	isPartial := c.Request.Method == http.MethodPatch
	result, getErr := services.UpdateUser(user, isPartial)
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

func DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errRes := errors.ResponseError("bad request")
		c.JSON(errRes.Status, errRes)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func FindByStatus(c *gin.Context) {
	status := c.Query("status")
	results, err := services.FindByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, results)
}

package users

import (
	"github.com/gin-gonic/gin"
	"github.com/restwzeasy/bookstore_users-api/pkg/api/errors"
	"net/http"
	"strconv"
)

type UserApi interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	SearchUser(c *gin.Context)
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	createdUser, saveErr := DoCreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := DoGetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me.")
}

package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nebojsaj1726/user-manager/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) Create(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user.ID = primitive.NewObjectID()

	if err := uc.UserUsecase.Create(c, &user); err != nil {
		if strings.Contains(err.Error(), "email must be unique") || strings.Contains(err.Error(), "age must be greater than 18") {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Something went wrong"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func (uc *UserController) Fetch(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid page parameter"})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid limit parameter"})
		return
	}

	users, err := uc.UserUsecase.Fetch(c, pageInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	total, err := uc.UserUsecase.Count(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"total": total,
	})
}

func (uc *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")

	objectID, valid := ValidateObjectID(c, id)
	if !valid {
		return
	}

	user, err := uc.UserUsecase.GetByID(c, objectID.Hex())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Update(c *gin.Context) {
	id := c.Param("id")

	objectID, valid := ValidateObjectID(c, id)
	if !valid {
		return
	}

	var user domain.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if err := uc.UserUsecase.Update(c, objectID.Hex(), &user); err != nil {
		if strings.Contains(err.Error(), "email must be unique") || strings.Contains(err.Error(), "age must be greater than 18") {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Something went wrong"})
		}
		return
	}

	updatedUser, err := uc.UserUsecase.GetByID(c, objectID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to retrieve updated user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}

func (uc *UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	objectID, valid := ValidateObjectID(c, id)
	if !valid {
		return
	}

	if err := uc.UserUsecase.Delete(c, objectID.Hex()); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "User deleted successfully"})
}

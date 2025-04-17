package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nebojsaj1726/user-manager/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateObjectID(c *gin.Context, id string) (primitive.ObjectID, bool) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid ID"})
		return primitive.NilObjectID, false
	}
	return objectID, true
}

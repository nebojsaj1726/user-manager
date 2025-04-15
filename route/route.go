package route

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/nebojsaj1726/user-manager/bootstrap"
	"github.com/nebojsaj1726/user-manager/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router *gin.Engine) {
	userGroup := router.Group("/users")
	NewUserRouter(env, timeout, db, userGroup)
}

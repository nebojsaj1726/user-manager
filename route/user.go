package route

import (
	"time"

	"github.com/gin-gonic/gin"

	// "github.com/nebojsaj1726/user-manager/api/controller"
	"github.com/nebojsaj1726/user-manager/bootstrap"
	// "github.com/nebojsaj1726/user-manager/domain"
	"github.com/nebojsaj1726/user-manager/mongo"
	// "github.com/nebojsaj1726/user-manager/repository"
	// "github.com/nebojsaj1726/user-manager/usecase"
)

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	// ur := repository.NewUserRepository(db, domain.CollectionUser)
	// controller := &controller.UserController{
	// 	UserUsecase: usecase.NewUserUseCase(ur, timeout),
	// }

	// group.GET("/", controller.Fetch)
	// group.POST("/", controller.Create)
	// group.GET("/:id", controller.GetByID)
	// group.PUT("/:id", controller.Update)
	// group.DELETE("/:id", controller.Delete)

	group.GET("/", func(c *gin.Context) {
		c.String(200, "User route is working!")
	})
}

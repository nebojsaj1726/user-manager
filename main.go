package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nebojsaj1726/user-manager/bootstrap"
	"github.com/nebojsaj1726/user-manager/route"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})
	route.Setup(env, timeout, db, router)

	addr := fmt.Sprintf("%s:%s", env.ServerHost, env.ServerPort)
	router.Run(addr)

}

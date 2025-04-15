package main

import (
	"fmt"
	"time"

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

	gin := gin.Default()
	gin.SetTrustedProxies([]string{"127.0.0.1"})

	route.Setup(env, timeout, db, gin)

	addr := fmt.Sprintf("%s:%s", env.ServerHost, env.ServerPort)
	gin.Run(addr)

}

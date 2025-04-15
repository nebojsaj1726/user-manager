package bootstrap

import (
	log "github.com/sirupsen/logrus"

	"github.com/nebojsaj1726/user-manager/mongo"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() Application {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)

	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)

	log.Infof("Server is now running at %s:%s", app.Env.ServerHost, app.Env.ServerPort)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}

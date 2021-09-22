package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {

	if err := server.Viper(); err != nil {
		logrus.Fatal(err)
	}

	c := container.New()
	server.Logrus()
	//Cfg provides configuration related providers to the container
	server.Cfg(c)
	//Mongo provides mongo client to the container
	server.Mongo(c)
	//MongoCollection provides constructor for all the collections of this project
	server.MongoCollections(c)
	//Redis Provides a constructor for creating redis client
	server.Redis(c)

	server.Store(c)
	server.Cache(c)

	server.Middleware(c)

	server.Handler(c)

	//All data
	server.Patient(c)
	server.Organization(c)

	server.WS(c)
	server.Event(c)

	server.Validator(c)

	c.Resolve(
		func(cfg cfg.Server, handler http.Handler) {
			logrus.Info("Starting server at port ", cfg.Port)
			http.ListenAndServe(":"+cfg.Port, handler)
		})
}

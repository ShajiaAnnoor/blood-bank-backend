package main

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/Aubichol/hrishi-backend/bootstrap/server"
	"gitlab.com/Aubichol/hrishi-backend/container"
	"gitlab.com/Aubichol/hrishi-backend/index"
)

func main() {
	if err := server.Viper(); err != nil {
		logrus.Fatal(err)
	}

	c := container.New()
	server.Cfg(c)
	server.Mongo(c)
	server.MongoCollections(c)

	index.Create(c)
}

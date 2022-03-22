package container

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

//defaultContainer encapsulates dig's registration logic
type defaultContainer struct {
	*dig.Container
}

//Register registers a provider
func (dc *defaultContainer) Register(provider interface{}) {
	if err := dc.Container.Provide(provider); err != nil {
		logrus.Fatal(err)
	}
}

//RegisterWithName registers a provider with name.
//This scenario is helpful when multiple providers give same service
func (dc *defaultContainer) RegisterWithName(provider interface{}, containerName string) {
	if err := dc.Container.Provide(provider, dig.Name(containerName)); err != nil {
		logrus.Fatal(err)
	}
}

func (dc *defaultContainer) Resolve(function interface{}) {
	if err := dc.Invoke(function); err != nil {
		logrus.Fatal(err)
	}
}

//RegisterGroup registers a provider that belongs to a group
func (dc *defaultContainer) RegisterGroup(provider interface{}, name string) {
	if err := dc.Provide(provider, dig.Group(name)); err != nil {
		logrus.Fatal(err)
	}
}

// New returns default Container
func New() Container {
	return &defaultContainer{dig.New()}
}

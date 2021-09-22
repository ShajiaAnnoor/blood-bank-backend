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
func (d *defaultContainer) Register(provider interface{}) {
	if err := d.Container.Provide(provider); err != nil {
		logrus.Fatal(err)
	}
}

//RegisterWithName registers a provider with name.
//This scenario is helpful when multiple providers give same service
func (d *defaultContainer) RegisterWithName(provider interface{}, containerName string) {
	if err := d.Container.Provide(provider, dig.Name(containerName)); err != nil {
		logrus.Fatal(err)
	}
}

func (d *defaultContainer) Resolve(function interface{}) {
	if err := d.Invoke(function); err != nil {
		logrus.Fatal(err)
	}
}

//RegisterGroup registers a provider that belongs to a group
func (d *defaultContainer) RegisterGroup(provider interface{}, name string) {
	if err := d.Provide(provider, dig.Group(name)); err != nil {
		logrus.Fatal(err)
	}
}

// New returns default Container
func New() Container {
	return &defaultContainer{dig.New()}
}

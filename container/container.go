package container

// Container provides a common interface for dig
type Container interface {
	Register(provider interface{})
	RegisterGroup(provider interface{}, name string)
	RegisterWithName(provider interface{}, containerName string)
	Resolve(function interface{})
}

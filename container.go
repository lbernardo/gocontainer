package gocontainer

type Container struct {
	modules map[string]interface{}
}

var container *Container

func init() {
	container = new(Container)
	container.modules = map[string]interface{}{}
}

func Register(name string, module interface{}) {
	container.modules[name] = module
}

func Get(name string) interface{} {
	return container.modules[name]
}

func Has(name string) bool {
	_, ok := container.modules[name]
	return ok
}

func Delete(name string) {
	delete(container.modules, name)
}

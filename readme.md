# gocontainer

Register modules and retrieve from other locations in your code

## Install

```shell
go get github.com/lbernardo/gocontainer
```

## Usage

### Register your implementation

```go
package main

type Repository interface {
	GetData() []string
}

type RepositoryInfra struct {}

func (r * RepositoryInfra) GetData() []string {
	return []string{"worked"}
}

func NewRepositoryInfra() *RepositoryInfra {
	return &RepositoryInfra{}
}

func main() {
	gocontainer.Register("repository_infra", NewRepositoryInfra())
	// .... in another location you can use: 
	repositoryInfra := gocontainer.Get("repository_infra").(Repository)
	fmt.Println(RepositoryInfra.GetData())
}
```

You can replace value for you tests

```go
package logic

import "github.com/lbernardo/gocontainer"

type Logic struct {
	repository Repository
}

func NewLogic() *Logic {
	return &Logic{
		repository: gocontainer.Get("repository_infra").(Repository),
	}
}

func (l *Logic) GetData() []string{
	return l.repository.GetData()
}

// In test:

func TestGetData() {
	gocontainer.Register("repository_infra", &MockRepository{})
	l = NewLogic()
	l.GetData()
	// ... Repository.GetData is mocked
}
```

## Another functions

- Has(string) : Check if exists register
- Delete(string): Delete register
- Register(string, interface{}): Register/Replace a implementation
- Get(name) interface{} : Get implementation by name
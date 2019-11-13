package external

import "fmt"

type ExternalInterface interface {
	Hello()
}

type external struct{}

func NewExternal() ExternalInterface {
	return &external{}
}

func (*external) Hello() {
	fmt.Println("external hello!")
}

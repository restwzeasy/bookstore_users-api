package users

import "fmt"

type ChildAApi interface {
	Notify(value string)
}

type ChildA struct {
	parent CRUDParent
}

func (c ChildA) Notify(value string) {
	fmt.Println("Notifying for value: " + value)
}




package users

import "fmt"

type CRUDApi interface {
	Save(value string)
	Delete(value string)
}

type CRUDParent struct{}

func (CRUDParent) Save(value string) {
	fmt.Println("Save method for input: " + value)
}

func (CRUDParent) Delete(value string) {
	fmt.Println("Delete method for input: " + value)
}


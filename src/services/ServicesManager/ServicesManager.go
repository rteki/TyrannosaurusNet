package main

import (
	"fmt"
	"interfaces"
)

type ServiceManager string

func (s ServiceManager) Run() {
	fmt.Print("Hello test")
}

func (s ServiceManager) Stop() {

}

func main() {}

func NewServiceManager() interfaces.Service {
	var s ServiceManager
	return s
}

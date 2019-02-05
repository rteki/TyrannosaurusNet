package main

import (
	"fmt"
	"interfaces"
)

type ServiceManager string

func (s ServiceManager) Run() {
	fmt.Println("Services Manager started.")
}

func (s ServiceManager) Stop() {
	fmt.Println("Services Manager stoped.")
}

func main() {}

func constructor() interfaces.Service {
	var s ServiceManager
	return s
}

var New interfaces.ServiceConstructor = constructor

package main

import (
	"fmt"
	"plugin"
)

func main() {
	plug, _ := plugin.Open("plugins/ServicesManager.so")
	sym, _ := plug.Lookup("Run")

	fmt.Println()

	sym.(func())()

}

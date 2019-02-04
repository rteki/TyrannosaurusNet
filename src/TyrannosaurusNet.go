package main

import (
	"plugin"
)

func main() {
	plug, _ := plugin.Open("plugins/ServicesManager.so")
	sym, _ := plug.Lookup("Run")
	sym.(func())()
}

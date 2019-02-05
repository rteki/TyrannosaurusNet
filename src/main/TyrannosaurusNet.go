package main

import (
	"interfaces"
	"plugin"
)

func main() {
	plug, _ := plugin.Open("plugins/ServicesManager.so")
	sim, _ := plug.Lookup("NewServiceManager")

	a := sim.(func() interfaces.Service)()

	a.Run()
}

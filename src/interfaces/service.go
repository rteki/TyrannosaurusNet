package interfaces

import (
	"plugin"
)

type Service interface {
	Run()
	Stop()
}

type ServiceConstructor func() Service

func LoadServiceFromPlugin(path string) Service {
	plug, _ := plugin.Open(path)
	sim, _ := plug.Lookup("New")

	New := sim.((*ServiceConstructor))

	return (*New)()
}

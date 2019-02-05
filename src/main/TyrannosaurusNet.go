package main

import (
	"interfaces"
)

func main() {
	s := interfaces.LoadServiceFromPlugin("./plugins/ServicesManager.so")

	s.Run()
	s.Stop()
}

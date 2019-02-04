ifndef PLUGINS_DIR
	export BUILD_DIR=$(CURDIR)/../build
	export PLUGINS_DIR=$(BUILD_DIR)/plugins
endif

all: services main
	mv TyrannosaurusNet $(BUILD_DIR)/

services: services/Makefile
	cd services; make


main: TyrannosaurusNet.go
	go build TyrannosaurusNet.go

clean:
	rm -rf $(BUILD_DIR)
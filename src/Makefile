ifndef PLUGINS_DIR
	export BUILD_DIR=$(CURDIR)/../build
	export PLUGINS_DIR=$(BUILD_DIR)/plugins
endif

all: main servicesTarget 
	mv TyrannosaurusNet $(BUILD_DIR)/

servicesTarget: services/Makefile
	cd services; make


main: TyrannosaurusNet.go
	go build TyrannosaurusNet.go

clean:
	rm -rf $(BUILD_DIR)
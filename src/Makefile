ifndef PLUGINS_DIR
	export BUILD_DIR=$(CURDIR)/../build
	export PLUGINS_DIR=$(BUILD_DIR)/plugins
endif

all: servicesTarget mainTarget packageBuilderTarget
	mv $(MAIN_EXE_NAME) $(BUILD_DIR)/
	mv $(UPDATE_BUILDER_NAME) $(BUILD_DIR)/

servicesTarget: services/Makefile
	@cd services; make; cd ..


mainTarget: main/TyrannosaurusNet.go
	go build -o $(MAIN_EXE_NAME) main

packageBuilderTarget: packageBuilder/packageBuilder.go
	go build -o $(UPDATE_BUILDER_NAME) packageBuilder

clean:
	@rm -rf $(BUILD_DIR)
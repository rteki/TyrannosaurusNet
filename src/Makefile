ifndef PLUGINS_DIR
	export BUILD_DIR=$(CURDIR)/../build
	export PLUGINS_DIR=$(BUILD_DIR)/plugins
endif

all: servicesTarget mainTarget
	mv TyrannosaurusNet $(BUILD_DIR)/

servicesTarget: services/Makefile
	@cd services; make; cd ..


mainTarget: main/TyrannosaurusNet.go
	go build -o TyrannosaurusNet tyrranosaurusNet

clean:
	@rm -rf $(BUILD_DIR)
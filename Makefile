export BUILD_DIR=$(CURDIR)/build
export PLUGINS_DIR=$(BUILD_DIR)/plugins
export GOPATH=$(CURDIR)

all: srcTarget resourcesTarget

srcTarget: src
	cd src; make; cd ..	

resourcesTarget: build
	cp -r resources ./build/resources

start: all
	cd build; ./TyrannosaurusNet

clean:
	rm -rf $(BUILD_DIR)
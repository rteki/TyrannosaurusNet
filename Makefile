export BUILD_DIR=$(CURDIR)/build
export PLUGINS_DIR=$(BUILD_DIR)/plugins
export GOPATH=$(CURDIR)

all:
	cd src; make

clean:
	rm -rf $(BUILD_DIR)
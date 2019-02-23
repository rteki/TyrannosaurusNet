export BUILD_DIR=$(CURDIR)/build
export PLUGINS_DIR=$(BUILD_DIR)/plugins
export GOPATH=$(CURDIR)

export MAIN_EXE_NAME=tNet
export UPDATE_BUILDER_NAME=tBuilder

all: srcTarget resourcesTarget

srcTarget: src
	cd src; make; cd ..	

resourcesTarget: build
	cp -r resources ./build/resources

start: all
	cd build; ./$(MAIN_EXE_NAME)

genKeys:
	openssl genrsa -out tServer.key 2048
	openssl ecparam -genkey -name secp384r1 -out tServer.key
	openssl req -new -x509 -sha256 -key tServer.key -out tServer.crt -days 7
	mv tServer.key build/
	mv tServer.crt build/

clean:
	rm -rf $(BUILD_DIR)
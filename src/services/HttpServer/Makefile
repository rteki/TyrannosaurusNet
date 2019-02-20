CURRENT_DIR_NAME = `pwd|sed "s/.\+\///g"`

all:
	go build -buildmode=plugin -o $(PLUGINS_DIR)/$(CURRENT_DIR_NAME).so services/HttpServer
	mv tServer.crt $(BUILD_DIR)/
	mv tServer.key $(BUILD_DIR)/

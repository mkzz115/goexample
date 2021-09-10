OUTPUT=./output
CONFIG=./config

APP_NAME=mk_server

BUILD = go build
GOFLAGS =


all:
	@echo $(APP_NAME)
	$(BUILD) $(GOFLAGS) -o ${OUTPUT}/${APP_NAME} app/main.go
	chmod u+x ${OUTPUT}/${APP_NAME}
	cp ${CONFIG}/conf.toml ${OUTPUT}/

clean:
	rm -rf ${OUTPUT}/*

.PHONY: clean fmt
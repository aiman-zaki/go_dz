GO := go
MODULE_PATH := github.com/aiman-zaki/go_dz_http 
CURRENT_DIR := ${shell pwd}
swagger := podman run --rm -it -e GOPATH=$(HOME)/go:/go -v $(HOME):$(HOME) -w $(CURRENT_DIR) quay.io/goswagger/swagger 
pkgs  = $(shell GOFLAGS=-mod=vendor $(GO) list ./... | grep -vE -e /vendor/ -e /pkg/swagger/)
#-------------
# build 
#------------
build:
	$(GO) build ${MODULE_PATH}
#-----------
# clean
#------------
clean:
	$(GO) clean ${MODULE_PATH}
#------------
# Swagger
#-----------
swagger.doc:
	podman run -i yousan/swagger-yaml-to-html < pkg/swagger/swagger.yml > doc/index.html

swagger.validate:
	${swagger} validate pkg/swagger/swagger.yml
#-----------
# go code generate
#------------
.PHONY: generate
## Generate go code
generate:
	@echo "==> generating go code"
	sh ./pkg/swagger/gen.sh

#----------
# dev
#----------
.PHONY: dev
dev:
	podman rm -f postgres
	podman run --network=host --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
	sleep 5
	podman exec -it postgres psql -U postgres -c "CREATE DATABASE dz"
	${swagger} generate spec -o ./swagger_ui/swagger.json
	go run main.go

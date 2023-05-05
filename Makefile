GOPATH = $(shell go env GOPATH)

run:
	@echo $GOPATH
	$(GOPATH)/bin/swag init --ot go --tags health --instanceName health
	$(GOPATH)/bin/swag init --ot go --tags greeting --instanceName greeting
	go run ./...

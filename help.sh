#!/bin/bash

# --ot means --output-type
$(go env GOPATH)/bin/swag init --ot go --tags health --instanceName health
$(go env GOPATH)/bin/swag init --ot go --tags greeting --instanceName greeting

# links to test
links=(
  "http://localhost:8080/swagger/health/index.html"
  "http://localhost:8080/swagger/greeting/index.html"
  "http://localhost:8080/api/v1/ping"
  "http://localhost:8080/greeting"
)

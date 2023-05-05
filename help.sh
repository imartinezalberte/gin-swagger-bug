#!/bin/bash

# --ot means --output-type
$(go env GOPATH)/bin/swag init --ot go --tags health --instanceName health
$(go env GOPATH)/bin/swag init --ot go --tags greeting --instanceName greeting

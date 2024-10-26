#!/bin/bash

# This shell is executed before docker build.
# github actions to add env_values for dockerhub password, server on dockerhub is opt Directory.
docker run -v "$(pwd)":/opt/recognitionsystem golangci/golangci-lint:v1.41 sh /opt/recognitionsystem/build/ci/static_check/run_golangci_lint.sh





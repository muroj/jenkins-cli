#!/bin/bash

set -euox pipefail

go fmt . ./pkg/jenkins ./pkg/instana/ ./cmd 

go vet . ./pkg/jenkins ./pkg/instana ./cmd

# Ignore errors for now
golint . ./pkg/jenkins ./pkg/instana ./cmd

go build -v

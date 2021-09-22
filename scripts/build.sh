#!/bin/bash

set -euox pipefail

go fmt . ./pkg/jenkins ./pkg/instana/ ./cmd 

go vet . ./pkg/jenkins ./pkg/instana ./cmd

golint -set_exit_status . ./pkg/jenkins ./pkg/instana ./cmd

go build -v

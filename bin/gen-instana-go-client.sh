#!/bin/bash

### This script assumes you have the `java` and `wget` commands on the path

export UNIT_NAME='tron' # for example: prod
export TENANT_NAME='ibmdataaiwai' # for example: awesomecompany

# Download the generator to your current working directory:
if [ ! -f openapi-generator-cli.jar ]; then
    curl https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/4.3.1/openapi-generator-cli-4.3.1.jar -o openapi-generator-cli.jar 
fi

# generate a client library that you can vendor into your repository
java -jar openapi-generator-cli.jar generate -i https://instana.github.io/openapi/openapi.yaml -g go \
    -o pkg/instana/openapi \
    --server-variables "tenant=${TENANT_NAME},unit=${UNIT_NAME}" \
    --skip-validate-spec

# (optional) format the Go code according to the Go code standard
gofmt -s -w pkg/instana/openapi

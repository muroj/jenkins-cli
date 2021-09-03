#!/bin/bash

JENKINS_URL="https://ghenkins.bigdatalab.ibm.com" \
JENKINS_USER="jmuro" \
JENKINS_API_TOKEN=$(cat ~/.creds/ghenkins-jmuro) \
INSTANA_API_KEY=$(cat ~/.creds/instana-api-token.txt) \
 ./ghestimator


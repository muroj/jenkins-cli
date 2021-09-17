#!/bin/bash

export JENKINS_URL="https://ghenkins.bigdatalab.ibm.com"
export JENKINS_USER="jmuro" 
export JENKINS_API_TOKEN=$(cat ~/.creds/ghenkins-jmuro | cut -f2 -d":") 
export INSTANA_API_KEY=$(cat ~/.creds/instana-api-token.txt) 
export JOB_URL="job/watson-engagement-advisor/job/clu-algorithms-service/job/master" 

./tronci.bin jenkins --user="$JENKINS_USER" --api-token="$JENKINS_API_TOKEN" --url="$JENKINS_URL" get build "$JOB_URL"

./tronci.bin jenkins --user="$JENKINS_USER" --api-token="$JENKINS_API_TOKEN" --url="$JENKINS_URL" version

./tronci.bin jenkins --user="$JENKINS_USER" --api-token="$JENKINS_API_TOKEN" --url="$JENKINS_URL" plugin list 


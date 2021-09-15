#!/bin/bash

export JENKINS_URL="https://ghenkins.bigdatalab.ibm.com"
export JENKINS_USER="jmuro" 
export JENKINS_API_TOKEN=$(cat ~/.creds/ghenkins-jmuro | cut -f2 -d":") 
export INSTANA_API_KEY=$(cat ~/.creds/instana-api-token.txt) 
export JOB_URL="job/watson-engagement-advisor/job/clu-algorithms-service/job/PR-1084" 

./ghestimator --jenkinsUser "$JENKINS_USER" --jenkinsAPIToken "$JENKINS_API_TOKEN" --jobURL "$JOB_URL" --instanaAPIKey "$INSTANA_API_KEY" 

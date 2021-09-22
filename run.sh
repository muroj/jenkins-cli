#!/bin/bash

#export JENKINS_URL="http://localhost:8080"
#export JENKINS_USER="admin" 
#export JENKINS_API_TOKEN="$(cat ~/.creds/localhost-jenkins)"

#export JENKINS_URL="https://ghenkins.bigdatalab.ibm.com"
#export JENKINS_USER="jmuro"
#export JENKINS_API_TOKEN=$(cat ~/.creds/ghenkins-jmuro | cut -f2 -d":")

export JENKINS_URL="https://wcp-tron-team-ci-jenkins.swg-devops.com"
export JENKINS_USER="jmuro@ibm.com" 
export JENKINS_API_TOKEN="$(cat ~/.creds/ghenkins-taas)"

#tronci jenkins --user="$JENKINS_USER" --api-token="$JENKINS_API_TOKEN" --url="$JENKINS_URL" get build "$JOB_URL"

tronci jenkins --user="$JENKINS_USER" --api-token="$JENKINS_API_TOKEN" --url="$JENKINS_URL" version

tronci jenkins --user="$JENKINS_USER" --api-token="$JENKINS_API_TOKEN" --url="$JENKINS_URL" plugin list 

#tronci jenkins --user="$JENKINS_USER" --api-token="$JENKINS_API_TOKEN" --url="$JENKINS_URL" plugin install --plugin-list="$(cat ghenkins-plugins-reduced.json)"

#tronci jenkins --user="$JENKINS_USER" --api-token="$JENKINS_API_TOKEN" --url="$JENKINS_URL" plugin install --plugin-list='[{"name": "kubernetes", "version": "1.30.1"}]'

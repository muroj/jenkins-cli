package jenkins

import (
	"log"
)

func GetBuild(url string, user string, apiToken string, projectUrl string, buildId int64) {

	jenkinsCreds := JenkinsCredentials{
		Username: user,
		APIToken: apiToken,
	}
	jenkinsClient := NewJenkinsClient(url, jenkinsCreds, false)
	buildInfo, err := GetBuildInfo(projectUrl, buildId, jenkinsClient)

	if err != nil {
		log.Fatalf("Failed to retrieve required build information: %s", err)
	}

	buildInfo.PrintBuildInfo()
}

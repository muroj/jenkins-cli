/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.ibm.com/jmuro/tronci/pkg/jenkins"
)

var (
	url           string
	user          string
	apiToken      string
	jobURL        string
	buildId       int64
	enableDebug   bool
	jenkinsClient *jenkins.JenkinsAPIClient
)

var jenkinsCmd = &cobra.Command{
	Use:   "jenkins",
	Short: "Run a command against a jenkins instance.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		jenkinsCreds := jenkins.JenkinsCredentials{
			Username: user,
			APIToken: apiToken,
		}
		jenkinsClient = jenkins.NewJenkinsClient(url, jenkinsCreds, false)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the version of the target jenkins",
	Run: func(cmd *cobra.Command, args []string) {
		jenkins.GetVersion(jenkinsClient)
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get data for a jenkins object.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("You must specify the type of resource to get.")
		}
		return nil
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Display info for a specified build",
	Run: func(cmd *cobra.Command, args []string) {
		projectUrl := args[0]
		jenkins.GetBuild(jenkinsClient, projectUrl, buildId)
	},
}

func init() {

	jenkinsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	jenkinsCmd.PersistentFlags().StringVar(&url, "url", "", "URL of the Jenkins host (required), e.g. \"https://ghenkins.bigdatalab.ibm.com/\"")
	jenkinsCmd.PersistentFlags().StringVar(&user, "user", "", "Jenkins username (required)")
	jenkinsCmd.PersistentFlags().StringVar(&apiToken, "api-token", "", "Jenkins API token (required)")
	jenkinsCmd.MarkPersistentFlagRequired("url")
	jenkinsCmd.MarkPersistentFlagRequired("user")
	jenkinsCmd.MarkPersistentFlagRequired("api-token")

	buildCmd.Flags().Int64Var(&buildId, "id", 0, "ID of the target build (required), e.g. 22. An value of 0 indicates the most recent build")
	getCmd.AddCommand(buildCmd)

	pluginCmd.AddCommand(installPluginsCmd)
	jenkinsCmd.AddCommand(versionCmd)
	jenkinsCmd.AddCommand(pluginCmd)
	jenkinsCmd.AddCommand(getCmd)
}
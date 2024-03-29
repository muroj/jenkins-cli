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
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.ibm.com/jmuro/tronci/pkg/jenkins"
)

var (
	url            string
	user           string
	apiToken       string
	jobURL         string
	buildID        int64
	enableDebug    bool
	decryptSecrets bool
	jenkinsClient  *jenkins.APIClient
)

var jenkinsCmd = &cobra.Command{
	Use:   "jenkins",
	Short: "Run a command against a jenkins instance.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		injectViperFlags(cmd)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the version of the target jenkins",
	Run: func(cmd *cobra.Command, args []string) {
		jenkinsCreds := jenkins.Credentials{
			Username: user,
			APIToken: apiToken,
		}
		jenkinsClient = jenkins.NewJenkinsClient(url, jenkinsCreds, enableDebug)
		jenkins.GetVersion(jenkinsClient)
	},
}

var listCredentialsCmd = &cobra.Command{
	Use:   "list-system-credentials",
	Short: "Output the jenkins system credentials as XML",
	Run: func(cmd *cobra.Command, args []string) {
		jenkinsCreds := jenkins.Credentials{
			Username: user,
			APIToken: apiToken,
		}
		jenkinsClient = jenkins.NewJenkinsClient(url, jenkinsCreds, enableDebug)
		jenkins.ListSystemCredentials(jenkinsClient, decryptSecrets)
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get data for a jenkins object.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must specify the type of resource to get")
		}
		return nil
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Display info for a specified build",
	Run: func(cmd *cobra.Command, args []string) {
		projectURL := args[0]
		jenkinsCreds := jenkins.Credentials{
			Username: user,
			APIToken: apiToken,
		}
		jenkinsClient = jenkins.NewJenkinsClient(url, jenkinsCreds, enableDebug)
		jenkins.GetBuild(jenkinsClient, projectURL, buildID)
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Initiate a safe restart on the target jenkins instance",
	Run: func(cmd *cobra.Command, args []string) {
		jenkinsCreds := jenkins.Credentials{
			Username: user,
			APIToken: apiToken,
		}
		jc := jenkins.NewJenkinsClient(url, jenkinsCreds, enableDebug)
		jenkins.SafeRestart(jc)
	},
}

func init() {
	jenkinsCmd.PersistentFlags().StringVar(&url, "url", "", "URL of the Jenkins host (required), e.g. \"https://ghenkins.bigdatalab.ibm.com/\"")
	jenkinsCmd.PersistentFlags().StringVar(&user, "user", "", "Jenkins username (required)")
	jenkinsCmd.PersistentFlags().StringVar(&apiToken, "api-token", "", "Jenkins API token (required)")
	jenkinsCmd.PersistentFlags().BoolVarP(&enableDebug, "debug", "v", false, "Enable debug output")
	jenkinsCmd.MarkPersistentFlagRequired("url")
	jenkinsCmd.MarkPersistentFlagRequired("user")
	jenkinsCmd.MarkPersistentFlagRequired("api-token")

	buildCmd.Flags().Int64Var(&buildID, "id", 0, "ID of the target build (required), e.g. 22. A value of 0 indicates the most recent build")
	buildCmd.MarkFlagRequired("id")
	getCmd.AddCommand(buildCmd)

	listCredentialsCmd.PersistentFlags().BoolVarP(&decryptSecrets, "decrypt", "d", false, "Output credential secrets in plain text")

	jenkinsCmd.AddCommand(versionCmd)
	jenkinsCmd.AddCommand(listCredentialsCmd)
	jenkinsCmd.AddCommand(pluginCmd)
	jenkinsCmd.AddCommand(getCmd)
	jenkinsCmd.AddCommand(restartCmd)
}

func injectViperFlags(cmd *cobra.Command) {
	vcfg := viper.Sub("jenkins")

	if vcfg != nil {
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if !f.Changed && vcfg.IsSet(f.Name) {
				cmd.Flags().Set(f.Name, vcfg.GetString(f.Name))
			}
		})
	}
}

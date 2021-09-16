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
	"strconv"

	"github.com/spf13/cobra"
	"github.ibm.com/jmuro/ghestimator/pkg/jenkins"
)

var (
	url         string
	user        string
	apiToken    string
	jobURL      string
	buildNumber int64
	enableDebug bool
)

// jenkinsCmd represents the jenkins command
var jenkinsCmd = &cobra.Command{
	Use:   "jenkins",
	Short: "Run a command against a jenkins instance.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jenkins called")
	},
}

var versionCmd = &cobra.Command{
	Use:   "jenkins version",
	Short: "Output the version of the target jenkins",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jenkins version called")
	},
	// func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		return errors.New("You must specify the type of resource to get.")
	// 	}
	// 	if jenkins.isValidResource(args[0]) {
	// 		return nil
	// 	}
	// 	return fmt.Errorf("invalid color specified: %s", args[0])
	// 	},
	// }
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get data for a jenkins object.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		switch args[0] {
		case "build":
			{
				projectUrl := args[1]
				buildId, _ := strconv.ParseInt(args[2], 10, 64)
				jenkins.GetBuild(url, user, apiToken, projectUrl, buildId)
			}
		case "project":
			{
				panic("Not implemented")
			}
		case "version":
			{
				panic("Not implemented")
			}
		case "nodes":
			{
				panic("Not implemented")
			}
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("You must specify the type of resource to get.")
		}
		return nil
	},
	//ValidArgs: []string{"build", "project", "version", "nodes"},
}

var metricsCmd = &cobra.Command{
	Use:   "getmetrics",
	Short: "Output metrics for a specified jenkins build.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jenkins build metrics")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jenkinsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jenkinsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	jenkinsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	jenkinsCmd.PersistentFlags().StringVar(&url, "url", "", "URL of the Jenkins host (required), e.g. \"https://ghenkins.bigdatalab.ibm.com/\"")
	jenkinsCmd.PersistentFlags().StringVar(&user, "user", "", "Jenkins username (required)")
	jenkinsCmd.PersistentFlags().StringVar(&apiToken, "api-token", "", "Jenkins API token (required)")
	jenkinsCmd.MarkPersistentFlagRequired("url")
	jenkinsCmd.MarkPersistentFlagRequired("user")
	jenkinsCmd.MarkPersistentFlagRequired("api-token")

	getCmd.AddCommand(versionCmd)
	jenkinsCmd.AddCommand(getCmd)
	rootCmd.AddCommand(jenkinsCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
	"github.ibm.com/jmuro/tronci/pkg/jenkins"
)

var pluginListJSON string

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Manage Jenkins plugins",
}

var listPluginsCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed Jenkins plugins",
	Run: func(cmd *cobra.Command, args []string) {
		jenkinsCreds := jenkins.Credentials{
			Username: user,
			APIToken: apiToken,
		}
		jenkinsClient = jenkins.NewJenkinsClient(url, jenkinsCreds, enableDebug)
		jenkins.ListPlugins(jenkinsClient)
	},
}

var installPluginsCmd = &cobra.Command{
	Use:   "install",
	Short: "List installed Jenkins plugins",
	Run: func(cmd *cobra.Command, args []string) {
		jenkinsCreds := jenkins.Credentials{
			Username: user,
			APIToken: apiToken,
		}
		jenkinsClient = jenkins.NewJenkinsClient(url, jenkinsCreds, enableDebug)
		jenkins.InstallPlugins(jenkinsClient, pluginListJSON)
	},
}

func init() {
	usage := `List of plugins to install specified as JSON. For example, "[{"name": "docker-plugin", "version": "1.2.3" }, ...]"`
	installPluginsCmd.Flags().StringVarP(&pluginListJSON, "plugin-list", "j", "", usage)
	installPluginsCmd.MarkFlagRequired("plugin-list")
	pluginCmd.AddCommand(listPluginsCmd)
	rootCmd.AddCommand(jenkinsCmd)
}

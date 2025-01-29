package commands

import (
	"log"

	"github.com/maurotory/jenkins-cli/pkg/config"
	"github.com/maurotory/jenkins-cli/pkg/jenkins"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Shows information",
	Long: `Connects to the Jenkins host using your 
	credentials and shows information about the server.

	By default, credentials are obtained from  ~/.jctl/config.json`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, err := cmd.Flags().GetString(configFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		conf, err := config.GetConfig(configPath)
		if err != nil {
			log.Fatalf("%v", err)
		}
		jClient, err := jenkins.ConnectToJenkins(conf)
		if err != nil {
			log.Fatalf("%v", err)
		}
		jClient.Info()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

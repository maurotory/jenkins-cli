/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/maurotory/jenkins-cli/pkg/config"
	"github.com/maurotory/jenkins-cli/pkg/jenkins"
	"github.com/spf13/cobra"
)

// whoamiCmd represents the whoami command
var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Shows information about user",
	Long: `Connects to the Jenkins host using your 
	credentials and shows information about the user.

	Credentials must be saved in ~/.jctl/config.json`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.GetConfig()
		if err != nil {
			log.Fatalf("%v", err)
		}
		jClient, err := jenkins.ConnectToJenkins(conf)
		if err != nil {
			log.Fatalf("%v", err)
		}
		err = jClient.Whoami()
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(whoamiCmd)
}

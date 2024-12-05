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

var jobID string

// buildCmd represents the "list build" subcommand
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Lists builds",
	Long:  "Lists all the builds",
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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists of jobs and build",
	Long:  `Commands that allows listing the diferent jenkins resources`,
}

func init() {
	listCmd.AddCommand(buildCmd)
	listCmd.PersistentFlags().String("job", "", "job path")

	// Add optional flags to the "job" subcommand
	buildCmd.Flags().StringVar(&jobID, "id", "", "Optional ID for the job")

	rootCmd.AddCommand(listCmd)
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

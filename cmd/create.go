/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/maurotory/jenkins-cli/pkg/config"
	"github.com/maurotory/jenkins-cli/pkg/errors"
	"github.com/maurotory/jenkins-cli/pkg/jenkins"
	"github.com/maurotory/jenkins-cli/pkg/parameters"
	"github.com/spf13/cobra"
)

var paramsFlag string = "params"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create command to create various Jenkins resources",
	Long:  `Create command to create various Jenkins resources`,
}

// buildCmd represents the "list builds" subcommand
var createBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Run a build",
	Long:  "Run a build",
	Run: func(cmd *cobra.Command, args []string) {
		job, err := cmd.Flags().GetString(jobFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if job == "" {
			log.Fatalf("%s: %s", errors.EmptyFlag, jobFlag)
		}
		paramsFile, err := cmd.Flags().GetString(paramsFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}

		conf, err := config.GetConfig()
		if err != nil {
			log.Fatalf("%v", err)
		}
		j, err := jenkins.ConnectToJenkins(conf)
		if err != nil {
			log.Fatalf("%v", err)
		}
		params, err := parameters.GetParameters(paramsFile)
		if err != nil {
			log.Fatalf("%v", err)
		}
		_, err = j.CreateJob(job, params)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	createCmd.AddCommand(createBuildCmd)
	createBuildCmd.PersistentFlags().String(jobFlag, "", "Mandatory ID for the job")
	createBuildCmd.PersistentFlags().String(paramsFlag, "", "Path of the parameters file")

	rootCmd.AddCommand(createCmd)
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"fmt"
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
	Short: "Create a resource of the selected path",
	Long:  `Commands that allows creating different Jenkins resources`,
}

// buildCmd represents the "list builds" subcommand
var createBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Creates a build",
	Long:  "Creates a build and prints this new build ID",
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
		configPath, err := cmd.Flags().GetString(configFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		conf, err := config.GetConfig(configPath)
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
		buildId, err := j.CreateJob(job, params)
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Printf("buildId: %d\n", buildId)
	},
}

func init() {
	createCmd.AddCommand(createBuildCmd)
	createBuildCmd.PersistentFlags().String(jobFlag, "", jobFlagMsg)
	createBuildCmd.PersistentFlags().StringP(paramsFlag, "p", ".env", "Path where the parameters json file is stored")

	rootCmd.AddCommand(createCmd)
}

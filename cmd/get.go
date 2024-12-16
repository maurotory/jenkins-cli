/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/maurotory/jenkins-cli/pkg/config"
	"github.com/maurotory/jenkins-cli/pkg/errors"
	"github.com/maurotory/jenkins-cli/pkg/jenkins"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get command to create various Jenkins resources",
	Long:  `Get command to create various Jenkins resources`,
}

// buildCmd represents the "list builds" subcommand
var getBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Get a build",
	Long:  "Get a build",
	Run: func(cmd *cobra.Command, args []string) {
		job, err := cmd.Flags().GetString(jobFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if job == "" {
			log.Fatalf("%s: %s", errors.EmptyFlag, jobFlag)
		}
		build, err := cmd.Flags().GetInt64(buildFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if build == 0 {
			log.Fatalf("%s: %s", errors.EmptyFlag, buildFlag)
		}
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
		err = j.GetBuild(job, build)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	createCmd.AddCommand(createBuildCmd)
	createBuildCmd.PersistentFlags().String(jobFlag, "", "Mandatory ID for the job")
	createBuildCmd.PersistentFlags().String(buildFlag, "", "Mandatory ID for the build")

	rootCmd.AddCommand(createCmd)
}

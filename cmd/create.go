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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create command to create various Jenkins resources",
	Long:  `Create command to create various Jenkins resources`,
}

// buildCmd represents the "list builds" subcommand
var createBuildCmd = &cobra.Command{
	Use:   "builds",
	Short: "Lists builds",
	Long:  "Lists all the builds",
	Run: func(cmd *cobra.Command, args []string) {
		job, err := cmd.Flags().GetString(jobFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if job == "" {
			log.Fatalf("%s: %s", errors.EmptyFlag, jobFlag)
		}

		conf, err := config.GetConfig()
		if err != nil {
			log.Fatalf("%v", err)
		}
		j, err := jenkins.ConnectToJenkins(conf)
		if err != nil {
			log.Fatalf("%v", err)
		}
		_, err = j.ListBuilds(job)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

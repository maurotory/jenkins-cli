/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"log"

	"github.com/maurotory/jenkins-cli/pkg/config"
	"github.com/maurotory/jenkins-cli/pkg/errors"
	"github.com/maurotory/jenkins-cli/pkg/jenkins"
	"github.com/spf13/cobra"
)

var folderFlag string = "folder"

var quantityFlag string = "quantity"

var listJobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "List of jobs",
	Long:  "Lists all jobs of the specified folder, by default lists the jobs of the main view.",
	Run: func(cmd *cobra.Command, args []string) {
		folder, err := cmd.Flags().GetString(folderFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		quantity, err := cmd.Flags().GetInt(quantityFlag)
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
		_, err = j.ListJobs(folder, quantity)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

// buildCmd represents the "list builds" subcommand
var listBuildsCmd = &cobra.Command{
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
		quantity, err := cmd.Flags().GetInt(quantityFlag)
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
		_, err = j.ListBuilds(job, quantity)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists a resource of the selected type",
	Long:  `Commands that allows listing different Jenkins resources`,
}

func init() {
	listCmd.AddCommand(listBuildsCmd)
	listBuildsCmd.PersistentFlags().String(jobFlag, "", "Full project name of the job. e.g: my-main-folder/my-sub-folder/my-job")

	listJobsCmd.PersistentFlags().String(folderFlag, "", "Parent folder path where to list items to")

	listCmd.AddCommand(listJobsCmd)
	listCmd.PersistentFlags().Int(quantityFlag, 10, "Max quantity of jobs to list, default is 10")

	rootCmd.AddCommand(listCmd)
}

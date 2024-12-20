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
var viewFlag string = "view"

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
		view, err := cmd.Flags().GetString(viewFlag)
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
		_, err = j.ListJobs(folder, view, quantity)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

var listViewsCmd = &cobra.Command{
	Use:   "views",
	Short: "List of views",
	Long:  "Lists all views",
	Run: func(cmd *cobra.Command, args []string) {
		// folder, err := cmd.Flags().GetString(folderFlag)
		// if err != nil {
		// log.Fatalf("%v", err)
		// }
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
		err = j.ListViews(quantity)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

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

var listArtifactsCmd = &cobra.Command{
	Use:   "artifacts",
	Short: "Lists artifacts",
	Long:  "Lists all the artifacts",
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
		err = j.ListArtifacts(job, build)
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
	listBuildsCmd.PersistentFlags().String(jobFlag, "", jobFlagMsg)
	listJobsCmd.PersistentFlags().StringP(viewFlag, "v", "", "View where to get builds from")

	listCmd.AddCommand(listArtifactsCmd)
	listArtifactsCmd.PersistentFlags().String(jobFlag, "", jobFlagMsg)
	listArtifactsCmd.PersistentFlags().Int64(buildFlag, 0, buildFlagMsg)

	listJobsCmd.PersistentFlags().String(folderFlag, "", "Parent folder path where to list jobs to")

	listCmd.AddCommand(listJobsCmd)
	listCmd.PersistentFlags().Int(quantityFlag, 10, "Max quantity of jobs to list, default is 10")

	listCmd.AddCommand(listViewsCmd)

	rootCmd.AddCommand(listCmd)
}

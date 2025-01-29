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
		configPath, err := cmd.Flags().GetString(configFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		conf, err := config.GetConfig(configPath)
		if err != nil {
			log.Fatalf("%v", err)
		}
		job, err := cmd.Flags().GetString(jobFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if job == "" {
			job = conf.JobId
			if job == "" {
				log.Fatalf("%s: %s", errors.EmptyFlag, jobFlag)
			}
		}
		quantity, err := cmd.Flags().GetInt(quantityFlag)
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
		configPath, err := cmd.Flags().GetString(configFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		conf, err := config.GetConfig(configPath)
		if err != nil {
			log.Fatalf("%v", err)
		}
		job, err := cmd.Flags().GetString(jobFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if job == "" {
			job = conf.JobId
			if job == "" {
				log.Fatalf("%s: %s", errors.EmptyFlag, jobFlag)
			}
		}
		latest, err := cmd.Flags().GetBool(latestFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		build, err := cmd.Flags().GetInt64(buildFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if build == 0 && !latest {
			log.Fatalf("%s: %s", errors.EmptyFlag, buildFlag)
		}
		j, err := jenkins.ConnectToJenkins(conf)
		if err != nil {
			log.Fatalf("%v", err)
		}
		err = j.ListArtifacts(job, build, latest)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists a resource of the selected type",
	Long:  `Commands that allows listing different Jenkins resources`,
}

func init() {
	listCmd.AddCommand(listBuildsCmd)
	listBuildsCmd.PersistentFlags().StringP(jobFlag, "j", "", jobFlagMsg)
	listJobsCmd.PersistentFlags().StringP(viewFlag, "v", "", "View where to get builds from")

	listCmd.AddCommand(listArtifactsCmd)
	listArtifactsCmd.PersistentFlags().StringP(jobFlag, "j", "", jobFlagMsg)
	listArtifactsCmd.PersistentFlags().Int64P(buildFlag, "b", 0, buildFlagMsg)
	listArtifactsCmd.PersistentFlags().BoolP(latestFlag, "l", false, latestFlagMsg)

	listJobsCmd.PersistentFlags().String(folderFlag, "", "Parent folder path where to list jobs to")

	listCmd.AddCommand(listJobsCmd)
	listCmd.PersistentFlags().IntP(quantityFlag, "q", 5, "Max quantity of jobs to list, default is 10")

	listCmd.AddCommand(listViewsCmd)

	rootCmd.AddCommand(listCmd)
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"log"
	"os"

	"github.com/maurotory/jenkins-cli/pkg/config"
	"github.com/maurotory/jenkins-cli/pkg/errors"
	"github.com/maurotory/jenkins-cli/pkg/jenkins"
	"github.com/spf13/cobra"
)

var artifactFlag string = "artifact"
var outputFlag string = "output"

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a resource of the selected type",
	Long:  `Commands that allows listing different Jenkins resources`,
}

var getBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Gets a build",
	Long:  "Gets a build and prints its information",
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
		err = j.GetBuild(job, build)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

var getArtifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "Gets an artifact",
	Long:  "Gets an artifact and stores in the specified path",
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
		artifact, err := cmd.Flags().GetString(artifactFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if artifact == "" {
			log.Fatalf("%s: %s", errors.EmptyFlag, artifactFlag)
		}
		output, err := cmd.Flags().GetString(outputFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if output == "" {
			output, err = os.Getwd()
			if err != nil {
				log.Fatalf("%v", err)
			}
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
		err = j.GetArtifact(job, artifact, output, build)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	getCmd.AddCommand(getBuildCmd)
	getCmd.AddCommand(getArtifactCmd)
	getCmd.PersistentFlags().String(jobFlag, "", "Full project name of the job. e.g: my-main-folder/my-sub-folder/my-job")
	getCmd.PersistentFlags().Int64(buildFlag, 0, "ID number of the build")
	getArtifactCmd.PersistentFlags().StringP(outputFlag, "o", "", "Folder path to save the artifact")
	getArtifactCmd.PersistentFlags().String(artifactFlag, "", "Artifact name")

	rootCmd.AddCommand(getCmd)
}

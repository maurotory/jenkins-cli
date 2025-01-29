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

var artifactFlag string = "artifact"
var outputFlag string = "output"
var printArtifactFlag string = "print"

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
		err = j.GetBuild(job, build, latest)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

var getArtifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "Gets an artifact",
	Long:  "Gets an artifact and saves it at your current directory path",
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
		printArtifact, err := cmd.Flags().GetBool(printArtifactFlag)
		if err != nil {
			log.Fatalf("%v", err)
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
		j, err := jenkins.ConnectToJenkins(conf)
		if err != nil {
			log.Fatalf("%v", err)
		}
		err = j.GetArtifact(job, artifact, output, build, latest, printArtifact)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	getCmd.AddCommand(getBuildCmd)
	getCmd.AddCommand(getArtifactCmd)
	getCmd.PersistentFlags().StringP(jobFlag, "j", "", jobFlagMsg)
	getCmd.PersistentFlags().BoolP(latestFlag, "l", false, latestFlagMsg)
	getCmd.PersistentFlags().Int64P(buildFlag, "b", 0, buildFlagMsg)
	getArtifactCmd.PersistentFlags().StringP(outputFlag, "o", "", "Custom filename path where to save the artifact")
	getArtifactCmd.PersistentFlags().StringP(artifactFlag, "a", "", "Artifact Name which will be downloded")
	getArtifactCmd.PersistentFlags().BoolP(printArtifactFlag, "p", false, "Print artifacts contents, if artifact contains ASCI characthers")
	getArtifactCmd.PersistentFlags().BoolP(latestFlag, "l", false, latestFlagMsg)

	rootCmd.AddCommand(getCmd)
}

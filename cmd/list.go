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

var jobFlag string = "jobId"

// buildCmd represents the "list build" subcommand
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Lists builds",
	Long:  "Lists all the builds",
	Run: func(cmd *cobra.Command, args []string) {
         job, err := cmd.Flags().GetString(jobFlag)
         if err != nil {
			log.Fatalf("%v", err)
         }
         if job == "" { 
             log.Fatalf("%s: %s",errors.EmptyFlag, jobFlag)
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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists of jobs and build",
	Long:  `Commands that allows listing the diferent jenkins resources`,
}

func init() {
	listCmd.AddCommand(buildCmd)
	buildCmd.PersistentFlags().String(jobFlag, "", "Mandatory ID for the job")

	rootCmd.AddCommand(listCmd)
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

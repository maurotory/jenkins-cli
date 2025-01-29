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

var followFlag string = "follow"

// showCmd represents the show command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Shows logs of a build",
	Long:  `Shows logs of a build`,
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
		build, err := cmd.Flags().GetInt64(buildFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		latest, err := cmd.Flags().GetBool(latestFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if build == 0 && !latest {
			log.Fatalf("%s: %s", errors.EmptyFlag, buildFlag)
		}
		follow, err := cmd.Flags().GetBool(followFlag)
		if err != nil {
			log.Fatalf("%v", err)
		}
		j, err := jenkins.ConnectToJenkins(conf)
		if err != nil {
			log.Fatalf("%v", err)
		}
		_, err = j.Logs(job, build, follow, latest)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	logsCmd.PersistentFlags().StringP(jobFlag, "j", "", "Full project name of the job. e.g: my-main-folder/my-sub-folder/my-job")
	logsCmd.PersistentFlags().Int64P(buildFlag, "b", 0, "ID number of the build")
	logsCmd.PersistentFlags().BoolP(followFlag, "f", false, "If set, the logs will be prompted in follow mode")
	logsCmd.PersistentFlags().BoolP(latestFlag, "l", false, latestFlagMsg)

	rootCmd.AddCommand(logsCmd)

}

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

var buildFlag string = "buildId"
var followFlag string = "follow"

// showCmd represents the show command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Shows logs of a build",
	Long:  `Shows logs of a build`,
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
		follow, err := cmd.Flags().GetBool(followFlag)
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
		_, err = j.Logs(job, build, follow)
		if err != nil {
			log.Fatalf("%v", err)
		}
	},
}

func init() {
	logsCmd.PersistentFlags().String(jobFlag, "", "Mandatory ID for the job")
	logsCmd.PersistentFlags().Int64(buildFlag, 0, "Mandatory ID for the build")
	logsCmd.PersistentFlags().Bool(followFlag, false, "If set, the logs will be prompted in follow mode")

	rootCmd.AddCommand(logsCmd)

}

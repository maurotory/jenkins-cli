/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var configFlag string = "config"

var buildFlag string = "buildId"
var buildFlagMsg string = "ID number of the build"

var jobFlag string = "jobId"
var jobFlagMsg string = "Full project name of the job. e.g: my-main-folder/my-sub-folder/my-job"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jctl",
	Short: "Jenkins CLI",
	Long:  `Easy to use CLI to interact with Jenkins.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP(configFlag, "c", "", "Path where the CLI Jenkins configuration is stored")
}

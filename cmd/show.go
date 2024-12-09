/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Shows logs of a build",
	Long: `Shows logs of a build`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logs called")
	},
}

func init() {
	logsCmd.PersistentFlags().String(jobFlag, "", "Mandatory ID for the job")

	rootCmd.AddCommand(logsCmd)

}

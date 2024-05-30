package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"mydata/internal/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tool",
	Long:  "Print the version number of tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.GetRawInfo())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

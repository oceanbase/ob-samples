package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mydata",
	Short: "A Fast and Flexible export data from remote mysql server",
	Long:  `A Fast and Flexible export data from remote mysql server`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
		}
	},
	SilenceUsage: true,
}

var debug bool

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "print stack log")

	rootCmd.Flags().SortFlags = false // 禁止flag排序
}

func Execute() {
	rootCmd.SilenceUsage = true

	if err := rootCmd.Execute(); err != nil {
		if debug {
			fmt.Printf("%+v\n", err)
		}
		os.Exit(1)
	}
}

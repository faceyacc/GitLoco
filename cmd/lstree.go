/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/faceyacc/gitloco/internals"
	"github.com/spf13/cobra"
)

// lstreeCmd represents the lstree command
var lstreeCmd = &cobra.Command{
	Use:   "ls-tree",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dummy_tree_sha := "I'm with stupid"
		res := internals.LsTree(dummy_tree_sha)
		fmt.Printf(res)

	},
}

func init() {
	rootCmd.AddCommand(lstreeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lstreeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lstreeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

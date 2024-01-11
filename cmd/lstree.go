/*
Copyright © 2024 NAME HERE <justfacey@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/faceyacc/gitloco/internals"
	"github.com/spf13/cobra"
)

// Usage: 'gitloco ls-tree <tree_sha>'
var lstreeCmd = &cobra.Command{
	Use:   "ls-tree",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args[0]) <= 30 {
			fmt.Fprintf(os.Stderr, "Incorrect tree hash\n")
			os.Exit(1)
		}

		tree_sha := args[0]

		// Print tree entries to stdout
		internals.LsTree(tree_sha)

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

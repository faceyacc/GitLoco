/*
Copyright © 2024 Ty Facey justfacey@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Usage: "gitloco commit-tree <tree_sha> -p <commit_sha> -m <message>"
var commitTreeCmd = &cobra.Command{
	Use:   "commit-tree",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get tree_sha from args
		if len(args[0]) < 30 {
			fmt.Fprintf(os.Stderr, "Incorrect blob hash\n")
			os.Exit(1)
		}

		// check if commit_sha was given from flag args

		// get message from flag args

		fmt.Println("commitTree called")
	},
}

func init() {
	rootCmd.AddCommand(commitTreeCmd)

	hashobjectCmd.Flags().String("p", "", "parent hash")

	hashobjectCmd.Flags().String("m", "", "commit message")
}

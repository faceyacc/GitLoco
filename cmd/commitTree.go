/*
Copyright © 2024 Ty Facey justfacey@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/faceyacc/gitloco/internals"
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
		if len(args[0]) < 40 {
			fmt.Fprintf(os.Stderr, "Incorrect hash tree hash\n")
			os.Exit(1)
		}

		tree_sha := args[0]

		// Check if a message was given
		message, err := cmd.Flags().GetString("m")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing you commit message")
		}
		if len(message) <= 0 {
			fmt.Fprintf(os.Stderr, "You must add a message to your commit :)")
		}

		parent_hash, err := cmd.Flags().GetString("p")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error with your parent hash")
		}

		// Check if parent hash is of correct sha-1 length
		if len(parent_hash) != 0 && len(parent_hash) < 40 {
			fmt.Fprintf(os.Stderr, "Incorrect parent hash\n")
			os.Exit(1)
		}

		sha := internals.CommitTree(tree_sha, parent_hash, message)
		fmt.Printf("TREE COMMIT SHA: %v", sha)
	},
}

func init() {
	rootCmd.AddCommand(commitTreeCmd)

	commitTreeCmd.Flags().String("p", "", "parent hash")

	commitTreeCmd.Flags().String("m", "", "commit message")
}

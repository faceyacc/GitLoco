/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/faceyacc/gitloco/internals"
	"github.com/spf13/cobra"
)

// hashobjectCmd represents the hashobject command
var hashobjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		writeFile, _ := cmd.Flags().GetString("w")

		if writeFile != "" {
			hash, err := internals.HashObject(writeFile)
			if err != nil {
				fmt.Print(err)
			}
			fmt.Print(hash)
		} else {
			fmt.Fprintf(os.Stderr, "Must give a file\n")
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(hashobjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashobjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	hashobjectCmd.Flags().String("w", "", "Help message for toggle")
}

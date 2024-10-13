package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "quiz",
    Short: "A simple quiz CLI application",
    Long:  `This application runs a simple quiz in the command line.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to the quiz! Run `quiz start` to begin.")
    },
}

// Execute runs the root command.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
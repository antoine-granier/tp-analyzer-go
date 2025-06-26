package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "loganalyzer",
    Short: "Analyseur de logs",
}

func Execute() error {
    return rootCmd.Execute()
}
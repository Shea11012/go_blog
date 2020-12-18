package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}


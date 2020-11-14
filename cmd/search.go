package cmd

import (
	"github.com/MattRighetti/passwdvault/configuration"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [NAME]",
	Short: "Returns available password with pattern [NAME]",
	Long:  "examples here...",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := configuration.CheckInitFile(); err != nil {
			return err
		}

		if err := configuration.ParseConfigurationFile(); err != nil {
			return err
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Do stuff here...
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

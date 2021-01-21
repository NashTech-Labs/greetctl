package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <resource><resource-id>",
	Short: "Command to create resources",
	Long: `This command creates customized resources like cards & quotes. example:
	greetctl create card eva -n="Eva Green" -o=thanksgiving -l=fr
	greetctl create card bob --name="Bob Marley" --occasion=birthday 
	`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

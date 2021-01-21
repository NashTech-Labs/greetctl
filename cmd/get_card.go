package cmd

import (
	"greetctl/models/cards"

	"github.com/spf13/cobra"
)

// cardCmd represents the card command
var getCardCmd = &cobra.Command{
	Use:     "card <resource_id>",
	Short:   "This command is to read contents from card. ",
	Aliases: []string{"cards"},
	Long: ` For example:
	greetctl create card eva -n="Eva Green" -o=thanksgiving -l=fr
	greetctl get card eva`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cards.FetchCardByID(args[0])
		} else {
			cards.FetchAllCards()
		}

	},
}

func init() {
	getCmd.AddCommand(getCardCmd)
}

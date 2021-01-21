package cmd

import (
	"fmt"
	"greetctl/models/cards"
	"os"

	"github.com/spf13/cobra"
)

// createCardCmd represents the card command
var createCardCmd = &cobra.Command{
	Use:   "card <name>",
	Short: "Create cards.",
	Long: ` This command creates cards. Example:
	greetctl create card eva -n="Eva Green" -o=thanksgiving -l=fr
	greetctl create card bob --name="Bob Marley" --occasion=birthday`,

	Run: func(cmd *cobra.Command, args []string) {
		cardID := args[0]
		file, err := os.Create(".db/" + cardID)
		defer file.Close()
		occasion, err := cmd.Flags().GetString("occasion")
		if err != nil {
			fmt.Println(err)
		}

		lang, err := cmd.Flags().GetString("language")
		if err != nil {
			fmt.Println(err)
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println(err)
		}

		config := cards.Config{
			Occasion: occasion,
			Language: lang,
			UserName: name,
			CardID:   cardID,
		}

		cards.CreateAndSaveCard(config)
	},
}

func init() {
	createCardCmd.PersistentFlags().StringP("occasion", "o", "", "Possible values: newyear, thanksgiving, birthday")
	createCardCmd.PersistentFlags().StringP("language", "l", "en", "Possible values: en, fr")
	createCardCmd.PersistentFlags().StringP("name", "n", "", "Name of the user to whom you want to greet")
	createCardCmd.MarkPersistentFlagRequired("name")
	createCardCmd.MarkPersistentFlagRequired("occasion")
	createCmd.AddCommand(createCardCmd)

}

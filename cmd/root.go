package cmd

import (
	"fmt"
	"os"

	"../internal"
	"github.com/spf13/cobra"
)

var startCaps bool

var rootCmd = &cobra.Command{
	Use:   "mockbob [word or sentence (quoted)]",
	Short: "Generate alternating-case text for Spongebob memes.",
	Long: `mockbob will take any set of input text, and return it in a Spongebob meme mocking format.
	Examples:
	mockbob "do you even lift bro" -> dO yOu EvEn LiFt BrO
	mockbob -c "do you even lift bro" -> Do YoU eVeN lIfT bRo
	mockbob herpderp -> hErPdErP
	mockbob -c herpderp -> HeRpDeRp`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		result := ""
		// Validate if stdin passed
		if internal.IsStdin() {
			stdin, err := internal.ReadStdin()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// Convert stdin text
			result = internal.Bobify(stdin, startCaps)
		} else {
			// If no stdin, ensure an arg is passed to consume
			if len(args) < 1 {
				// If cli invoked without data, show help panel
				cmd.Help()
			} else {
				result = internal.Bobify(args[0], startCaps)
			}
		}
		fmt.Println(result)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&startCaps, "start-caps", "c", false, "start the text with a capital letter")
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tlkamp/mockbob/internal"
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
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		result := internal.Bobify(args[0], startCaps)
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

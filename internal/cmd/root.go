package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tlkamp/mockbob/internal/bobs"
)

var (
	versionInfo string = "No version specified"
	gitRev      string = "No revision specified"
	buildDate   string = "No date specified"

	versionMessage string = fmt.Sprintf(`Version: %v
Revision: %v
Build Date: %v`, versionInfo, gitRev, buildDate)

	startCaps  bool
	randomCaps bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display information about the version of mockbob.",
	Long:  `Displays version information about mockbob, including information from the SCM, git.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(versionMessage)
	},
	Hidden:     true,
	Deprecated: "version is deprecated and will be removed.",
}

type Bob interface {
	// Bobify accepts a string as input and returns the bobified version.
	Bobify(string) string
}

var rootCmd = &cobra.Command{
	Use:   "mockbob [word or sentence (quoted)]",
	Short: "Generate alternating-case text for Spongebob memes.",
	Long: `mockbob will take any set of input text, and return it in a Spongebob meme mocking format.

Examples:
  mockbob "do you even lift bro" -> dO yOu EvEn LiFt BrO
  mockbob -c "do you even lift bro" -> Do YoU eVeN lIfT bRo
  mockbob herpderp -> hErPdErP
  mockbob -c herpderp -> HeRpDeRp
  mockbob -r herpaderp ->HerPAdErP`,
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var b Bob

		b = bobs.NewStandardBobifier(startCaps)

		if randomCaps {
			b = bobs.NewRandomBobifier()
		}

		result := ""
		// Validate if stdin passed
		if isStdin() {
			stdin, err := readStdin()
			if err != nil {
				return err
			}
			// Convert stdin text
			result = b.Bobify(stdin)
		} else {
			// If no stdin, ensure an arg is passed to consume
			if len(args) < 1 {
				// If cli invoked without data, show help panel
				cmd.Help()
			} else {
				result = b.Bobify(args[0])
			}
		}
		cmd.Println(result)
		return nil
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
	rootCmd.Flags().BoolVarP(&randomCaps, "random-caps", "r", false, "randomize the capital letters through the text")
	rootCmd.AddCommand(versionCmd)
}

func isStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func readStdin() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	text := ""

	if scanner.Err() != nil {
		return text, scanner.Err()
	}

	for scanner.Scan() {
		text += scanner.Text()
	}

	return text, nil
}

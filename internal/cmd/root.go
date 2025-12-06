package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	leetBob "github.com/tlkamp/mockbob/internal/adapters/bobs/leet"
	"github.com/tlkamp/mockbob/internal/adapters/bobs/random"
	"github.com/tlkamp/mockbob/internal/adapters/bobs/standard"
	"github.com/tlkamp/mockbob/internal/channeler"
	"github.com/tlkamp/mockbob/internal/core/domain"
	"github.com/tlkamp/mockbob/internal/core/ports"
)

var (
	startCaps  bool
	randomCaps bool
	leet       bool
)

const (
	startCapsFlag  = "start-caps"
	randomCapsFlag = "random-caps"
	leetFlag       = "leet"
)

var rootCmd = &cobra.Command{
	Use:   "mockbob [word or sentence]",
	Short: "Generate alternating-case text for Spongebob memes.",
	Long: `mockbob will take any set of input text, and return it in a Spongebob meme mocking format.

Examples:
  mockbob herpderp     -> hErPdErP

  mockbob herp a derp  -> hErP a DeRp
  
  mockbob -c herpderp  -> HeRpDeRp
  mockbob -r herpaderp -> HerPAdErP
  mockbob -l herpaderp -> h3rp4d3rp

  echo "herpaderp" | mockbob   -> hErPdErP`,
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bob ports.Bob

		switch {
		case randomCaps:
			bob = random.NewRandomBobifier()
		case leet:
			bob = leetBob.NewLeetBobifier()
		default:
			bob = standard.NewStandardBobifier(startCaps)
		}

		app := domain.NewApp(bob)

		// Validate if stdin passed
		if isStdin() {
			c := channeler.Channeler(cmd.InOrStdin())
			for s := range c {
				cmd.Print(app.Process((s)))
			}
			return nil
		}

		// If no stdin, ensure an arg is passed to consume
		if len(args) == 0 {
			return errors.New("input is required")
		}

		input := strings.Join(args, " ")

		cmd.Println(app.Process(input))
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
	rootCmd.Flags().BoolVarP(&startCaps, startCapsFlag, "c", false, "start the text with a capital letter")
	rootCmd.Flags().BoolVarP(&randomCaps, randomCapsFlag, "r", false, "randomize the capital letters through the text")
	rootCmd.Flags().BoolVarP(&leet, leetFlag, "l", false, "convert text to 1337 5p34k")
	rootCmd.MarkFlagsMutuallyExclusive(startCapsFlag, randomCapsFlag, leetFlag)
}

func isStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

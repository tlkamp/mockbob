package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tlkamp/mockbob/internal/bobs"
)

var (
	startCaps  bool
	randomCaps bool
	leet       bool
)

type Bob interface {
	// Bobify accepts a string as input and returns the bobified version.
	Bobify(string) string
}

var rootCmd = &cobra.Command{
	Use:   "mockbob [word or sentence]",
	Short: "Generate alternating-case text for Spongebob memes.",
	Long: `mockbob will take any set of input text, and return it in a Spongebob meme mocking format.

Examples:
  mockbob "do you even lift bro" -> dO yOu EvEn LiFt BrO
  mockbob -c "do you even lift bro" -> Do YoU eVeN lIfT bRo
  mockbob herpderp -> hErPdErP
  mockbob -c herpderp -> HeRpDeRp
  mockbob -r herpaderp -> HerPAdErP
  mockbob -l herpaderp -> h3rp4d3rp`,
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var b Bob

		b = bobs.NewStandardBobifier(startCaps)

		if randomCaps {
			b = bobs.NewRandomBobifier()
		}

		if leet {
			b = bobs.NewLeetBobifier()
		}

		// Validate if stdin passed
		if isStdin() {
			c := channeler(cmd.InOrStdin())
			for s := range c {
				fmt.Print(b.Bobify((s)))
			}
			return nil
		}

		// If no stdin, ensure an arg is passed to consume
		if len(args) == 0 {
			return errors.New("input is required")
		}

		input := strings.Join(args, " ")

		cmd.Println(b.Bobify(input))
		return nil
	},
}

func channeler(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		buf := bufio.NewReader(r)

		for {
			l, err := buf.ReadString('\n')
			if l != "" {
				ch <- l
			}

			if err != nil {
				break
			}
		}

		close(ch)

	}()

	return ch
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
	rootCmd.Flags().BoolVarP(&leet, "leet", "l", false, "convert text to 1337 5p34k")
}

func isStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

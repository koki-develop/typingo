package cmd

import (
	"os"

	"github.com/koki-develop/typingo/internal/game"
	"github.com/spf13/cobra"
)

var (
	numWords int
)

var rootCmd = &cobra.Command{
	Use:   "typingo",
	Short: "Typing game written in Go",
	RunE: func(cmd *cobra.Command, args []string) error {
		m := game.New(&game.GameConfig{
			NumWords: numWords,
		})
		if err := game.Run(m); err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVar(&numWords, "num-words", 10, "the number of words")
}

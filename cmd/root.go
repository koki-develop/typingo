package cmd

import (
	"os"

	"github.com/koki-develop/typingo/internal/game"
	"github.com/spf13/cobra"
)

var (
	numTexts int
	beep     bool
)

var rootCmd = &cobra.Command{
	Use:   "typingo",
	Short: "Typing game written in Go",
	RunE: func(cmd *cobra.Command, args []string) error {
		m := game.New(&game.GameConfig{
			NumTexts: numTexts,
			Beep:     beep,
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
	rootCmd.Flags().IntVarP(&numTexts, "num-texts", "n", 10, "the number of texts")
	rootCmd.Flags().BoolVarP(&beep, "beep", "b", true, "whether to beep when mistaken")
}

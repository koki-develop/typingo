package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/koki-develop/typingo/internal/game"
	"github.com/spf13/cobra"
)

var (
	version string

	flagNumTexts int
	flagBeep     bool
)

var rootCmd = &cobra.Command{
	Use:   "typingo",
	Short: "Typing game written in Go",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if flagNumTexts < 1 {
			return fmt.Errorf(`invalid argument %d for "-n, --num-texts" flag: must be at least 1`, flagNumTexts)
		}

		m := game.New(&game.GameConfig{
			NumTexts: flagNumTexts,
			Beep:     flagBeep,
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
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}

	rootCmd.Version = version

	rootCmd.Flags().IntVarP(&flagNumTexts, "num-texts", "n", 10, "the number of texts")
	rootCmd.Flags().BoolVarP(&flagBeep, "beep", "b", true, "whether to beep when mistaken")
}

package utils

import (
	"fmt"

	"github.com/spf13/cobra"
)

func PrintDuration(durationCalculator func(bool) (string, error)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var includeSeconds bool
		var err error

		if includeSeconds, err = cmd.Flags().GetBool("sec"); err != nil {
			panic(err)
		}

		var durationOutput string
		if durationOutput, err = durationCalculator(includeSeconds); err != nil {
			panic(err)
		}

		fmt.Println(durationOutput)
	}
}

package cmd

import (
	"fmt"

	"github.com/eibhleag/trictrac/core"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "set the value at a key",
	Run: func(cmd *cobra.Command, args []string) {
		key, value := parseKeyValue(args)
		c := core.OpenCollection()
		result := c.Set(key, value)
		fmt.Printf("%s\n", formatResult(key, result))
		c.Close()
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}

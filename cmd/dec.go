package cmd

import (
	"fmt"

	"github.com/eibhleag/trictrac/core"
	"github.com/spf13/cobra"
)

// decCmd represents the dec command
var decCmd = &cobra.Command{
	Use:   "dec [key]",
	Short: "decrement the value at a key",
	Run: func(cmd *cobra.Command, args []string) {
		key, decrement := parseKeyValueOptional(1, args)
		c := core.OpenCollection()
		value := c.Dec(key, decrement)
		fmt.Printf("%s\n", formatResult(key, value))
		c.Close()
	},
}

func init() {
	RootCmd.AddCommand(decCmd)
}

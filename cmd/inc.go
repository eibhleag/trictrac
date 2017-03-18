package cmd

import (
	"fmt"

	"github.com/eibhleag/trictrac/core"
	"github.com/spf13/cobra"
)

// incCmd represents the inc command
var incCmd = &cobra.Command{
	Use:   "inc [key] [+value]",
	Short: "increment the value at a key",
	Run: func(cmd *cobra.Command, args []string) {
		key, increment := parseKeyValueOptional(1, args)
		c := core.OpenCollection()
		value := c.Inc(key, increment)
		fmt.Printf("%s\n", formatResult(key, value))
		c.Close()
	},
}

func init() {
	RootCmd.AddCommand(incCmd)
}

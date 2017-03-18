package cmd

import (
	"fmt"

	"github.com/eibhleag/trictrac/core"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "get the value at a key",
	Run: func(cmd *cobra.Command, args []string) {
		c := core.OpenCollection()
		key := args[0]
		value := c.Get(key)
		fmt.Printf("%s\n", formatResult(key, value))
		c.Close()
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}

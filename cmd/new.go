package cmd

import (
	"fmt"

	"github.com/eibhleag/trictrac/core"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [key] [value?]",
	Short: "create a new key",
	Run: func(cmd *cobra.Command, args []string) {
		key, decrement := parseKeyValueOptional(1, args)
		c := core.OpenCollection()
		result := c.New(key, decrement)
		fmt.Printf("%s\n", formatResult(key, result))
		c.Close()
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
}

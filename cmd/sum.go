package cmd

import (
	"fmt"
	"os"

	"github.com/eibhleag/trictrac/core"
	"github.com/spf13/cobra"
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum [prefix]",
	Short: "sum all keys with a given prefix",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("too many arguments")
			os.Exit(-1)
		} else if len(args) < 1 {
			fmt.Println("prefix required")
			os.Exit(-1)
		}
		c := core.OpenCollection()
		key := args[0]
		value := c.SumPrefix(key)
		fmt.Printf("%s\n", formatResult(key, value))
		c.Close()

	},
}

func init() {
	RootCmd.AddCommand(sumCmd)
}

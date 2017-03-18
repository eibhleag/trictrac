package cmd

import (
	"container/list"
	"fmt"
	"os"

	"github.com/eibhleag/trictrac/core"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all keys and their values",
	Run: func(cmd *cobra.Command, args []string) {
		c := core.OpenCollection()
		var list *list.List
		if len(args) > 1 {
			fmt.Println("too many arguments")
			os.Exit(-1)
		} else if len(args) == 1 {
			list = c.ListPrefix(args[0])
		} else {
			list = c.List()
		}
		for e := list.Front(); e != nil; e = e.Next() {
			pair := e.Value.(core.Pair)
			fmt.Printf("%s = %d\n", pair.Key, pair.Value)
		}
		c.Close()
	},
}

func init() {
	listCmd.Flags().BoolP("keys-only", "k", false, "Only show keys (no values)")
	RootCmd.AddCommand(listCmd)
}

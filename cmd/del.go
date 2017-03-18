package cmd

import (
	"github.com/eibhleag/trictrac/core"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete the provided key",
	Run: func(cmd *cobra.Command, args []string) {
		c := core.OpenCollection()
		key := args[0]
		c.Del(key)
		c.Close()
	},
}

func init() {
	RootCmd.AddCommand(delCmd)
}

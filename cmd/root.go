package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tt",
	Short: "low volume, local metrics",

	// placeholder for default action:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func parseKeyValue(args []string) (string, uint64) {
	if len(args) < 1 {
		fmt.Println("key required")
		os.Exit(-1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments")
		os.Exit(-1)
	}

	key := args[0]
	var value uint64
	var err error
	value, err = strconv.ParseUint(args[1], 10, 64)

	if err != nil {
		fmt.Printf("error: value must be an integer, got: %s\n", args[1])
		os.Exit(-1)
	}

	return key, value
}

func parseKeyValueOptional(defaultValue uint64, args []string) (string, uint64) {
	if len(args) > 1 {
		return parseKeyValue(args)
	}

	if len(args) < 1 {
		fmt.Println("key required")
		os.Exit(-1)
	}

	return args[0], defaultValue
}

var excludeKeys bool

func formatResult(key string, value uint64) string {
	if !excludeKeys {
		return fmt.Sprintf("%s = %d", key, value)
	} else {
		return fmt.Sprintf("%d", value)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&excludeKeys, "exclude-keys", "x", false, "exclude keys from output")
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

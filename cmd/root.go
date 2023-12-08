package cmd

import (
	"github.com/netdoop/netdoop/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCommand.PersistentFlags().StringVarP(&utils.ConfigFile, "config", "c", utils.DefaultConfigFile, "Config file")
	rootCommand.PersistentFlags().BoolVarP(&utils.VerboseMode, "verbose", "v", utils.VerboseMode, "Turn on verbose mode")
	rootCommand.PersistentFlags().BoolVarP(&utils.DebugMode, "debug", "", utils.DebugMode, "Turn on debug mode")
	rootCommand.Flags().SortFlags = false
	rootCommand.PersistentFlags().SortFlags = false

	rootCommand.CompletionOptions.DisableDefaultCmd = true
	rootCommand.AddCommand(serverRunCommand)
}

var rootCommand = &cobra.Command{
	Use:   "netdoop",
	Short: "NetDoop Server",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCommand.Execute()
}

package cmd

import (
	"os"

	"github.com/sanmesh-kakade/ftf/cmd/module"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ftf",
	Short: "FTF CLI is a command-line interface (CLI) tool that aids in the development of Facets.Cloud modules via innersourcing flow",
	Long:  `FTF CLI is a command-line interface (CLI) tool that aids in the development of Facets.Cloud modules via innersourcing flow. It packs a set of commands that help in generating, validating, testing and registering Facets.Cloud modules.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(module.ModuleCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommandPalettes()
}

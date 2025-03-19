package module

import (
	"github.com/spf13/cobra"
)

// moduleCmd represents the module command
var ModuleCmd = &cobra.Command{
	Use:   "module",
	Short: "Module is a palette that contains terraform module based commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

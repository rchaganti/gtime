package gtime

import (
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View existing timezones gtime configuration",
	Long:  "You can view existing timezone configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	configCmd.AddCommand(viewCmd)
}

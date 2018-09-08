package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	convertCmd.PersistentFlags().String("path", "p", "width of picture")
	convertCmd.PersistentFlags().String("width", "w", "width")
}

var path string
var width int

var convertCmd = &cobra.Command{
	Use:   "convert [OPTIONS]",
	Short: "Convert image to ASCII art",
	Run: func(cmd *cobra.Command, args []string) {
		lang := cmd.Flag("path").Value.String()
		fmt.Printf("%s world :)\n", lang)
	},
}

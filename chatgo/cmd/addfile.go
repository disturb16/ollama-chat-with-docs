package cmd

import (
	"chatgo/chromastore"

	"github.com/spf13/cobra"
)

var addFileCmd = &cobra.Command{
	Use:     "addf",
	Short:   "adds a file to the vectore store",
	Long:    ``,
	Args:    cobra.ExactArgs(1),
	Example: "chatgo addf /path/to/file",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		chromastore.AddDoc(cmd.Context(), &store, filePath)
	},
}

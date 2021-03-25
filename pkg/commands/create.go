package commands

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/texsch/pkg/commands/create"
	"github.com/Matt-Gleich/texsch/pkg/configuration"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"github.com/Matt-Gleich/texsch/pkg/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a document",
	Run: func(cmd *cobra.Command, args []string) {
		location.ChdirProjectRoot()
		classes := configuration.GetClassNames()

		if utils.GetBoolFlag(cmd, "classes") {
			for _, class := range classes {
				fmt.Println(class)
			}
			os.Exit(0)
		}
		if utils.GetBoolFlag(cmd, "doctypes") {
			for _, docType := range create.DocumentTypes {
				fmt.Println(docType)
			}
			os.Exit(0)
		}

		path := create.Document(cmd, classes)
		create.ClipboardAndOpen(path)
	},
}

func init() {
	createCmd.Flags().StringP("name", "n", "", "Name of the document")
	createCmd.Flags().StringP("type", "t", "", "Document type (e.g. Worksheet)")
	createCmd.Flags().StringP("class", "c", "", "Class name for the document (e.g. Physics)")
	createCmd.Flags().Bool("classes", false, "Output the user's classes")
	createCmd.Flags().Bool("doctypes", false, "Output the available document types")
	rootCmd.AddCommand(createCmd)
}

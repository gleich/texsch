package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gleich/statuser/v2"
	"github.com/gleich/texsch/pkg/commands/create"
	"github.com/gleich/texsch/pkg/config"
	"github.com/gleich/texsch/pkg/location"
	"github.com/gleich/texsch/pkg/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a document",
	Run: func(cmd *cobra.Command, args []string) {
		location.ChdirProjectRoot()
		classes := config.ClassNames(config.Read())

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
		if utils.GetBoolFlag(cmd, "templates") {
			files, err := ioutil.ReadDir("./texsch/templates")
			if err != nil {
				statuser.Error("Failed to get a list of all the templates", err, 1)
			}
			for _, file := range files {
				if !file.IsDir() {
					fmt.Println(strings.TrimSuffix(file.Name(), ".txt"))
				}
			}
			os.Exit(0)
		}

		path := create.Document(cmd, classes)
		create.Post(cmd, path)
	},
}

func init() {
	createCmd.Flags().StringP("name", "n", "", "Name of the document")
	createCmd.Flags().StringP("type", "t", "", "Document type (e.g. Worksheet)")
	createCmd.Flags().StringP("class", "c", "", "Class name for the document (e.g. Physics)")
	createCmd.Flags().String("template", "", "Template for the document")
	createCmd.Flags().Bool("classes", false, "Output the user's classes")
	createCmd.Flags().Bool("doctypes", false, "Output the available document types")
	createCmd.Flags().Bool("templates", false, "Output the available templates")
	createCmd.Flags().Bool("no-clipboard", false, "If the clipboard configuration should be ignored")
	createCmd.Flags().Bool("no-editor", false, "If the editor configuration should be ignored")
	rootCmd.AddCommand(createCmd)
}

package commands

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/configure"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure texsch with an interactive prompt",
	Long: `
Configure the following for texsch:

	1. General configuration (stored in texsch/configuration.yaml)
		1. What is your first and last name?
		2. What year are you in? (Freshman, Sophmore, Junior, or Senior)
	2. Class configuration (stored in texsch/classes.yaml)
		Ask the following for each class:
		1. What is the name of the class?
		2. What time is the class (i.e. period, block)?
		3. What is the name of your teacher (with title prefix; i.e. Mr. or Ms.)?
	3. Template configuration (templates stored in texsch/templates)
		1. Do you want to use the default templates?
		2. Do you want the default templates to have emojis?

`,
	Run: func(cmd *cobra.Command, args []string) {
		location.ChdirProjectRoot()
		configure.Write(
			configure.General(),
			configure.Classes(),
			configure.Templates(),
		)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

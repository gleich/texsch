package utils

import (
	"fmt"

	"github.com/gleich/statuser/v2"
	"github.com/spf13/cobra"
)

// Get a flag of type string from a cobra command
func GetStringFlag(cmd *cobra.Command, name string) string {
	flag, err := cmd.Flags().GetString(name)
	if err != nil {
		statuser.Error(fmt.Sprintf("Failed to get %v flag", name), err, 1)
	}
	return flag
}

// Get a flag of type boolean from a cobra command
func GetBoolFlag(cmd *cobra.Command, name string) bool {
	flag, err := cmd.Flags().GetBool(name)
	if err != nil {
		statuser.Error(fmt.Sprintf("Failed to get %v flag", name), err, 1)
	}
	return flag
}

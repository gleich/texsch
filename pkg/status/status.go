package status

import "github.com/fatih/color"

func Success(msg string) {
	c := color.New(color.FgGreen).Add(color.Bold)
	c.Println("✔", msg)
}

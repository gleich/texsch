package status

import (
	"fmt"

	"github.com/fatih/color"
)

func Step(emoji string, msg string) {
	fmt.Println("_______________________________________________________________________________________")
	fmt.Printf("︳%v  %v▕\n", emoji, msg)
	fmt.Println()
}

func Success(msg string) {
	c := color.New(color.FgGreen).Add(color.Bold)
	c.Println("✔", msg)
}

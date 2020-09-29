package status

import (
	"fmt"
)

func Step(emoji string, msg string) {
	fmt.Println("_______________________________________________________________________________________")
	fmt.Printf("︳%v  %v▕\n", emoji, msg)
	fmt.Println()
}

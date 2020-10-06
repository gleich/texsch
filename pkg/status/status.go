package status

import (
	"fmt"

	"github.com/enescakir/emoji"
)

func Step(e emoji.Emoji, msg string) {
	fmt.Println("_______________________________________________________________________________________")
	fmt.Print(emoji.Sprintf("︳%v  %v▕\n", e, msg))
	fmt.Println()
}

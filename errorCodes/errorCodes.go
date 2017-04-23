package errorCodes

import (
	"fmt"
	"os"
)

func GolemKill(error int, print string) {
	fmt.Printf(print)
	os.Exit(error)
}

func GolemError(print string) {
	fmt.Printf(print)
}
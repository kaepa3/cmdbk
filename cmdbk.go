package cmdbk

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var Flg = true

//	Start callback
func Start(action func(cmd string)) {
	go func() {
		for Flg {
			var input string
			fmt.Scanln(&input)
			action(input)
		}
	}()
}

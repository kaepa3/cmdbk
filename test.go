package main

import (
	"fmt"
	"time"

	"github.com/kaepa/cmdbk"
)

var result = ""

func main() {
	cmdbk.Start(func(text string) {
		result = texts
		cmdbk.Flg = false
	})
	fmt.Println("foo")
	///	fmt.Fprint(os.Stdout, "foo\n")
	for cmdbk.Flg == true {
		time.Sleep(100 * time.Millisecond)
	}

	if result != "foo" {
		fmt.Println("cant call:", result)
	}
}

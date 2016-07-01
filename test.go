package main

import (
	"fmt"
	"time"

	"./cmdbk"
)

var result = ""

func main() {
	cmdbk.Start(func(text string) {
		result = text
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

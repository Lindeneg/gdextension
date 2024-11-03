package utils

import (
	"fmt"
	"os"
)

type Crasher struct {
	Root string
}

func (c Crasher) DieCleanOnErr(err error) {
	if err == nil {
		return
	}
	fmt.Println(err)
	os.RemoveAll(c.Root)
	os.Exit(1)
}

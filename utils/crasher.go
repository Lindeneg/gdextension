package utils

import (
	"fmt"
	"os"
)

type Crasher struct {
	Root string
}

func NewCrasher(root string) Crasher {
	return Crasher{Root: root}
}

func (c Crasher) DieCleanOnErr(err error) {
	if err == nil {
		return
	}
	fmt.Println(err)
	os.RemoveAll(c.Root)
	os.Exit(1)
}

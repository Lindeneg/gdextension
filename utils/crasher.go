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

func (c Crasher) DieCleanOnErr(err error, removeRoot bool) {
	if err == nil {
		return
	}
	fmt.Println("Setup failed with this error:")
	fmt.Println(err)
	if removeRoot {
		os.RemoveAll(c.Root)
	}
	os.Exit(1)
}

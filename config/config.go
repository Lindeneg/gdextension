package config

import (
	"errors"
	"flag"
	"fmt"

	"github.com/lindeneg/godot-utils/godot"
	"github.com/lindeneg/godot-utils/project"
)

type C struct {
	P           *project.P
	G           *godot.G
	Jobs        int
	WithUtils   bool
	WithMacros  bool
	WithGlobals bool
	WithLogger  bool
	WithExample bool
	WithAll     bool
	DryRun      bool
}

func New() *C {
	c := C{}
	c.P = &project.P{}
	c.G = &godot.G{}
	flag.IntVar(&c.Jobs, "jobs", 14, "corresponds to scons '-j' flag")
	flag.BoolVar(&c.WithUtils, "utils", false, "include c++ utility files (default false)")
	flag.BoolVar(&c.WithMacros, "macros", false, "include c++ macro files (default false)")
	flag.BoolVar(&c.WithGlobals, "globals", false, "include c++ global files (default false)")
	flag.BoolVar(&c.WithLogger, "logger", false, "include c++ logger files (default false)")
	flag.BoolVar(&c.WithExample, "example", false, "include godot c++ example (default false)")
	flag.BoolVar(&c.WithAll, "all", false, "include all c++ files (default false)")
	flag.BoolVar(&c.DryRun, "dry", false, "do not write anything (default false)")

	flag.Usage = func() {
		fmt.Println("gdextension [OPTIONS] project_name godot_binary")
		fmt.Println("Arguments:\n  project_name\n\tname of project (required)\n  godot_binary\n\tpath to godot executable (required)")
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("Example:")
		fmt.Println("godot-utils -all -j 20 foobar /bin/godot")
	}
	return &c
}

func (c *C) Validate() error {
	var err error
	args := flag.Args()
	if len(args) != 2 {
		return errors.New("error: positional args missing")
	}
	if err := c.P.Set(args[0]); err != nil {
		return err
	}
	if err := c.G.Set(args[1]); err != nil {
		return err
	}
	if c.WithAll {
		c.WithGlobals = true
		c.WithMacros = true
		c.WithUtils = true
		c.WithLogger = true
		c.WithExample = true
		return err
	}
	if c.WithExample {
		c.WithMacros = true
		c.WithUtils = true
	} else if c.WithMacros && !c.WithUtils {
		c.WithUtils = true
	}
	return err
}

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CompileOpts struct {
	Cwd      string
	J        int
	Platform string
	Arch     string
}

func CompileScons(co CompileOpts) error {
	return compileScons(co)
}

func CompileSconsEx(co CompileOpts) error {
	return compileScons(co, "custom_api_file=../extension_api.json")
}

func SilentCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func compileScons(co CompileOpts, args ...string) error {
	if err := os.Chdir(co.Cwd); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", co.Cwd, err)
	}
	args = append(args,
		fmt.Sprintf("-j%d", co.J),
		"platform="+co.Platform,
		"arch="+co.Arch,
		"target=template_debug",
		"compiledb=yes",
	)
    fmt.Printf("\tcwd: '%s'\n\topts: '%s'\n", co.Cwd, strings.Join(args, " "))
	fmt.Println("Please wait.. This could take a few mins..")
	if err := SilentCmd("scons", args...); err != nil {
		return err
	}
	return nil
}

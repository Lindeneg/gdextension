package godot

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

const (
	boolVariable = "BoolVariable"
	enumVariable = "EnumVariable"
	targetPrefix = "from SCons.Variables import"
)

var (
	target      = []byte(targetPrefix + " *")
	patches     = []string{"targets", "windows", "linux", "ios"}
	replacement = map[string]string{
		"windows": boolVariable,
		"linux":   boolVariable,
		"ios":     boolVariable,
		"targets": fmt.Sprintf("%s, %s", boolVariable, enumVariable),
	}
)

// there's an import error from scons when compiling on windows
// with godot-cpp versions lower than 4.3. The error is fixed in this PR:
// https://github.com/godotengine/godot-cpp/pull/1504/files#diff-769d3a0fa0df88f8c7410e5350aeb9b2dab4fa58f9a4d4639b746e9d3483b706
// but if a lower verison of godot is detected and we're on windows, we apply the patches directly.
func patch(root string) error {
	for _, patch := range patches {
		fp := filepath.Join(root, "godot-cpp", "tools", patch+".py")
		r := fmt.Sprintf("%s %s", targetPrefix, replacement[patch])
		if err := patchFile(fp, r); err != nil {
			return err
		}
		fmt.Printf("\tpatched '%s'\n\t  --%s\n\t  ++%s\n", fp, string(target), r)
	}
	return nil

}

func patchFile(filePath string, replacement string) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(filePath), "tmp-patch-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	nb := bytes.Replace(b, target, []byte(replacement), 1)
	if _, err := tmp.Write(nb); err != nil {
		return err
	}
	tmp.Close()
	if err := os.Rename(tmp.Name(), filePath); err != nil {
		return err
	}
	return nil
}

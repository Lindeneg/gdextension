package godot

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/lindeneg/godot-utils/templates"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var (
	patchTargets = []string{"targets", "windows", "linux", "ios"}
)

// there's an import error from scons when compiling on windows
// with godot-cpp versions lower than 4.3. The error is fixed in this PR:
// https://github.com/godotengine/godot-cpp/pull/1504/files#diff-769d3a0fa0df88f8c7410e5350aeb9b2dab4fa58f9a4d4639b746e9d3483b706
// but if a lower verison of godot is detected and we're on windows, we apply the patches directly.
// TODO: figure out how go-diff supports `patch` generated diff files.
func patch(root string, tpls *templates.T) error {
	dmp := diffmatchpatch.New()
	for _, target := range patchTargets {
		fn := filepath.Join(root, "godot-cpp", "tools", target+".py")
		f, err := os.Create(fn)
		if err != nil {
			return err
		}
		defer f.Close()
		text, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		patch, err := tpls.Patch(target)
		if err != nil {
			return err
		}
		diff := dmp.DiffMain(string(text), patch, false)
		patches := dmp.PatchMake(diff)
		updatedText, _ := dmp.PatchApply(patches, string(text))
		if _, err = f.WriteAt([]byte(updatedText), 0); err != nil {
			return err
		}
		fmt.Printf("Succesfully patched '%s'\n", fn)
	}
	return nil

}

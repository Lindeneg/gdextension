package godot

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lindeneg/godot-utils/templates"
	"github.com/lindeneg/godot-utils/utils"
)

const (
	Remote          = "https://github.com/godotengine/godot-cpp"
	MinMajorVersion = 3
	MinMinorVersion = 1
)

type G struct {
	Bin     string
	Tag     string
	Version version
}

func (g *G) Set(s string) error {
	if s == "" {
		return errors.New("must not be empty")
	}
	g.Bin = s
	if err := g.setVersion(); err != nil {
		return err
	}
	return g.setTag()
}

func (g *G) String() string {
	return g.Bin
}

func (g *G) Clone(rootPath string) error {
	if err := os.Chdir(rootPath); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", rootPath, err)
	}
	if err := utils.SilentCmd("git", "init"); err != nil {
		return fmt.Errorf("failed to initialize git repo: %w", err)
	}
	if err := utils.SilentCmd("git", "submodule", "add", Remote); err != nil {
		return fmt.Errorf("failed to add submodule: %w", err)
	}
	godotCppPath := filepath.Join(rootPath, "godot-cpp")
	if err := os.Chdir(godotCppPath); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", godotCppPath, err)
	}
	if err := utils.SilentCmd("git", "checkout", "tags/"+g.Tag); err != nil {
		return fmt.Errorf("failed to checkout tag %s: %w", g.Tag, err)
	}
	if err := utils.SilentCmd("git", "submodule", "update", "--init"); err != nil {
		return fmt.Errorf("failed to update submodules: %w", err)
	}
	return nil
}

func (g *G) DumpExtension(rootPath string) error {
	if err := os.Chdir(rootPath); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", rootPath, err)
	}
	if err := utils.SilentCmd(g.Bin, "--dump-extension-api", "--headless"); err != nil {
		return err
	}
	return nil
}

func (g *G) Patch(root string, tpls *templates.T) error {
	return patch(root, tpls)
}

func (g *G) setVersion() error {
	cmd := exec.Command(g.Bin, "--version")
	output, err := cmd.Output()
	if err != nil {
		return errors.New("failed to find godot binary version")
	}
	v, err := newVersionFromString(string(output))
	if err != nil {
		return err
	}
	g.Version = v
	return nil
}

func (g *G) setTag() error {
	cmd := exec.Command("git", "ls-remote", "--tags", Remote)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to list remote tags: %w", err)
	}
	tag := fmt.Sprintf("godot-%s-stable", g.Version.String())
	tagRef := fmt.Sprintf("refs/tags/%s", tag)
	if !strings.Contains(string(output), tagRef) {
		return fmt.Errorf("tag '%s' does not exist on '%s'", tag, Remote)
	}
	g.Tag = tag
	return nil
}

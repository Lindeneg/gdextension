package project

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lindeneg/godot-utils/templates"
)

type P struct {
	NameL  string
	NameU  string
	Parent string
	Root   string
	Src    string
	Godot  string
}

func (p *P) Set(s string) error {
	if s == "" {
		return errors.New("must not be empty")
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	p.NameL = strings.ToLower(filepath.Base(s))
	p.NameU = strings.ToUpper(p.NameL)
	p.Parent = wd
	p.Root = filepath.Join(p.Parent, p.NameL)
	if directoryExists(p.Root) {
		return fmt.Errorf("directory '%s' already exists", p.NameL)
	}
	p.Src = filepath.Join(p.Root, "src")
	p.Godot = filepath.Join(p.Root, p.NameL)
	return nil
}

func (p *P) String() string {
	return p.NameL
}

func (p *P) CreateFolders(t *templates.T) error {
	folders := []string{
		p.Root,
		p.Src,
		p.Godot,
		filepath.Join(p.Root, "misc"),
		filepath.Join(p.Src, "core"),
		filepath.Join(p.Godot, "bin"),
	}
	if t.Cfg.WithLogger {
		folders = append(folders, filepath.Join(p.Src, "objects"))
	}
	if t.Cfg.WithExample {
		folders = append(folders, filepath.Join(p.Src, "nodes"))
	}
	for _, f := range folders {
		if err := os.Mkdir(f, 0755); err != nil {
			return err
		}
	}
	return nil
}

func (p *P) CreateFiles(t *templates.T) error {
	if err := p.createFile(filepath.Join(p.Root, ".gitignore"), t.GitIgnore); err != nil {
		return err
	}
	if err := p.createFile(filepath.Join(p.Root, "SConstruct"), t.Scons); err != nil {
		return err
	}
	if err := p.createFile(filepath.Join(p.Root, "misc", "build.sh"), t.Build); err != nil {
		return err
	}
	if err := p.createFile(filepath.Join(p.Root, "misc", "run.sh"), t.Run); err != nil {
		return err
	}
	if err := p.createFile(
		filepath.Join(p.Godot, "bin",
			fmt.Sprintf("%s-extension.gdextension", t.Cfg.ProjectName)),
		t.GDExtension); err != nil {
		return err
	}
	if err := p.createFile(
		filepath.Join(p.Godot, "project.godot"), t.ProjectGodot); err != nil {
		return err
	}
	if err := p.createFile(
		filepath.Join(p.Godot, "icon.svg"), t.Icon); err != nil {
		return err
	}
	if err := p.createFile(
		filepath.Join(p.Src, "core", "register_types.h"), t.RegisterH); err != nil {
		return err
	}
	if err := p.createFile(
		filepath.Join(p.Src, "core", "register_types.cpp"), t.RegisterCPP); err != nil {
		return err
	}
	if t.Cfg.WithMacros {
		if err := p.createFile(
			filepath.Join(p.Src, "core", "macros.h"), t.MacrosH); err != nil {
			return err
		}
	}
	if t.Cfg.WithUtils {
		if err := p.createFile(
			filepath.Join(p.Src, "core", "utils.h"), t.UtilsH); err != nil {
			return err
		}
		if err := p.createFile(
			filepath.Join(p.Src, "core", "utils.cpp"), t.UtilsCPP); err != nil {
			return err
		}
	}
	if t.Cfg.WithGlobal {
		if err := p.createFile(
			filepath.Join(p.Src, "core", "global.h"), t.GlobalH); err != nil {
			return err
		}
		if err := p.createFile(
			filepath.Join(p.Src, "core", "global.cpp"), t.GlobalCPP); err != nil {
			return err
		}

	}
	if t.Cfg.WithLogger {
		if err := p.createFile(
			filepath.Join(p.Src, "objects", "logger.h"), t.LoggerH); err != nil {
			return err
		}
		if err := p.createFile(
			filepath.Join(p.Src, "objects", "logger.cpp"), t.LoggerCPP); err != nil {
			return err
		}
	}
	if t.Cfg.WithExample {
		if err := p.createFile(
			filepath.Join(p.Src, "nodes", "player.h"), t.ExampleH); err != nil {
			return err
		}
		if err := p.createFile(
			filepath.Join(p.Src, "nodes", "player.cpp"), t.ExampleCPP); err != nil {
			return err
		}
	}
	return nil
}

func (p *P) createFile(name string, writer func(io.Writer) error) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	writer(f)
	if err := f.Sync(); err != nil {
		return err
	}
	return nil

}

func directoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

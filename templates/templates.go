package templates

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"strings"
	"text/template"
)

//go:embed *.gotpl
var FS embed.FS

type Template struct {
	tpl *template.Template
}

func (t Template) Execute(w io.Writer, data interface{}) error {
	tpl := t.tpl
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return err
	}
	io.Copy(w, &buf)
	return nil
}

type Config struct {
	Platform         string
	Arch             string
	Executable       string
	ProjectName      string
	ProjectNameUpper string
	GodotVersion     string
	MajorMinor       string
	Jobs             int
	WithGlobal       bool
	WithUtils        bool
	WithMacros       bool
	WithLogger       bool
	WithExample      bool
}

type T struct {
	Cfg          Config
	build        Template
	gdExtension  Template
	gitIgnore    Template
	globalCPP    Template
	globalH      Template
	patchIos     Template
	patchLinux   Template
	patchTargets Template
	patchWindows Template
	icon         Template
	loggerCPP    Template
	loggerH      Template
	macrosH      Template
	exampleCPP   Template
	exampleH     Template
	projectGodot Template
	registerCPP  Template
	registerH    Template
	run          Template
	scons        Template
	utilsCPP     Template
	utilsH       Template
}

func New(c Config) *T {
	t := T{Cfg: c}
	t.build = parseFS(FS, "build.gotpl")
	t.gdExtension = parseFS(FS, "extension.gdextension.gotpl")
	t.gitIgnore = parseFS(FS, "gitignore.gotpl")
	t.globalCPP = parseFS(FS, "global.cpp.gotpl")
	t.globalH = parseFS(FS, "global.h.gotpl")
	t.patchIos = parseFS(FS, "godot_tools_ios.gotpl")
	t.patchLinux = parseFS(FS, "godot_tools_linux.gotpl")
	t.patchTargets = parseFS(FS, "godot_tools_targets.gotpl")
	t.patchWindows = parseFS(FS, "godot_tools_windows.gotpl")
	t.icon = parseFS(FS, "icon.svg.gotpl")
	t.loggerCPP = parseFS(FS, "logger.cpp.gotpl")
	t.loggerH = parseFS(FS, "logger.h.gotpl")
	t.macrosH = parseFS(FS, "macros.h.gotpl")
	t.exampleCPP = parseFS(FS, "player.cpp.gotpl")
	t.exampleH = parseFS(FS, "player.h.gotpl")
	t.projectGodot = parseFS(FS, "project.godot.gotpl")
	t.registerCPP = parseFS(FS, "register_types.cpp.gotpl")
	t.registerH = parseFS(FS, "register_types.h.gotpl")
	t.run = parseFS(FS, "run.gotpl")
	t.scons = parseFS(FS, "scons.gotpl")
	t.utilsCPP = parseFS(FS, "utils.cpp.gotpl")
	t.utilsH = parseFS(FS, "utils.h.gotpl")
	return &t
}

func (t *T) Build(w io.Writer) error {
	return t.build.Execute(w, struct {
		Platform string
		Arch     string
		Jobs     int
	}{t.Cfg.Platform, t.Cfg.Arch, t.Cfg.Jobs})
}

func (t *T) GDExtension(w io.Writer) error {
	return t.gdExtension.Execute(w, struct {
		ProjectName  string
		GodotVersion string
	}{t.Cfg.ProjectName, t.Cfg.GodotVersion})
}

func (t *T) GitIgnore(w io.Writer) error {
	return t.gitIgnore.Execute(w, nil)
}

func (t *T) GlobalCPP(w io.Writer) error {
	return t.globalCPP.Execute(w, struct {
		ProjectName string
	}{t.Cfg.ProjectName})
}

func (t *T) GlobalH(w io.Writer) error {
	return t.globalH.Execute(w, struct {
		ProjectName      string
		ProjectNameUpper string
	}{t.Cfg.ProjectName, t.Cfg.ProjectNameUpper})
}

func (t *T) Patch(s string) (string, error) {
	sb := strings.Builder{}
	var err error
	switch s {
	case "targets":
		err = t.patchTargets.Execute(&sb, nil)
	case "ios":
		err = t.patchIos.Execute(&sb, nil)
	case "linux":
		err = t.patchLinux.Execute(&sb, nil)
	case "windows":
		err = t.patchWindows.Execute(&sb, nil)
	default:
		return "", fmt.Errorf("patch template '%s' not found", s)
	}
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}

func (t *T) Icon(w io.Writer) error {
	return t.icon.Execute(w, nil)
}

func (t *T) LoggerCPP(w io.Writer) error {
	return t.loggerCPP.Execute(w, struct {
		ProjectName string
	}{t.Cfg.ProjectName})
}

func (t *T) LoggerH(w io.Writer) error {
	return t.loggerH.Execute(w, struct {
		ProjectName      string
		ProjectNameUpper string
	}{t.Cfg.ProjectName, t.Cfg.ProjectNameUpper})
}

func (t *T) MacrosH(w io.Writer) error {
	return t.macrosH.Execute(w, struct {
		ProjectNameUpper string
	}{t.Cfg.ProjectNameUpper})
}

func (t *T) ExampleCPP(w io.Writer) error {
	return t.exampleCPP.Execute(w, struct {
		ProjectName string
		WithLogger  bool
	}{t.Cfg.ProjectName, t.Cfg.WithLogger})
}

func (t *T) ExampleH(w io.Writer) error {
	return t.exampleH.Execute(w, struct {
		ProjectName      string
		ProjectNameUpper string
		WithLogger       bool
	}{t.Cfg.ProjectName, t.Cfg.ProjectNameUpper, t.Cfg.WithLogger})
}

func (t *T) ProjectGodot(w io.Writer) error {
	return t.projectGodot.Execute(w, struct {
		ProjectName string
		MajorMinor  string
	}{t.Cfg.ProjectName, t.Cfg.MajorMinor})
}

func (t *T) RegisterCPP(w io.Writer) error {
	return t.registerCPP.Execute(w, struct {
		ProjectName string
		WithLogger  bool
		WithExample bool
	}{t.Cfg.ProjectName, t.Cfg.WithLogger, t.Cfg.WithExample})
}

func (t *T) RegisterH(w io.Writer) error {
	return t.registerH.Execute(w, struct {
		ProjectName      string
		ProjectNameUpper string
	}{t.Cfg.ProjectName, t.Cfg.ProjectNameUpper})
}

func (t *T) Run(w io.Writer) error {
	return t.run.Execute(w, struct {
		Executable  string
		ProjectName string
	}{t.Cfg.Executable, t.Cfg.ProjectName})
}

func (t *T) Scons(w io.Writer) error {
	return t.scons.Execute(w, struct {
		ProjectName      string
		ProjectNameUpper string
		Jobs             int
	}{t.Cfg.ProjectName, t.Cfg.ProjectNameUpper, t.Cfg.Jobs})
}

func (t *T) UtilsCPP(w io.Writer) error {
	return t.utilsCPP.Execute(w, struct {
		ProjectName string
	}{t.Cfg.ProjectName})
}

func (t *T) UtilsH(w io.Writer) error {
	return t.utilsH.Execute(w, struct {
		ProjectName      string
		ProjectNameUpper string
	}{t.Cfg.ProjectName, t.Cfg.ProjectNameUpper})

}

func parseFS(fs fs.FS, patterns ...string) Template {
	tpl := template.New(patterns[0])
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	return Template{tpl}
}

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/lindeneg/godot-utils/config"
	"github.com/lindeneg/godot-utils/godot"
	"github.com/lindeneg/godot-utils/templates"
	"github.com/lindeneg/godot-utils/utils"
)

func main() {
	cfg := config.New()

	flag.Parse()

	if err := cfg.Validate(); err != nil {
		fmt.Println(err)
		fmt.Println("run with '-h' for usage info")
		os.Exit(1)
	}

	fmt.Println("Setting up new Godot GDExtension C++ Project")
	fmt.Printf("Found project name '%s'\n", cfg.P)
	fmt.Printf("Detected platform '%s' on arch '%s'\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Detected Godot version '%s'\n", cfg.G.Version)
	fmt.Printf("Using tag '%s' from remote '%s'\n", cfg.G.Tag, godot.Remote)

	if cfg.DryRun {
		os.Exit(0)
	}

	tpls := templates.New(templates.Config{
		Platform:         runtime.GOOS,
		Arch:             runtime.GOARCH,
		Executable:       cfg.G.Bin,
		ProjectName:      cfg.P.NameL,
		ProjectNameUpper: cfg.P.NameU,
		GodotVersion:     cfg.G.Version.String(),
		MajorMinor:       cfg.G.Version.MajorMinor(),
		Jobs:             cfg.Jobs,
		WithGlobal:       cfg.WithGlobals,
		WithUtils:        cfg.WithUtils,
		WithMacros:       cfg.WithMacros,
		WithLogger:       cfg.WithLogger,
		WithExample:      cfg.WithExample,
	})

	steps := utils.MakeSteps().
		Add(
			utils.Step{
				Msg:      "Creating core folders",
				Callback: func() error { return cfg.P.CreateFolders(tpls) },
			},
			utils.Step{
				Msg:      "Creating core files",
				Callback: func() error { return cfg.P.CreateFiles(tpls) },
			},
			utils.Step{
				Msg:      "Cloning godot-cpp",
				Callback: func() error { return cfg.G.Clone(cfg.P.Root) },
			},
		)

	if tpls.Cfg.Platform == "windows" && cfg.G.Version.Major <= 4 && cfg.G.Version.Minor < 3 {
		steps.Add(
			utils.Step{
				Msg:      "Patching godot-cpp tools",
				Callback: func() error { return cfg.G.Patch(cfg.P.Root, tpls) },
			},
		)
	}

	steps.Add(
		utils.Step{
			Msg:      "Dumping GDExtension API JSON file",
			Callback: func() error { return cfg.G.DumpExtension(cfg.P.Root) },
		},
		utils.Step{
			Msg: "Compiling godot-cpp with custom extension API",
			Callback: func() error {
				return utils.CompileSconsEx(utils.CompileOpts{
					Cwd:      filepath.Join(cfg.P.Root, "godot-cpp"),
					J:        tpls.Cfg.Jobs,
					Platform: tpls.Cfg.Platform,
					Arch:     tpls.Cfg.Arch,
				})
			},
		},
		utils.Step{
			Msg: "Compiling main project",
			Callback: func() error {
				return utils.CompileScons(utils.CompileOpts{
					Cwd:      cfg.P.Root,
					J:        tpls.Cfg.Jobs,
					Platform: tpls.Cfg.Platform,
					Arch:     tpls.Cfg.Arch,
				})
			},
		},
	)

	c := utils.Crasher{Root: cfg.P.Root}
	c.DieCleanOnErr(steps.Execute())

	fmt.Println("GDExtension setup succesfully completed.")
}

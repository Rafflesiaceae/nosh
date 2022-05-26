package main

import (
	"fmt"
	"os"

	noshOs "github.com/Rafflesiaceae/nosh/os"

	"github.com/Rafflesiaceae/nosh/fs"
	"github.com/Rafflesiaceae/nosh/lang"

	"go.starlark.net/lib/json"
	"go.starlark.net/lib/math"
	"go.starlark.net/resolve"
	"go.starlark.net/starlark"
)

func usage(retCode int) {
	f := os.Stderr
	if retCode == 0 {
		f = os.Stdout
	}

	fmt.Fprintln(f, "usage: -c <cmd> | <script>")
	os.Exit(retCode)
}

func version() {
	fmt.Print("0.0.1")
	os.Exit(0)
}

func run(scriptPath string, src interface{}) {

	resolve.AllowGlobalReassign = true
	resolve.AllowRecursion = true
	resolve.AllowSet = true
	// resolve.LoadBindsGlobally = true

	// Builtins
	predeclared := starlark.StringDict{
		"args":     noshOs.Args,
		"assert":   starlark.NewBuiltin("assert", lang.Assert),
		"cd":       fs.ModuleChdir,
		"chdir":    fs.ModuleChdir,
		"cmp":      fs.ModuleCmp,
		"cmp_file": fs.ModuleCmpFile,
		"cmp_path": fs.ModuleCmpPath,
		"copy":     fs.ModuleCopy,
		"cp":       fs.ModuleCopy,
		"exists":   fs.ModuleExists,
		"exit":     noshOs.ModuleExit,
		"expand":   noshOs.ModuleExpand,
		"fail":     starlark.NewBuiltin("fail", lang.Fail),
		"find":     fs.ModuleFind,
		"fs":       fs.Module,
		"getenv":   noshOs.ModuleGetenv,
		"json":     json.Module,
		"math":     math.Module,
		"mkdir":    fs.ModuleMkdir,
		"move":     fs.ModuleMove,
		"mv":       fs.ModuleMove,
		"os":       noshOs.Module,
		"popd":     fs.ModulePopd,
		"pwd":      fs.ModulePwd,
		"quit":     noshOs.ModuleQuit,
		"read":     fs.ModuleRead,
		"readdir":  fs.ModuleReaddir,
		"remove":   fs.ModuleRemove,
		"run":      noshOs.ModuleRun,
		"setenv":   noshOs.ModuleSetenv,
		"touch":    fs.ModuleTouch,
		"write":    fs.ModuleWrite,
	}

	// Execute Starlark program in a file.
	thread := &starlark.Thread{Name: "nosh"}
	_, err := starlark.ExecFile(thread, scriptPath, src, predeclared)

	switch err := err.(type) {
	case *starlark.EvalError:
		fmt.Fprintf(os.Stderr, "%s\n", err.Backtrace())
		noshOs.PresetExit()
	case nil: // success
	default:
		fmt.Fprintf(os.Stderr, "Error in %v\n", err)
		noshOs.PresetExit()
	}
}

func main() {
	// "Cli"
	args := os.Args[1:]
	if len(args) < 1 {
		usage(1)
	}

	switch args[0] {
	case "-c":
		if len(args) == 1 {
			usage(1)
		}
		if len(args) > 2 {
			noshOs.SetArgs(args[2:])
		}
		run("-c", args[1])
	case "-h", "--help":
		usage(0)
	case "-v", "--version":
		version()
	default:
		if len(args) > 1 {
			noshOs.SetArgs(args[1:])
		}
		run(args[0], nil)
	}
}

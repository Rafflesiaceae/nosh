package main

import (
	"fmt"
	"os"

	noshos "github.com/Rafflesiaceae/nosh/os"

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
		"args":   noshos.Args,
		"assert": starlark.NewBuiltin("assert", lang.Assert),
		"exists": fs.ModuleExists,
		"exit":   noshos.ModuleExit,
		"expand": noshos.ModuleExpand,
		"find":   fs.ModuleFind,
		"fs":     fs.Module,
		"json":   json.Module,
		"math":   math.Module,
		"os":     noshos.Module,
		"getenv": noshos.ModuleGetenv,
		"setenv": noshos.ModuleSetenv,
		"quit":   noshos.ModuleQuit,
		"read":   fs.ModuleRead,
		"remove": fs.ModuleRemove,
		"run":    noshos.ModuleRun,
		"touch":  fs.ModuleTouch,
		"write":  fs.ModuleWrite,
	}

	// Execute Starlark program in a file.
	thread := &starlark.Thread{Name: "nosh"}
	_, err := starlark.ExecFile(thread, scriptPath, src, predeclared)

	switch err := err.(type) {
	case *starlark.EvalError:
		fmt.Fprintf(os.Stderr, "%s\n", err.Backtrace())
		os.Exit(1)
	case nil: // success
	default:
		fmt.Fprintf(os.Stderr, "Error in %v\n", err)
		os.Exit(1)
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
			noshos.SetArgs(args[2:])
		}
		run("-c", args[1])
	case "-h", "--help":
		usage(0)
	case "-v", "--version":
		version()
	default:
		if len(args) > 1 {
			noshos.SetArgs(args[1:])
		}
		run(args[0], nil)
	}
}

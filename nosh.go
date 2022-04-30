package main

import (
	"fmt"
	"nosh/fs"
	"nosh/lang"
	noshos "nosh/os"
	"os"

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

	fmt.Fprintln(f, "usage: <script>")
	os.Exit(retCode)
}

func version() {
	fmt.Print("0.0.1")
	os.Exit(0)
}

func run(scriptPath string) {

	resolve.AllowGlobalReassign = true
	resolve.AllowRecursion = true
	resolve.AllowSet = true
	// resolve.LoadBindsGlobally = true

	// Builtins
	predeclared := starlark.StringDict{
		"assert": starlark.NewBuiltin("assert", lang.Assert),
		"fs":     fs.Module,
		"json":   json.Module,
		"math":   math.Module,
		"os":     noshos.Module,
	}

	// Execute Starlark program in a file.
	thread := &starlark.Thread{Name: "nosh"}
	_, err := starlark.ExecFile(thread, scriptPath, nil, predeclared)
	if err != nil {
		panic(err)
	}

}

func main() {
	// "Cli"
	args := os.Args[1:]
	if len(args) != 1 {
		usage(1)
	}

	switch args[0] {
	case "-h", "--help":
		usage(0)
	case "-v", "--version":
		version()
	default:
		run(args[0])
	}
}

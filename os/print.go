package os

import (
	"fmt"
	"os"
	"strings"

	"github.com/Rafflesiaceae/nosh/internal"
	"go.starlark.net/starlark"
)

func Print(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		stderr bool = false
		sep         = " "
	)

	if err := internal.UnpackKwargs("print", kwargs, "sep?", &sep, "stderr?", &stderr); err != nil {
		return starlark.Bool(false), err
	}

	// build output string
	buf := new(strings.Builder)
	for i, v := range args {
		if i > 0 {
			buf.WriteString(sep)
		}
		if s, ok := starlark.AsString(v); ok {
			buf.WriteString(s)
		} else if b, ok := v.(starlark.Bytes); ok {
			buf.WriteString(string(b))
		} else {
			buf.WriteString(v.String())
		}
	}
	s := buf.String()

	if stderr {
		fmt.Fprintln(os.Stderr, s)
	} else {
		fmt.Fprintln(os.Stdout, s)
	}

	return starlark.None, nil
}

func Printf(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var err error

	var (
		fargs  []starlark.Value
		stderr bool = false
	)

	if fargs, err = internal.UnpackPositionalVarargs("print", args, starlark.None); err != nil {
		return starlark.Bool(false), err
	}

	if err = internal.UnpackKwargs("print", kwargs, "stderr?", &stderr); err != nil {
		return starlark.Bool(false), err
	}

	// split format-string
	if len(fargs) < 1 {
		return nil, fmt.Errorf("printf requires at least one argument (format-string)")
	}
	formatString, _ := starlark.AsString(fargs[0])
	fargs = fargs[1:]

	// []starlark.Value -> []interface{}
	var anyFargs []interface{}
	for _, v := range fargs {
		if s, ok := starlark.AsString(v); ok {
			anyFargs = append(anyFargs, s)
		} else {
			anyFargs = append(anyFargs, v)
		}
	}

	if stderr {
		_, err = fmt.Fprintf(os.Stderr, formatString, anyFargs...)
		if err != nil {
			return starlark.Bool(false), err
		}
	} else {
		_, err = fmt.Fprintf(os.Stdout, formatString, anyFargs...)
		if err != nil {
			return starlark.Bool(false), err
		}

	}

	return starlark.None, nil
}

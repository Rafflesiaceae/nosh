package internal

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"go.starlark.net/starlark"
)

// UnpackOneArg copied from ~/.go/pkg/mod/go.starlark.net@v0.0.0-20220328144851-d1966c6b9fcd/starlark/unpack.go:217:15
func UnpackOneArg(v starlark.Value, ptr interface{}) error {
	// On failure, don't clobber *ptr.
	switch ptr := ptr.(type) {
	case starlark.Unpacker:
		return ptr.Unpack(v)
	case *starlark.Value:
		*ptr = v
	case *string:
		s, ok := starlark.AsString(v)
		if !ok {
			return fmt.Errorf("got %s, want string", v.Type())
		}
		*ptr = s
	case *bool:
		b, ok := v.(starlark.Bool)
		if !ok {
			return fmt.Errorf("got %s, want bool", v.Type())
		}
		*ptr = bool(b)
	case *int, *int8, *int16, *int32, *int64,
		*uint, *uint8, *uint16, *uint32, *uint64, *uintptr:
		return starlark.AsInt(v, ptr)
	case *float64:
		f, ok := v.(starlark.Float)
		if !ok {
			return fmt.Errorf("got %s, want float", v.Type())
		}
		*ptr = float64(f)
	case **starlark.List:
		list, ok := v.(*starlark.List)
		if !ok {
			return fmt.Errorf("got %s, want list", v.Type())
		}
		*ptr = list
	case **starlark.Dict:
		dict, ok := v.(*starlark.Dict)
		if !ok {
			return fmt.Errorf("got %s, want dict", v.Type())
		}
		*ptr = dict
	case *starlark.Callable:
		f, ok := v.(starlark.Callable)
		if !ok {
			return fmt.Errorf("got %s, want callable", v.Type())
		}
		*ptr = f
	case *starlark.Iterable:
		it, ok := v.(starlark.Iterable)
		if !ok {
			return fmt.Errorf("got %s, want iterable", v.Type())
		}
		*ptr = it
	default:
		// v must have type *V, where V is some subtype of starlark.Value.
		ptrv := reflect.ValueOf(ptr)
		if ptrv.Kind() != reflect.Ptr {
			log.Panicf("internal error: not a pointer: %T", ptr)
		}
		paramVar := ptrv.Elem()
		if !reflect.TypeOf(v).AssignableTo(paramVar.Type()) {
			// The value is not assignable to the variable.

			// Detect a possible bug in the Go program that called Unpack:
			// If the variable *ptr is not a subtype of Value,
			// no value of v can possibly work.
			if !paramVar.Type().AssignableTo(reflect.TypeOf(new(starlark.Value)).Elem()) {
				log.Panicf("pointer element type does not implement Value: %T", ptr)
			}

			// Report Starlark dynamic type error.
			//
			// We prefer the Starlark Value.Type name over
			// its Go reflect.Type name, but calling the
			// Value.Type method on the variable is not safe
			// in general. If the variable is an interface,
			// the call will fail. Even if the variable has
			// a concrete type, it might not be safe to call
			// Type() on a zero instance. Thus we must use
			// recover.

			// Default to Go reflect.Type name
			paramType := paramVar.Type().String()

			// Attempt to call Value.Type method.
			func() {
				defer func() { recover() }()
				paramType = paramVar.MethodByName("Type").Call(nil)[0].String()
			}()
			return fmt.Errorf("got %s, want %s", v.Type(), paramType)
		}
		paramVar.Set(reflect.ValueOf(v))
	}
	return nil
}

// UnpackKwargs copied from go.starlark.net@v0.0.0-20220328144851-d1966c6b9fcd/starlark/unpack.go:93:6
func UnpackKwargs(fnname string, kwargs []starlark.Tuple, pairs ...interface{}) error {
	nparams := len(pairs) / 2
	var defined intset
	defined.init(nparams)

	paramName := func(x interface{}) (name string, skipNone bool) { // (no free variables)
		name = x.(string)
		if strings.HasSuffix(name, "??") {
			name = strings.TrimSuffix(name, "??")
			skipNone = true
		} else if name[len(name)-1] == '?' {
			name = name[:len(name)-1]
		}

		return name, skipNone
	}

	// keyword arguments
kwloop:
	for _, item := range kwargs {
		name, arg := item[0].(starlark.String), item[1]
		for i := 0; i < nparams; i++ {
			pName, skipNone := paramName(pairs[2*i])
			if pName == string(name) {
				// found it
				if defined.set(i) {
					return fmt.Errorf("%s: got multiple values for keyword argument %s",
						fnname, name)
				}

				if skipNone {
					if _, isNone := arg.(starlark.NoneType); isNone {
						continue kwloop
					}
				}

				ptr := pairs[2*i+1]
				if err := UnpackOneArg(arg, ptr); err != nil {
					return fmt.Errorf("%s: for parameter %s: %s", fnname, name, err)
				}
				continue kwloop
			}
		}
		err := fmt.Errorf("%s: unexpected keyword argument %s", fnname, name)
		names := make([]string, 0, nparams)
		for i := 0; i < nparams; i += 2 {
			param, _ := paramName(pairs[i])
			names = append(names, param)
		}
		// if n := spell.Nearest(string(name), names); n != "" {
		// 	err = fmt.Errorf("%s (did you mean %s?)", err.Error(), n)
		// }
		return err
	}

	return nil
}

// UnpackPositionalVarargs unpacks all positional varargs and returns them,
// pass an `argTypeBin` if you want to have runtime type-checking, pass
// `starlark.None` if you don't want to force all varargs to a specific type
func UnpackPositionalVarargs(fnname string, args starlark.Tuple, argTypeBin interface{}) (posargs []starlark.Value, err error) {
	// @XXX @TODO use generics once they arrive
	result := make([]starlark.Value, 0)

	// Unpack positional args
	for _, arg := range args {
		if argTypeBin != starlark.None {
			if err := UnpackOneArg(arg, &argTypeBin); err != nil {
				return nil, fmt.Errorf("unpacking: %s", arg)
			}
		}
		result = append(result, arg)
	}

	return result, nil
}

// UnpackPositionalVarargsString unpacks all positional varargs and returns them,
// pass an `argTypeBin` if you want to have runtime type-checking, pass
// `starlark.None` if you don't want to force all varargs to a specific type
func UnpackPositionalVarargsString(fnname string, args starlark.Tuple) (posargs []string, err error) {
	result := make([]string, 0)

	// Unpack positional args
	for _, arg := range args {
		{
			// On failure, don't clobber *ptr.
			s, ok := starlark.AsString(arg)
			if !ok {
				return nil, fmt.Errorf("got %s, want string", arg.Type())
			}
			result = append(result, s)
		}
	}

	return result, nil
}

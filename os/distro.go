package os

import (
	"fmt"
	"runtime"

	"go.starlark.net/starlark"
	strlk "go.starlark.net/starlark"
)

// {{{1 Distro
type Distro struct {
	Os   string
	Arch string
}

var (
	_ strlk.Value    = &Distro{}
	_ strlk.HasAttrs = &Distro{}
)

func (di *Distro) Freeze()               { /* @TODO */ }
func (di *Distro) Hash() (uint32, error) { /* @TODO */ return 0, nil }
func (di *Distro) Truth() strlk.Bool     { return true }
func (di *Distro) Type() string          { return "Distro" }
func (di *Distro) String() string        { return fmt.Sprintf("%s (%s)", di.Os, di.Arch) }

func (di *Distro) Attr(name string) (strlk.Value, error) {
	return map[string]strlk.Value{
		// @TODO Cmd
		"os":   strlk.String(di.Os),
		"arch": strlk.String(di.Arch),
	}[name], nil
}

func (fe *Distro) AttrNames() []string {
	return []string{"os", "arch"}
}

func distro(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

	if err := starlark.UnpackArgs("distro", args, kwargs); err != nil {
		return nil, err
	}

	return &Distro{
		Os:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}, nil
}

package fs

import (
	"fmt"
	"log"

	"go.starlark.net/starlark"
	strlk "go.starlark.net/starlark"

	"github.com/fsnotify/fsnotify"
)

// {{{1 WatchEvent
type WatchEvent struct {
	/**
	One of:
		CHMOD
		CREATE
		REMOVE
	*/
	Operation string
	FileEntry
}

var (
	_ strlk.Value    = &WatchEvent{}
	_ strlk.HasAttrs = &WatchEvent{}
)

func (we *WatchEvent) Freeze()               { /* @TODO */ }
func (we *WatchEvent) Hash() (uint32, error) { /* @TODO */ return 0, nil }
func (we *WatchEvent) Truth() strlk.Bool     { return true }
func (we *WatchEvent) Type() string          { return "WatchEvent" }

// func (we *WatchEvent) String() string        { return we.AbsPath }
func (we *WatchEvent) String() string { return we.Operation + "\t" + we.AbsPath }

func (we *WatchEvent) Attr(name string) (result strlk.Value, err error) {
	switch name {
	case "abspath":
		return strlk.String(we.AbsPath), nil
	case "basename":
		return strlk.String(we.BaseName), nil
	case "dir":
		return strlk.String(we.Dir), nil
	case "ext":
		return strlk.String(we.Ext), nil
	case "name":
		return strlk.String(we.Name), nil
	case "path":
		return strlk.String(we.Path), nil

	case "operation":
		return strlk.String(we.Operation), nil
	}
	return starlark.None, fmt.Errorf("Unknown attr: %s", name)
}

func (we *WatchEvent) AttrNames() []string {
	return []string{
		"abspath",
		"basename",
		"dir",
		"ext",
		"name",
		"path",

		"operation",
	}
}

func watch(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var (
		path     string
		callback starlark.Callable
	)

	if err := starlark.UnpackArgs("watch", args, kwargs, "path", &path, "callback", &callback); err != nil {
		return nil, err
	}

	// @TODO improve, current implementation is very basic
	// @TODO add non-blocking option
	// @TODO improve error handling (use error channel, don't use 'panic')
	{
		// Create new watcher.
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return nil, err
		}
		defer watcher.Close()

		// Start listening for events.
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}

					// populate starlark event
					strlkEvent := &WatchEvent{
						Operation: event.Op.String(),
					}
					err = strlkEvent.FileEntry.Init(event.Name)
					if err != nil {
						panic(err)
					}

					_, err = starlark.Call(thread, callback, starlark.Tuple{strlkEvent}, nil)
					if err != nil {
						panic(err)
					}

				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					log.Println("ERROR", err)
				}
			}
		}()

		err = watcher.Add(path)
		if err != nil {
			return nil, err
		}

		<-make(chan struct{}) // @XXX
	}

	return starlark.None, nil
}

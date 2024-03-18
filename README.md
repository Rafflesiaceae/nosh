# nosh

[![CI](https://github.com/Rafflesiaceae/nosh/actions/workflows/ci.yml/badge.svg)](https://github.com/Rafflesiaceae/nosh/actions/workflows/ci.yml)

Take `starlark`, take `go` and wrap them into a cross-plat shell of sorts.

Motivation is to do basic shell tasks in `bazel` with something cross-plat, easy
to provision via `bazel` and which tastes like `starlark`.

See [./quicktest.nosh](./quicktest.nosh) for examples.

## Todo
- [X] print / printf
- [X] assert
- [ ] regex
- [X] defer
- [ ] signal handling
- [ ] set -e (NOTE: there's no exceptions in starlark)
- [ ] set -x
- [X] json
- [X] math
- [ ] xml
- [ ] ini/cfg (?)
- [ ] glob/globby
- [ ] hash
- [ ] base64
- [ ] templating
- [X] getenv/setenv
- [X] expand env
    + [ ] also expand ~ tilde ?
- [X] sleep
- [ ] dates
- [ ] parallel
- [X] args
- [ ] cli / argparse ?
- [X] cli -c
- [ ] tee

- [X] os.run
	+ [X] capture
	+ [X] env
	+ [ ] pipes
	+ [X] redirections
	+ [ ] background
	+ [ ] timeouts
- [X] os.distro/kind/os-type/uname
- [ ] os.pkill <pid> <name-*>
- [ ] os.https://nim-lang.org/docs/dynlib.html
- [ ] os.getTempDir
- [ ] os.sendSignal
- [ ] os.registry (WIN-ONLY)
- [X] os.random
- [ ] os.clipboard

- [X] fs.chdir / cd
- [X] fs.popd
- [X] fs.pwd / pwd
- [X] fs.find
- [X] fs.join
- [X] fs.realpath
- [X] fs.basename
- [ ] fs.symlink-handling
- [X] fs.copy
- [X] fs.move
- [X] fs.exists
	+ [X] fs.is_file
	+ [X] fs.is_dir
- [X] fs.mkdir
- [ ] fs.samefile
- [ ] fs.stat
- [X] fs.read
- [X] fs.write
- [X] fs.remove
- [X] fs.touch
- [ ] fs.mktemp
- [X] fs.watch

- [ ] net.ping
- [ ] net.curl

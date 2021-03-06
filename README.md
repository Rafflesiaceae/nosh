# nosh

[![CI](https://github.com/Rafflesiaceae/nosh/actions/workflows/ci.yml/badge.svg)](https://github.com/Rafflesiaceae/nosh/actions/workflows/ci.yml)

A cross-platform "shell", maybe, eventually, but not really.

Maybe eventually useful to at least prototype cross-plat shell-tasks instead of
having to write `.bat` and `.sh` side-by-side and ontop trying to balance `bash`
version differences.

The idea sounds so simple, take `go` take `starlark` look towards `bash` look towards `python`,
put 'em together, what can go wrong?

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

- [ ] net.ping
- [ ] net.curl

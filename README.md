# nosh

[![CI](https://github.com/Rafflesiaceae/nosh/actions/workflows/main.yml/badge.svg)](https://github.com/Rafflesiaceae/nosh/actions/workflows/main.yml)

A cross-platform "shell", maybe, eventually, but not really.

Maybe eventually useful to at least prototype cross-plat shell-tasks instead of
having to write `.bat` and `.sh` side-by-side and ontop trying to balance `bash`
version differences.

The idea sounds so simple, take `go` take `starlark` take `bash` take `python`,
put 'em together, what can go wrong?

## Todo
- [X] assert
- [ ] regex
- [ ] defer (?)
- [ ] signal handling
- [ ] set -e (NOTE: there's no exceptions in starlark)
- [/] set -x
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
- [ ] PWD
- [X] sleep
- [ ] dates
- [ ] parallel
- [ ] args
- [ ] cli / argparse ?
- [X] cli -c

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
- [X] fs.find
- [ ] fs.join
- [ ] fs.realpath
- [ ] fs.basename
- [ ] fs.symlink-handling
- [X] fs.copy
- [X] fs.move
- [X] fs.exists
	+ [X] fs.is_file
	+ [X] fs.is_dir
- [X] fs.mkdir
- [ ] fs.stat
- [X] fs.read
- [X] fs.write
- [X] fs.remove
- [X] fs.touch

- [ ] sh.stat
- [ ] sh.mv
- [ ] sh.cp
	+ [ ] -r
- [ ] sh.mktemp
- [ ] sh.printf
- [ ] sh.echo
- [ ] sh.tee

- [ ] net.ping
- [ ] net.curl

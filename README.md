# nosh
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
- [ ] PWD / cd
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
	+ [ ] redirections
	+ [ ] background
	+ [ ] timeouts
- [X] os.distro/kind/os-type/uname
- [ ] os.pkill <pid> <name-*>
- [ ] os.https://nim-lang.org/docs/dynlib.html
- [ ] os.getTempDir
- [ ] os.sendSignal
- [ ] os.registry (WIN-ONLY)

- [X] fs.find
- [ ] fs.join
- [ ] fs.realpath
- [ ] fs.basename
- [ ] fs.symlink-handling

- [ ] sh.stat
- [ ] sh.mv
- [ ] sh.cp
	+ [ ] -r
- [ ] sh.find
- [ ] sh.mktemp
- [ ] sh.cat/read_file
- [ ] sh.printf
- [ ] sh.echo
- [ ] sh.tee

- [ ] os.clipboard

- [ ] net.ping
- [ ] net.curl

#!/usr/bin/env nosh
# [RUN] ./cycle.sh

# test assert
assert(7+1, 8)
assert(1, 2, xfail=True)

# test starlark features
## @XXX f"...{foo}..." formatting strings don't work yet
assert("{}".format("qwe"), "qwe")

# test fs.find
found_count = 0
for x in fs.find("."):
    found_count+=1
assert(found_count >= 62, True)

# test os.distro
res = os.distro()
assert(res.os in ["linux", "netbsd", "openbsd", "darwin", "windows"])
assert(len(res.arch) >= 3)

# test os.run
res = run(os.executable(), "--version", capture=["stdout"])
assert(res.stdout, "0.0.1")
assert(res.exitCode, 0)

res = run(os.executable(), "--version", capture=["stdout->devnull"])
assert(res.stdout, "0.0.1", xfail=True)
assert(res.exitCode, 0)

res = run(os.executable(), "-c", 'print(args)', 'nosh_first_arg', 'nosh_second_arg')
assert(res.stderr, '["nosh_first_arg", "nosh_second_arg"]\n')
assert(res.exitCode, 0)

# test setenv/getenv/expand
setenv("nosh_test_env", "nosh_test_env_has_content")
assert("nosh_test_env=nosh_test_env_has_content" in getenv())
assert(getenv("nosh_test_env"),  "nosh_test_env_has_content")
assert(expand("$nosh_test_env"), "nosh_test_env_has_content")

# test os.run & expand
res = run(os.executable(), "-c", 'print(expand("$nosh_test_env"))', env=["nosh_test_env=nosh_test_env_has_content"])
assert(res.stderr, "nosh_test_env_has_content\n") # @TODO `print` should not write to stderr but stdout by default
assert(res.exitCode, 0)

# test os.exists
assert(exists(os.executable()))

# test fs.exists/fs.touch/fs.remove - basic
test_dir         = "/tmp/nosh-quicktest/testd/"
test_dir_file    = "/tmp/nosh-quicktest/testd/testf"
test_dir_file_2  = "/tmp/nosh-quicktest/testd/testf2"

touch(test_dir)
assert(exists(test_dir, dir=True, file=False))
touch(test_dir_file)
assert(exists(test_dir_file, dir=False, file=True))
touch(test_dir_file_2)
assert(exists(test_dir_file_2, dir=False, file=True))

remove(test_dir_file_2)
assert(exists(test_dir_file_2, dir=False, file=True), xfail=True)
remove(test_dir)
assert(exists(test_dir_file, dir=False, file=True), xfail=True)
assert(exists(test_dir, dir=True, file=False), xfail=True)

# test fs.remove - force (ignores missing files)
test_dir_missing = "/tmp/nosh-quicktest/testd/missing_directory"
remove(test_dir_missing, force=True)

# test fs.write/fs.read
remove("/tmp/otto_neurath", force=True)
write(path="/tmp/otto_neurath", contents="oh\n")
write(path="/tmp/otto_neurath", contents="woe\n", append=True)
contents = read("/tmp/otto_neurath")
assert(contents, "oh\nwoe\n")
remove("/tmp/otto_neurath")
assert(exists("/tmp/otto_neurath", dir=False, file=True), xfail=True)

# test os.run - redir/append
run(os.executable(), "-c", 'print("otto")', capture=["stderr->/tmp/otto_neurath"])
run(os.executable(), "-c", 'print("neurath")', capture=["stderr->>/tmp/otto_neurath"])
assert(exists("/tmp/otto_neurath", dir=False, file=True))
assert(read("/tmp/otto_neurath"), "otto\nneurath\n")
fs.remove("/tmp/otto_neurath")
assert(exists("/tmp/otto_neurath", dir=False, file=True), xfail=True)

print("quicktest passed successfully")

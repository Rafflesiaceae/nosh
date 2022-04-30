# [RUN] ./cycle.sh

# test assert
assert(7+1, 8)
assert(1, 2, xfail=True)

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

# test setenv/getenv/expand
setenv("nosh_test_env", "nosh_test_env_has_content")
assert("nosh_test_env=nosh_test_env_has_content" in getenv())
assert(getenv("nosh_test_env"),  "nosh_test_env_has_content")
assert(expand("$nosh_test_env"), "nosh_test_env_has_content")

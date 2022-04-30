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

res = run(os.executable(), "--version", capture=[])
assert(res.stdout, "0.0.1", xfail=True)
assert(res.exitCode, 0)

# [RUN] ./cycle.sh

# test assert
assert(7+1, 8)
assert(1, 2, xfail=True)

# test fs.find
found_count = 0
for x in fs.find("."):
    found_count+=1
assert(found_count >= 62, True)
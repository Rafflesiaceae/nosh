package internal

// intset copied from go.starlark.net@v0.0.0-20220328144851-d1966c6b9fcd/starlark/unpack.go:312:8
type intset struct {
	small uint64       // bitset, used if n < 64
	large map[int]bool //    set, used if n >= 64
}

func (is *intset) init(n int) {
	if n >= 64 {
		is.large = make(map[int]bool)
	}
}

func (is *intset) set(i int) (prev bool) {
	if is.large == nil {
		prev = is.small&(1<<uint(i)) != 0
		is.small |= 1 << uint(i)
	} else {
		prev = is.large[i]
		is.large[i] = true
	}
	return
}

func (is *intset) get(i int) bool {
	if is.large == nil {
		return is.small&(1<<uint(i)) != 0
	}
	return is.large[i]
}

func (is *intset) len() int {
	if is.large == nil {
		// Suboptimal, but used only for error reporting.
		len := 0
		for i := 0; i < 64; i++ {
			if is.small&(1<<uint(i)) != 0 {
				len++
			}
		}
		return len
	}
	return len(is.large)
}

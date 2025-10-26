package piscine

func Abort(a, b, c, d, e int) int {
	// first comparison of a with the other values
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if a > d {
		a, d = d, a
	}
	if a > e {
		a, e = e, a
	}
	// first comparison of b with the other values
	if b > c {
		b, c = c, b
	}
	if b > d {
		b, d = d, b
	}
	if b > e {
		b, e = e, b
	}
	// first comparison of c with the other values
	if c > d {
		c, d = d, c
	}
	if c > e {
		c, e = e, c
	}
	// first comparison of d with the other values
	if d > e {
		d, e = e, d
	}

	return c
}

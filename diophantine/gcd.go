package diophantine

func isqrt(x int64) int64 {
	if x < 0 {
		panic("negative sqrt")
	}
	op := x
	res := int64(0)

	// "one" starts at the highest power of four <= than the argument
	one := int64(1) << 62
	for one > op {
		one >>= 2
	}

	for one != 0 {
		if op >= res+one {
			op -= res + one
			res += one * 2
		}
		res >>= 1
		one >>= 2
	}
	return res
}

func gcd2(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcd(terms ...int64) int64 {
	g := terms[0]
	for _, t := range terms[1:] {
		g = gcd2(g, t)
		if g == 1 {
			break
		}
	}
	return g
}

type exGCD struct {
	// Coefficients
	U, V int64

	// Greatest common divisor
	Gcd int64

	// Quotients
	T, S int64
}

func extendedGCD(a, b int64) exGCD {

	s, old_s := int64(0), int64(1)
	t, old_t := int64(1), int64(0)
	r, old_r := b, a

	for r != 0 {
		q := old_r / r
		old_r, r = r, old_r-q*r
		old_s, s = s, old_s-q*s
		old_t, t = t, old_t-q*t
	}

	return exGCD{
		U:   old_s,
		V:   old_t,
		Gcd: old_r,
		T:   t,
		S:   s,
	}
}

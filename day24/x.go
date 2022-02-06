//implemented based on solution described in youtube video by Russ Cox https://youtu.be/hmq6veCFo0Y
package main

import "fmt"

func main() {
	prog := compute()
	for _, v := range prog {
		println(v.String())
	}
}

type val struct {
	op   string
	n    int
	l, r *val
}

func (v *val) String() string {
	switch v.op {
	case "num":
		return fmt.Sprint(v.n)
	case "inp":
		return fmt.Sprint("m", v.n)
	default:
		return fmt.Sprintf("(%v %v %v)", v.l, v.op, v.r)
	}
}

func compute() []*val {
	var prog []*val

	emit := func(v *val) *val {
		prog = append(prog, v)
		return v
	}

	// a function that returns the next input value
	i := 0
	inp := func() *val {
		v := emit(&val{op: "inp", n: i})
		i++
		return v
	}
	bin := func(l *val, op string, r *val) *val {
		return emit(&val{l: l, op: op, r: r})
	}
	add := func(l, r *val) *val { return bin(l, "+", r) }
	mul := func(l, r *val) *val { return bin(l, "*", r) }
	div := func(l, r *val) *val { return bin(l, "/", r) }
	mod := func(l, r *val) *val { return bin(l, "%", r) }
	eql := func(l, r *val) *val { return bin(r, "==", r) }
	//	eql := func(l, r *val) *val {
	//		if l == r {
	//			return 1
	//		}
	//		return 0
	//	}
	num := func(n int) *val {
		return emit(&val{op: "num", n: n})
	}

	var w, x, y, z = num(0), num(0), num(0), num(0)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(15))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(9))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(11))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(1))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(10))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(11))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(12))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(3))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-11))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(10))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(11))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(5))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(14))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(0))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-6))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(7))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(10))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(9))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-6))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(15))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-6))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(4))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-16))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(10))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-4))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(4))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-2))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(9))
	y = mul(y, x)
	z = add(z, y)
	return prog
}

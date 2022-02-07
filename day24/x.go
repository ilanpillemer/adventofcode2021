//implemented based on solution described in youtube video by Russ Cox https://youtu.be/hmq6veCFo0Y
package main

import (
	"fmt"
)

func main() {
	prog := compute()
	opt(prog)
	dump(prog)
}

func opt(prog []*val) {
	for _, v := range prog {
		switch {
		case v.op == "*" && v.l.op == "num" && v.r.op == "num":
			*v = val{op: "num", n: v.l.n * v.r.n, t: v.t}
		case v.op == "+" && v.l.op == "num" && v.r.op == "num":
			*v = val{op: "num", n: v.l.n + v.r.n, t: v.t}
		case v.op == "%" && v.l.op == "num" && v.r.op == "num":
			*v = val{op: "num", n: v.l.n % v.r.n, t: v.t}
		}
	}
}

func dump(prog []*val) {
	count := make(map[*val]int)
	for _, v := range prog {
		count[v.l]++
		count[v.r]++
	}

	str := make(map[*val]string)
	for _, v := range prog {
		var x string
		switch v.op {
		// if its a constant print it out when it appears rather than its "single static assignment name"
		// and dont print our its SSA
		case "inp", "num": // if its an input print it out rather than its "single static assignment name"
			x = v.Init()
		default:
			//if it appears only once then print its assignment, rather than its single static assignment name
			// and dont print our its SSA

			x = fmt.Sprintf("(%v %v %v)", str[v.l], v.op, str[v.r])
			if count[v] >= 2 {
				//if it apears more than once then print out its SSA name, ands its assignment
				fmt.Println(v.Name(), "=", x)
				x = v.Name()
			}

		}
		str[v] = x
	}

	fmt.Println(str[prog[len(prog)-1]])
}

type val struct {
	t    int
	op   string
	n    int
	l, r *val
}

func (v *val) Name() string {
	// print the single static assignment's "name"
	return fmt.Sprint("t", v.t)
}

func (v *val) Init() string {
	// print what it was initially assigned in terms of SSA
	switch v.op {
	case "num":
		return fmt.Sprint(v.n)
	case "inp":
		return fmt.Sprint("m", v.n)
	default:
		return fmt.Sprintf("(%v %v %v)", v.l.Name(), v.op, v.r.Name())
	}
}

func (v *val) String() string {
	return fmt.Sprintf("(%v = %v)", v.Name(), v.Init())
}

func compute() []*val {
	var prog []*val
	t := 0
	emit := func(v *val) *val {
		t++
		v.t = t
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
	eql := func(l, r *val) *val { return bin(l, "==", r) }
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
	_ = z
	return prog
}

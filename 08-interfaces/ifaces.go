package ifaces

import "io"

type Employee struct {
	age int
}

func (e *Employee) SetAge(a int) {
	e.age = a
}

func (e Employee) Age() int {
	return e.age
}

type Customer struct {
	age int
}

func (c *Customer) SetAge(a int) {
	c.age = a
}

func (c Customer) Age() int {
	return c.age
}

type Person interface {
	Age() int
}

func Eldest(args ...Person) int {
	max := 0
	for _, v := range args {
		if v.Age() > max {
			max = v.Age()
		}
	}
	return max
}

func EldestObj(args ...interface{}) interface{} {
	max := 0
	var res interface{}
	for _, v := range args {
		switch o := v.(type) {
		case Customer:
			if o.age > max {
				max = o.age
				res = o
			}
		case Employee:
			if o.age > max {
				max = o.age
				res = o
			}
		}
	}
	return res
}

func printStr(w io.Writer, args ...interface{}) {
	for _, v := range args {
		if o, ok := v.(string); ok {
			w.Write([]byte(o))
		}
	}
}

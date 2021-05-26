package ifaces

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestOlder(t *testing.T) {
	c := Customer{}
	c.SetAge(20)
	e := Employee{}
	e.SetAge(30)
	type args struct {
		args []iface
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"#1", args{[]iface{c}}, 20},
		{"#2", args{[]iface{e}}, 30},
		{"#3", args{[]iface{e, c}}, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Older(tt.args.args...); got != tt.want {
				t.Errorf("Older() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOlderObj(t *testing.T) {
	c := Customer{}
	c.SetAge(20)
	e := Employee{}
	e.SetAge(30)
	var a1, a2, a3 []interface{}
	a1 = append(a1, c)
	a2 = append(a2, e)
	a3 = append(a3, e, c)
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"#1", args{a1}, c},
		{"#2", args{a2}, e},
		{"#3", args{a3}, e},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OlderObj(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OlderObj() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printStr(t *testing.T) {
	var a1, a2, a3 []interface{}
	a1 = append(a1, os.Stdout, 10, 1.4, "qweqwe")
	a2 = append(a2, os.Stdout, 10)
	a3 = append(a3, os.Stdout, 10, "asdqwe", "123123")
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{"#1", args{a1}, "qweqwe"},
		{"#2", args{a2}, ""},
		{"#3", args{a3}, "asdqwe123123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			printStr(w, tt.args.args...)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("printStr() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

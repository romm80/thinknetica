package ifaces

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestEldest(t *testing.T) {
	c := Customer{}
	c.SetAge(20)
	e := Employee{}
	e.SetAge(30)
	type args struct {
		args []Person
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "#1",
			args: args{[]Person{c}},
			want: 20},
		{
			name: "#2",
			args: args{[]Person{e}},
			want: 30,
		},
		{
			name: "#3",
			args: args{[]Person{e, c}},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eldest(tt.args.args...); got != tt.want {
				t.Errorf("Eldest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEldestObj(t *testing.T) {
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
		{
			name: "#1",
			args: args{a1},
			want: c,
		},
		{
			name: "#2",
			args: args{a2},
			want: e,
		},
		{
			name: "#3",
			args: args{a3},
			want: e,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EldestObj(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EldestObj() = %v, want %v", got, tt.want)
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
		{
			name:  "#1",
			args:  args{a1},
			wantW: "qweqwe",
		},
		{
			name:  "#2",
			args:  args{a2},
			wantW: "",
		},
		{
			name:  "#3",
			args:  args{a3},
			wantW: "asdqwe123123",
		},
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

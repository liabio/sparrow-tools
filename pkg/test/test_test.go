package test

import "testing"

func TestA(t *testing.T) {
	var x *struct {
		s [][32]byte
	}
	println(len(x.s[9]))
}

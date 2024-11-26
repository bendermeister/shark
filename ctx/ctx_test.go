package ctx_test

import (
	"shark/ctx"
	"slices"
	"testing"
)

func TestParseString(t *testing.T) {
	in0 := `
[[ctx]]
name = "name"
path = "./../ctx"
tag = ["out", "out:a", "out:b", "out:a:aa"]
`

	cs, err := ctx.ParseString(in0)
	if err != nil {
		t.Fatal(err)
	}

	if len(cs) != 1 {
		t.Fatal("parse something else than 1")
	}

	c := cs[0]

	if c.Name != "name" {
		t.Fatal("name parsed incorrectly")
	}
	if c.Path != "./../ctx" {
		t.Fatal("path parsed incorrectly")
	}

	tags, err := c.Expand("out")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(tags, []string{"out"}) {
		t.Fatal("out expanded incorrectly")
	}

	tags, err = c.Expand("a")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(tags, []string{"out", "a"}) {
		t.Fatal("a expanded incorrectly")
	}

	tags, err = c.Expand("aa")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(tags, []string{"out", "a", "aa"}) {
		t.Fatal("aa expanded incorrectly")
	}

	tags, err = c.Expand("b")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(tags, []string{"out", "b"}) {
		t.Fatal("b expanded incorrectly")
	}
}

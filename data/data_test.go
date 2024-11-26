package data_test

import (
	"shark/ctx"
	"shark/data"
	"slices"
	"testing"
)

func TestParse(t *testing.T) {
	ctx0 := `
[[ctx]]
tag = ["a", "a:b", "a:b:c", "a:aa"]
`

	in0 := `
[[entry]]
title = "title"
desc = "desc"
tag = "c"
value = "1.23"

[[entry]]
title = "title1"
desc = "desc1"
tag = "aa"
value = "12.30"
`

	cs, err := ctx.ParseString(ctx0)
	if err != nil {
		t.Fatal(err)
	}
	c := cs[0]

	entries, err := data.ParseString(&c, in0)
	if err != nil {
		t.Fatal(err)
	}
	e := entries[0]

	if e.Title != "title" {
		t.Fatal("title incorrect")
	}
	if e.Desc != "desc" {
		t.Fatal("desc incorrect")
	}
	if !slices.Equal(e.Tag, []string{"a", "b", "c"}) {
		t.Fatal("tag incorrect")
	}
	if e.Value != 123 {
		t.Fatal("value incorrect")
	}

	e = entries[1]

	if e.Title != "title1" {
		t.Fatal("title incorrect")
	}
	if e.Desc != "desc1" {
		t.Fatal("desc incorrect")
	}
	if !slices.Equal(e.Tag, []string{"a", "aa"}) {
		t.Fatal("tag incorrect")
	}
	if e.Value != 1230 {
		t.Fatal("value incorrect")
	}
}

package data

import (
	"os"
	"shark/ctx"
	"strconv"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type date struct {
	Body time.Time
}

func (a *date) UnmarshalText(text []byte) error {
	var err error
	a.Body, err = time.Parse("2006-01-02", string(text))
	return err
}

type parseValue struct {
	Body int32
}

func (v *parseValue) UnmarshalText(text []byte) error {
	// TODO: what if we parse '12.3' this will be wrongly parsed
	s := strings.Replace(string(text), ".", "", 1)
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}
	v.Body = int32(val)
	return nil
}

type parseEntry struct {
	Date  date       `toml:"date"'`
	Value parseValue `toml:"value"`
	Title string     `toml:"title"`
	Desc  string     `toml:"desc"`
	Tag   string     `toml:"tag"`
}

type Entry struct {
	Date  time.Time
	Value int32
	Title string
	Desc  string
	Tag   []string
}

func ParseString(c *ctx.Ctx, text string) ([]Entry, error) {
	data := struct {
		Entries []parseEntry `toml:"entry"`
	}{}
	_, err := toml.Decode(string(text), &data)
	if err != nil {
		return nil, err
	}
	entries := make([]Entry, len(data.Entries))

	for i, _ := range data.Entries {
		tags, err := c.Expand(data.Entries[i].Tag)
		if err != nil {
			return nil, err
		}
		entries[i] = Entry{
			Title: data.Entries[i].Title,
			Desc:  data.Entries[i].Desc,
			Tag:   tags,
			Date:  data.Entries[i].Date.Body,
			Value: data.Entries[i].Value.Body,
		}
	}
	return entries, nil
}

func ParseFile(c *ctx.Ctx, path string) ([]Entry, error) {
	text, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseString(c, string(text))
}

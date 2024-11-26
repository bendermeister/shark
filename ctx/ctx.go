package ctx

import (
	"strings"

	"github.com/BurntSushi/toml"
)

type Ctx struct {
	Name string
	Path string
	tag  map[string][]string
}

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

var ErrTagExist = &Error{"tag already exist"}

func (c *Ctx) insertTag(t string) error {
	if c.tag == nil {
		c.tag = make(map[string][]string)
	}
	tree := strings.Split(t, ":")
	tail := tree[len(tree)-1]
	_, exists := c.tag[tail]
	if exists {
		return ErrTagExist
	}
	c.tag[tail] = tree
	return nil
}

var ErrTagNoExist = &Error{"tag does not exist"}

func (c *Ctx) Expand(t string) ([]string, error) {
	val, ok := c.tag[t]
	if !ok {
		return nil, ErrTagExist
	}
	return val, nil
}

type parsePath struct {
	Body string
}

func (dir *parsePath) UnmarshalText(text []byte) error {
	// TODO: check if path exist
	dir.Body = string(text)
	return nil
}

type parseCtx struct {
	Name string    `toml:"name"`
	Path parsePath `toml:"path"`
	Tag  []string  `toml:"tag"`
}

func ParseString(text string) ([]Ctx, error) {
	data := struct {
		Ctx []parseCtx `toml:"ctx"`
	}{}
	_, err := toml.Decode(text, &data)
	if err != nil {
		return nil, err
	}

	ctx := make([]Ctx, len(data.Ctx))

	for i, _ := range data.Ctx {
		ctx[i] = Ctx{
			Name: data.Ctx[i].Name,
			Path: data.Ctx[i].Path.Body,
		}

		for _, t := range data.Ctx[i].Tag {
			err = ctx[i].insertTag(t)
			if err != nil {
				return nil, err
			}
		}
	}

	return ctx, nil
}

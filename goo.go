package goo

import (
  "io"
)

type Template struct {
  name string
}


func New(name string) *Template {
  return &Template{
    name: name,
  }
}

func (t *Template) Name() string {
  return t.name
}

func (t *Template) Parse(input string) (*Template, error) {
  return t, nil
}

func (t *Template) Execute(wr io.Writer, data interface{}) (err error) {
  return
}

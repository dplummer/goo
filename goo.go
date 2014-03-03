package goo

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

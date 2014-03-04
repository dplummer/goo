package goo_test

import (
  "testing"
  "bytes"
  "github.com/dplummer/goo"
)

func TestName(t *testing.T) {
  name := "test"
  result := goo.New(name).Name()
  if name != result {
    t.Errorf("%q != %q", name, result)
  }
}

func TestEmpty(t *testing.T) {
  b := new(bytes.Buffer)
  parsed, err := goo.New("empty").Parse("")
  if err != nil {
    t.Errorf("Empty: parse error: %s", err)
  }
  err = parsed.Execute(b, nil)
  if err != nil {
    t.Errorf("Empty: execute error: %s", err)
  }
  result := b.String()
  if "" != result {
    t.Errorf("Empty failed %q != %q", result)
  }
}

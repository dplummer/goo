package goo_test

import (
  "testing"
  "github.com/dplummer/goo"
)

func TestName(t *testing.T) {
  name := "test"
  result := goo.New(name).Name()
  if name != result {
    t.Errorf("%q != %q", name, result)
  }
}

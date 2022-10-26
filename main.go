package main

import (
	"github.com/TimothyStiles/poly/io/genbank"
	"github.com/gopherjs/gopherjs/js"
	"github.com/miratronix/jopher"
	// "github.com/miratronix/jopher"
)

type Genbank struct {
  *js.Object
  Sequence string       `js:"sequence"` // will be changed and include reader, writer, and byte slice.
}
func main() {
  js.Module.Get("exports").Set("genbank", jopher.Promisify(newGenbank))
}

func newGenbank(path string) *js.Object {
  g := Genbank{Object: js.Global.Get("Object").New()}
  file, _ := genbank.Read(path)
  g.Sequence = file.Sequence
  return g.Object
}

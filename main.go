package main

import (
	"github.com/gopherjs/gopherjs/js"
  "github.com/isaacguerreir/poly.js/src/go/genbank"
)

func main() {
  genbankModule := genbank.GenbankFactory()
  js.Module.Get("exports").Set("genbank", genbankModule)
}

package main

import (
	"github.com/gopherjs/gopherjs/js"
  "github.com/isaacguerreir/poly.js/go/genbank"
  "github.com/isaacguerreir/poly.js/go/fasta"
)

func main() {
  genbankModule := genbank.GenbankFactory()
  fastaModule := fasta.FastaFactory()
  js.Module.Get("exports").Set("genbank", genbankModule)
  js.Module.Get("exports").Set("fasta", fastaModule)
}

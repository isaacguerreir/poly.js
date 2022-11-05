package fasta

import (
	"bytes"

	"github.com/TimothyStiles/poly/io/fasta"
	"github.com/gopherjs/gopherjs/js"
	"github.com/miratronix/jopher"
)

type Fasta struct {
  *js.Object
  Name     string `js:"name"`
  Sequence string `js:"sequence"`
}

type FastaModule struct {
  *js.Object
  Read    func(args ...interface{}) *js.Object `js:"read"`
  Parse   func(args ...interface{}) *js.Object `js:"parse"`
}

func FastaFactory() *js.Object {
  module := FastaModule{Object: jopher.NewObject()}
  module.Read = jopher.Promisify(Read)
  module.Parse = jopher.Promisify(Parse)
  return module.Object
}

func Read(path string) []*js.Object {
	data, _ := fasta.Read(path)
	return mapFastas(data)
}

func Parse(buffer []byte) []*js.Object { 
  reader := bytes.NewReader(buffer)
	data, _ := fasta.Parse(reader)
	return mapFastas(data)
}

func mapFastas(fastas []fasta.Fasta) []*js.Object {
  mappedFastas := []*js.Object{}
  for _, fasta := range fastas {
    mappedFasta := mapFasta(fasta).Object
    mappedFastas = append(mappedFastas, mappedFasta)
  }
  return mappedFastas
}

func mapFasta(source fasta.Fasta) Fasta {
  target := Fasta{Object: jopher.NewObject()}
	target.Sequence = source.Sequence
  target.Name =  source.Name
  
  return target 
}

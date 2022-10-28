package genbank

import (
	"fmt"

	"github.com/TimothyStiles/poly/io/genbank"
	"github.com/gopherjs/gopherjs/js"
	"github.com/miratronix/jopher"
)

type GenbankModule struct {
  *js.Object
  Read    func(args ...interface{}) *js.Object `js:"read"`
}

type GenbankJs struct {
	*js.Object
  Meta     MetaJs      `js:"meta"`
	Features []FeatureJs `js:"features"` 
	Sequence string      `js:"sequence"`
}

type MetaJs struct {
  *js.Object
	Date                 string            `js:"date"`
	Definition           string            `js:"definition"`
	Accession            string            `js:"accession"`
	Version              string            `js:"version"`
	Keywords             string            `js:"keywords"`
	Organism             string            `js:"organism"`
	Source               string            `js:"source"`
	Taxonomy             []string          `js:"taxonomy"`
	Origin               string            `js:"origin"`
	Locus                LocusJs           `js:"locus"`
	References           []ReferenceJs     `js:"references"`
	Other                map[string]string `js:"other"`
	Name                 string            `js:"name"`
	SequenceHash         string            `js:"sequence_hash"`
	SequenceHashFunction string            `js:"hash_function"`
}

// Reference holds information for one reference in a Meta struct.
type ReferenceJs struct {
  *js.Object
	Authors string `js:"authors"`
	Title   string `js:"title"`
	Journal string `js:"journal"`
	PubMed  string `js:"pub_med"`
	Remark  string `js:"remark"`
	Range   string `js:"range"`
}

// Locus holds Locus information in a Meta struct.
type LocusJs struct {
  *js.Object
	Name             string `js:"name"`
	SequenceLength   string `js:"sequence_length"`
	MoleculeType     string `js:"molecule_type"`
	GenbankDivision  string `js:"genbank_division"`
	ModificationDate string `js:"modification_date"`
	SequenceCoding   string `js:"sequence_coding"`
	Circular         bool   `js:"circular"`
	Linear           bool   `js:"linear"`
}

type FeatureJs struct {
  *js.Object
	Type                 string            `js:"type"`
	Description          string            `js:"description"`
	Attributes           map[string]string `js:"attributes"`
	SequenceHash         string            `js:"sequence_hash"`
	SequenceHashFunction string            `js:"hash_function"`
	Sequence             string            `js:"sequence"`
	Location             LocationJs          `js:"location"`
}

type LocationJs struct {
  *js.Object
	Start             int        `js:"start"`
	End               int        `js:"end"`
	Complement        bool       `js:"complement"`
	Join              bool       `js:"join"`
	FivePrimePartial  bool       `js:"five_prime_partial"`
	ThreePrimePartial bool       `js:"three_prime_partial"`
	GbkLocationString string     `js:"gbk_location_string"`
	SubLocations      []LocationJs `js:"sub_locations"`
}

func GenbankFactory() *js.Object {
  module := GenbankModule{Object: jopher.NewObject()}
  module.Read = jopher.Promisify(Read)
  return module.Object
}

func Read(path string) *js.Object {
  fmt.Println(path)
	data, _ := genbank.Read(path)
  fmt.Println(data)
	return mapGenbank(data).Object
}

func mapGenbank(source genbank.Genbank) GenbankJs {
  target := GenbankJs{Object: jopher.NewObject()}
	target.Sequence = source.Sequence
  target.Meta = mapMeta(source.Meta)
  target.Features = mapFeatures(source.Features)
  
  return target 
}

func mapFeatures(features []genbank.Feature) []FeatureJs {
  var featuresJs []FeatureJs

  if len(features) == 0 {
    return featuresJs
  }

  for _, feature := range features {
    featureJs := &FeatureJs{Object: jopher.NewObject()}
    featureJs.Type = feature.Type
    featureJs.Description = feature.Description
    featureJs.Attributes = feature.Attributes
    featureJs.SequenceHash = feature.SequenceHash
    featureJs.SequenceHashFunction = feature.SequenceHashFunction
    featureJs.Sequence = feature.Sequence
    featureJs.Location = mapLocation(feature.Location)
    
    featuresJs = append(featuresJs, *featureJs)
  }

  return featuresJs
}

func mapMeta(meta genbank.Meta) MetaJs {
  metaJs := &MetaJs{Object: jopher.NewObject()}

  metaJs.Date                      = meta.Date                                
  metaJs.Definition                = meta.Definition          
  metaJs.Accession                 = meta.Accession           
  metaJs.Version                   = meta.Version             
  metaJs.Keywords                  = meta.Keywords            
  metaJs.Organism                  = meta.Organism            
  metaJs.Source                    = meta.Source              
  metaJs.Taxonomy                  = meta.Taxonomy            
  metaJs.Origin                    = meta.Origin              
  metaJs.Locus                     = mapLocus(meta.Locus)               
  metaJs.References                = mapReferences(meta.References)          
  metaJs.Other                     = meta.Other               
  metaJs.Name                      = meta.Name                
  metaJs.SequenceHash              = meta.SequenceHash        
  metaJs.SequenceHashFunction      = meta.SequenceHashFunction

  return *metaJs
}

func mapLocus(locus genbank.Locus) LocusJs {
  locusJs := &LocusJs{Object: jopher.NewObject()}

  locusJs.Name             = locus.Name            
  locusJs.SequenceLength   = locus.SequenceLength  
  locusJs.MoleculeType     = locus.MoleculeType    
  locusJs.GenbankDivision  = locus.GenbankDivision 
  locusJs.ModificationDate = locus.ModificationDate
  locusJs.SequenceCoding   = locus.SequenceCoding  
  locusJs.Circular         = locus.Circular        
  locusJs.Linear           = locus.Linear          

  return *locusJs
}

func mapReferences(references []genbank.Reference) []ReferenceJs {
  var referencesJs []ReferenceJs
  if len(references) == 0 {
    return referencesJs
  }

  for _, reference := range references {
    referenceJs := &ReferenceJs{Object: jopher.NewObject()}
    referenceJs.Authors   = reference.Authors
    referenceJs.Title     = reference.Title  
    referenceJs.Journal   = reference.Journal
    referenceJs.PubMed    = reference.PubMed 
    referenceJs.Remark    = reference.Remark 
    referenceJs.Range     = reference.Range  
    referencesJs = append(referencesJs, *referenceJs)
  }

  return referencesJs
} 

func mapLocation(location genbank.Location) LocationJs {
  locationJs := &LocationJs{Object: jopher.NewObject()}
  locationJs.Start              = location.Start                              
  locationJs.End                = location.End              
  locationJs.Complement         = location.Complement       
  locationJs.Join               = location.Join             
  locationJs.FivePrimePartial   = location.FivePrimePartial 
  locationJs.ThreePrimePartial  = location.ThreePrimePartial
  locationJs.GbkLocationString  = location.GbkLocationString
  
  var subLocationsJs []LocationJs
  if len(location.SubLocations) == 0 {
    locationJs.SubLocations       = subLocationsJs     
    return *locationJs
  }

  for _, subLocation := range location.SubLocations {
    subLocationsJs = append(subLocationsJs, mapLocation(subLocation))
  }

  locationJs.SubLocations = subLocationsJs     
  return *locationJs
}

/*
  Module and interface declaration for Poly
*/
  declare namespace genbank {
    function read(path: string): Promise<GenbankJs>
  }

  export interface GenbankJs {
    meta: MetaJs;
    features: FeatureJs[];
    sequence: string;
  }
  export interface MetaJs {
    date: string;
    definition: string;
    accession: string;
    version: string;
    keywords: string;
    organism: string;
    source: string;
    taxonomy: string[];
    origin: string;
    locus: LocusJs;
    references: ReferenceJs[];
    other: { [key: string]: string};
    name: string;
    sequenceHash: string;
    sequenceHashFunction: string;
  }
  /**
   * Reference holds information for one reference in a Meta struct.
   */
  export interface ReferenceJs {
    authors: string;
    title: string;
    journal: string;
    pubMed: string;
    remark: string;
    range: string;
  }
  /**
   * Locus holds Locus information in a Meta struct.
   */
  export interface LocusJs {
    name: string;
    sequenceLength: string;
    moleculeType: string;
    genbankDivision: string;
    modificationDate: string;
    sequenceCoding: string;
    circular: boolean;
    linear: boolean;
  }
  export interface FeatureJs {
    type: string;
    description: string;
    attributes: { [key: string]: string};
    sequenceHash: string;
    sequenceHashFunction: string;
    sequence: string;
    location: LocationJs;
  }
  export interface LocationJs {
    start: number /* int */;
    end: number /* int */;
    complement: boolean;
    join: boolean;
    fivePrimePartial: boolean;
    threePrimePartial: boolean;
    gbkLocationString: string;
    subLocations: LocationJs[];
  }

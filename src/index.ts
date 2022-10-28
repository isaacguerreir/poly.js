// import { FeatureJs, genbank, GenbankJs } from './poly'
import { genbank } from "./poly/poly";

genbank.read("./data/benchling.gb")
  .then((data) => console.log(data))


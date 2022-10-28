import { genbank } from '../lib'

genbank.read("./data/benchling.gb")
  .then((res) => console.log(res))


const poly = require('./poly/src/poly')


poly.genbank.read("../data/benchling.gb")
  .then((res) => console.log(res))



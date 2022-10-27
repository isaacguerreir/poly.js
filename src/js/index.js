const { genbank } = require("../../build/poly")

// const obj = importGenbank("../data/benchling.gb") 
genbank.read("./data/benchling.gb")
  .then((res) => { 
    console.log(res)
    // res.features.map((feature) => console.log(feature))
  })
.catch((err) => console.log(err))


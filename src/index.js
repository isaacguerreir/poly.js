const { genbank } = require("../build/poly")

// const obj = importGenbank("../data/benchling.gb") 
genbank("./data/benchling.gb").then((res) => console.log(res)).catch((err) => console.log(err))

import { genbank } from "../src";

(async () => {
  const sequenceData = await genbank.read('./data/benchling.gb')
  console.log(sequenceData)
})()

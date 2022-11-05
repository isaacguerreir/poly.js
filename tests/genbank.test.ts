import { readFileSync } from "fs";
import { genbank } from "../src";

(
  async () => {
    const fromRead = await genbank.read("./data/benchling.gb")
    console.log("From read sequence", fromRead.sequence)

    const buffer = readFileSync('./data/benchling.gb')
    const fromParse = await genbank.parse(new Uint8Array(buffer))
    console.log("From parse sequence: ", fromParse.sequence)
  }
)()

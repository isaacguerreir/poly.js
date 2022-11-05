import { readFileSync } from "fs";
import { fasta } from "../src";

(
  async () => {
    const fromRead = await fasta.read("./data/base.fasta")
    console.log("From read sequence", fromRead)

    const buffer = readFileSync("./data/base.fasta")
    const fromParse = await fasta.parse(new Uint8Array(buffer))
    console.log("From parse sequence: ", fromParse)
  }
)()

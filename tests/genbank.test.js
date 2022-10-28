"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const src_1 = require("../src");
(async () => {
    const sequenceData = await src_1.genbank.read('./data/benchling.gb');
    console.log(sequenceData);
})();

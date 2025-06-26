"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const assert_1 = require("@std/assert");
const main_ts_1 = require("./main.ts");
Deno.test(function addTest() {
    (0, assert_1.assertEquals)((0, main_ts_1.add)(2, 3), 5);
});
//# sourceMappingURL=main_test.js.map
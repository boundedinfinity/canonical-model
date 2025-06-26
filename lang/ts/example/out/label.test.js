"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const node_test_1 = require("node:test");
const label_1 = require("./label");
const uuid_1 = require("uuid");
(0, node_test_1.test)('Generate create table', async () => {
    const repo = new label_1.LabelSqlliteRepository();
    await repo.ensureCreate();
    const label = new label_1.Label({
        id: (0, uuid_1.v4)(),
        name: 'Work'
    });
    await repo.insert(label);
    const found = await repo.select({
        name: { contains: 'Work' }
    });
    console.log(found);
});
//# sourceMappingURL=label.test.js.map
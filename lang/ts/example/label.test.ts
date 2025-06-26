import { test } from 'node:test'
import { Label, LabelSqlliteRepository } from './label'
import { v4 as uuid } from 'uuid';

test('Generate create table', async () => {
    const repo = new LabelSqlliteRepository()
    await repo.ensureCreate()

    const label = new Label({
        id: uuid(),
        name: 'Work'
    })

    await repo.insert(label)

    const found = await repo.select({
        name: { contains: 'Work' }
    })

    console.log(found)
})
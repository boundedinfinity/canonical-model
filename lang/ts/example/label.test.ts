import { test } from 'node:test'
import { Label } from './label'
import { Manager } from './repository'
import { v4 as uuid } from 'uuid';
import { createClient } from "@libsql/client";
import fs from 'node:fs'

// https://bun.sh/docs/api/sqlite
// https://docs.turso.tech/libsql#encryption-at-rest

const dbname = 'bounded.db'

test('Generate create table', async () => {
    // if (fs.existsSync(dbname))
    //     fs.rmSync(dbname)
    const db = createClient({ url: 'file:bounded.db' })
    const manager = new Manager(db)
    await manager.create()

    for (const datum of data)
        await manager.save(new Label(datum))
    db.close()

    const x = 1
})


const data = [
    {
        kind: 'label-full',
        id: '5b87891c-58ab-4a54-9d7c-6fee35ac6b27',
        name: 'Work',
        description: 'Label for work related items'
    },
    {
        kind: 'label-full',
        id: '97c14ec4-c28c-4452-8542-b1444f2cb450',
        name: 'Home',
        description: 'Label for home related items'
    }
]
import { assertEquals } from "@std/assert";
import { BoundedGenerator } from './bounded-gen.js'

function save(file: string, text: string) {
    const encoder = new TextEncoder();
    Deno.writeFileSync(file, encoder.encode(text))
}

function normal(text: string): string {
    return text
        .trim()
        .split('\n')
        .map(line => line.trim())
        .filter(line => line !== '')
        .join('\n')
        .trim()
}

Deno.test('Generate bounded object', () => {
    const actual = new BoundedGenerator().emit(
        {
            kind: 'object',
            name: 'Label',
            properties: [
                {
                    kind: 'string',
                    name: 'id',
                    min: 36,
                    max: 36,
                    anyOf: ['-'],
                    primaryKey: true
                },
                {
                    kind: 'string',
                    name: 'name',
                    min: 2,
                    max: 25,
                    unique: true,
                    indexed: true,
                    searchable: true,
                },
                {
                    kind: 'string',
                    name: 'description',
                    optional: true,
                    min: 2,
                    max: 500,
                    searchable: true
                }
            ]
        },
        {
            kind: 'object',
            name: 'Name Prefix',
            properties: [
                {
                    kind: 'string',
                    name: 'id',
                    min: 36,
                    max: 36,
                    anyOf: ['-'],
                    primaryKey: true
                },
                {
                    kind: 'string',
                    name: 'name',
                    min: 2,
                    max: 50,
                    indexed: true,
                    unique: true,
                },
                {
                    kind: 'array',
                    name: 'abbr',
                    optional: true,
                    min: 1,
                    items: {
                        kind: 'string',
                        min: 1,
                        max: 10,
                        indexed: true,
                        searchable: true
                    }
                },
                {
                    kind: 'string',
                    name: 'description',
                    optional: true,
                    min: 2,
                    max: 500
                },
                {
                    kind: 'array',
                    name: 'labels',
                    optional: true,
                    min: 1,
                    items: {
                        kind: 'ref',
                        ref: 'Label'
                    }
                },
            ]
        }
    )

    for (const [path, content] of Object.entries(actual)) {
        save(path, content)
    }

    // assertEquals(normal(actual), normal(`
    //     Test
    // `))
})


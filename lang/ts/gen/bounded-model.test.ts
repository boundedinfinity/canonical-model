import { assertEquals } from "@std/assert";
import { BoundedGenerator } from './bounded-model.ts'

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
    const actual = new BoundedGenerator().gen(
        {
            kind: 'object',
            name: 'Label',
            properties: [
                {
                    kind: 'string',
                    name: 'id',
                    min: 36,
                    max: 36,
                    anyOf: ['-']
                },
                {
                    kind: 'string',
                    name: 'name',
                    min: 2,
                    max: 50
                },
                {
                    kind: 'string',
                    name: 'description',
                    optional: true,
                    min: 2,
                    max: 500
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
                    anyOf: ['-']
                },
                {
                    kind: 'string',
                    name: 'name',
                    min: 2,
                    max: 50
                },
                {
                    kind: 'array',
                    name: 'abbr',
                    optional: true,
                    min: 1,
                    items: {
                        kind: 'string',
                        min: 1,
                        max: 10
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

    save('./gen-output/generate-bounded-object.gen.ts', actual)

    assertEquals(normal(actual), normal(`
        Test
    `))
})


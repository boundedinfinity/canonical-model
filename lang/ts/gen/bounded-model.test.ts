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
            name: 'Name Prefix',
            properties: [
                {
                    kind: 'string',
                    name: 'id',
                    min: 36,
                    max: 36,
                    includes: '-'
                }
            ]
        }
    )

    save('./gen-output/generate-bounded-object.gen.ts', actual)

    assertEquals(normal(actual), normal(`
        Test
    `))
})


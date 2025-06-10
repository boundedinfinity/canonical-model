import { assertEquals } from "@std/assert";
import { TsGenerator, TsPropertyType } from './ts-gen-model.ts'
import { array } from "zod/v4";

function normal(text: string): string {
    return text
        .trim()
        .split('\n')
        .map(line => line.trim())
        .filter(line => line !== '')
        .join('\n')
        .trim()
}


function save(file: string, text: string) {
    const encoder = new TextEncoder();
    Deno.writeFileSync(file, encoder.encode(text))
}

Deno.test('Generate if statement', () => {
    const actual = new TsGenerator().gen({
        kind: 'if',
        conditions: [`abbr.length == 36`],
        body: ['throw new Error("something is wrong")']
    })

    save('./gen-output/generate-interface.gen.ts', actual)

    assertEquals(normal(actual), normal(`
        export interface NamePrefixInterface {
            id: string
            name: string
            abbr?: string[]
            desc?: string
        }
    `))
})

Deno.test('Generate Interface', () => {
    const actual = new TsGenerator().gen({
        kind: 'interface-type',
        name: "NamePrefixInterface",
        exported: true,
        properties: [
            { kind: 'property', key: 'id', value: { kind: 'string-type' } },
            { kind: 'property', key: 'name', value: { kind: 'string-type' }, },
            { kind: 'property', key: 'abbr', value: { kind: 'string-type', array: true }, optional: true },
            { kind: 'property', key: 'desc', value: { kind: 'string-type' }, optional: true },
        ]
    })

    save('./gen-output/generate-interface.gen.ts', actual)

    assertEquals(normal(actual), normal(`
        export interface NamePrefixInterface {
            id: string
            name: string
            abbr?: string[]
            desc?: string
        }
    `))
})

Deno.test('Generate Class', () => {
    const namePrefixProperties: TsPropertyType[] = [
        {
            kind: 'property',
            key: 'id',
            value: { kind: 'string-type' }
        },
        {
            kind: 'property',
            key: 'name',
            value: { kind: 'string-type' }
        },
        {
            kind: 'property',
            key: 'abbrs',
            optional: true,
            value: { kind: 'string-type', array: true }
        },
        {
            kind: 'property',
            key: 'description',
            optional: true,
            value: { kind: 'string-type' }
        },
    ]


    const actual = new TsGenerator().gen(
        {
            kind: 'class',
            name: "NamePrefix",
            exported: true,
            properties: namePrefixProperties,
            methods: [
                {
                    kind: 'method',
                    name: 'constructor',
                    args: [
                        {
                            kind: 'property',
                            key: 'values',
                            value: {
                                kind: 'object-literal',
                                properties: namePrefixProperties,
                            }
                        }
                    ],
                    body: namePrefixProperties.map(property => {
                        return `this.${property.key} = values.${property.key}`
                    })
                },
            ]
        },
        {
            kind: 'class',
            name: 'NamePrefixValidator',
            methods: [
                {
                    kind: 'method',
                    name: 'validate',
                    args: [
                        {
                            kind: 'property',
                            key: 'namePrefix',
                            value: {
                                kind: 'custom-type',
                                name: 'NamePrefix'
                            }
                        }
                    ],
                    body: [
                        `if(namePrefix.name.length < 2) {
                            throw new Error('name prefix name less than 2 characters')
                        }`
                    ]
                }
            ]
        }
    )

    save('./gen-output/generate-class.gen.ts', actual)

    assertEquals(normal(actual), normal(`
        export class NamePrefix {
            value: NamePrefixInterface

            constructor(
                value: NamePrefixInterface
            ) {
                this.value = value
            }
        }
    `))
})

Deno.test('Generate Enum', () => {
    const actual = new TsGenerator().gen({
        kind: 'enum-type',
        name: "NamePrefixConvertType",
        exported: true,
        values: [
            { name: 'Full', value: 'full' },
            { name: 'Embedded', value: 'embedded' }
        ],
    })

    save('./gen-output/generate-enum.gen.ts', actual)

    assertEquals(normal(actual), normal(`
        export enum NamePrefixConvertType {
            Full = 'full',
            Embedded = 'embedded'
        }
    `))
})
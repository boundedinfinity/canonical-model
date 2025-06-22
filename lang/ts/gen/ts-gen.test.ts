import * as assert from "@std/assert";
import { TsGenerator } from './ts-gen.ts'
import { stringUtils as su } from './utils.ts'

// Deno.test('Test primitive', () => {
//     const gen = new TsGenerator()

//     let actual = gen.emit({ kind: 'primitive', value: 'string' })
//     assertEquals(actual, `'string'`)

//     actual = gen.emit({ kind: 'primitive', value: 1 })
//     assertEquals(actual, '1')
// })

// Deno.test('Test type', () => {
//     const gen = new TsGenerator()

//     let actual = gen.emit({ kind: 'type', type: 'boolean' })
//     assertEquals(actual, `boolean`)

//     actual = gen.emit({ kind: 'type', type: 'number' })
//     assertEquals(actual, 'number')
// })

// Deno.test('Test modifiers', () => {
//     const gen = new TsGenerator()

//     let actual = gen.emit({ kind: 'modifiers', types: [] })
//     assertEquals(actual, ``)

//     actual = gen.emit({ kind: 'modifiers', types: ["export", 'const'] })
//     assertEquals(actual, 'export const ')
// })

// Deno.test('Test property', () => {
//     const gen = new TsGenerator()

//     let actual = gen.emit({ kind: 'property', name: 'arg0', type: { kind: 'type', type: 'string' } })
//     assertEquals(actual, `arg0: string`)

//     actual = gen.emit({ kind: 'property', name: 'arg0', type: { kind: 'type', type: 'string' }, array: false })
//     assertEquals(actual, `arg0: string`)

//     actual = gen.emit({ kind: 'property', name: 'arg0', type: { kind: 'type', type: 'string' }, array: true })
//     assertEquals(actual, `arg0: string[]`)

//     actual = gen.emit({
//         kind: 'property',
//         name: 'arg0',
//         type: { kind: 'type', type: 'string' },
//         array: true,
//         optional: true,
//         variadic: true
//     })
//     assertEquals(actual, `...arg0?: string[]`)

//     actual = gen.emit({
//         kind: 'property',
//         name: 'arg0',
//         type: { kind: 'type', type: 'string' },
//         variadic: true
//     })
//     assertEquals(actual, `...arg0: string[]`)

//     actual = gen.emit({
//         kind: 'property',
//         comments: {
//             kind: 'comment',
//             lines: ['Comment Line 1', 'Comment Line 2']
//         },
//         name: 'arg0',
//         type: { kind: 'type', type: 'string' },
//         variadic: true
//     })
//     assertEquals(actual, `// Comment Line 1\n// Comment Line 2\n...arg0: string[]`)
// })


// Deno.test('Test arguments', () => {
//     const gen = new TsGenerator()

//     let actual = gen.emit({ kind: 'arguments', sep: ', ' })
//     assertEquals(actual, `()`)

//     actual = gen.emit({
//         kind: 'arguments',
//         sep: ', ',
//         items: [
//             {
//                 kind: 'property',
//                 name: 'arg0',
//                 type: { kind: 'type', type: 'string' }
//             }
//         ]
//     })
//     assertEquals(actual, `(arg0: string)`)

//     actual = gen.emit({
//         kind: 'arguments',
//         sep: ', ',
//         items: [
//             {
//                 kind: 'property',
//                 name: 'arg0',
//                 type: { kind: 'type', type: 'string' }
//             },
//             {
//                 kind: 'property',
//                 name: 'arg1',
//                 type: { kind: 'type', type: 'number' },
//                 optional: true,
//                 variadic: true
//             }
//         ]
//     })
//     assertEquals(actual, `(arg0: string, ...arg1?: number[])`)
// })

// Deno.test('Test parameter-object', () => {
//     const gen = new TsGenerator()

//     let actual = gen.emit({
//         kind: 'parameter-object',
//         sep: ', ',
//         items: [
//             {
//                 kind: 'property',
//                 name: 'arg0',
//                 type: { kind: 'type', type: 'string' }
//             },

//         ]
//     })
//     assertEquals(actual, `{arg0: string}`)

//     actual = gen.emit({
//         kind: 'parameter-object',
//         sep: ', ',
//         items: [
//             {
//                 kind: 'property',
//                 name: 'arg0',
//                 type: { kind: 'type', type: 'string' }
//             },
//             {
//                 kind: 'property',
//                 name: 'arg1',
//                 type: { kind: 'type', type: 'number' },
//                 optional: true,
//                 variadic: true
//             }
//         ]
//     })
//     assertEquals(actual, `{arg0: string, ...arg1?: number[]}`)
// })

// Deno.test('Test literal-object', () => {
//     const gen = new TsGenerator()

//     let actual = gen.emit({
//         kind: 'literal-object',
//         sep: ', ',
//         items: [
//             {
//                 name: 'arg0',
//                 value: {
//                     kind: 'literal',
//                     value: { kind: 'primitive', value: 'string0' }
//                 }
//             },
//         ]
//     })
//     assertEquals(actual, `{arg0: 'string0'}`)

//     actual = gen.emit({
//         kind: 'literal-object',
//         sep: ', ',
//         items: [
//             {
//                 name: 'arg0',
//                 value: {
//                     kind: 'literal',
//                     value: { kind: 'primitive', value: 'string0' }
//                 }
//             },
//             {
//                 name: 'arg1',
//                 value: {
//                     kind: 'literal',
//                     value: {
//                         kind: 'literal-object',
//                         sep: ', ',
//                         items: [
//                             {
//                                 name: 'arg0',
//                                 value: {
//                                     kind: 'literal',
//                                     value: { kind: 'primitive', value: 'string0' }
//                                 }
//                             },
//                         ]
//                     }

//                 }
//             },
//         ]
//     })
//     assertEquals(actual, `{arg0: 'string0', arg1: {arg0: 'string0'}}`)
// })

// Deno.test('Test function-decl', () => {
//     const gen = new TsGenerator()

//     let actual = gen.emit({
//         kind: 'function-decl',
//         modifiers: { kind: 'modifiers', types: ['export'] },
//         name: 'fn0',
//     })
//     assertEquals(actual, `export function fn0()`)


//     actual = gen.emit({
//         kind: 'function-decl',
//         modifiers: { kind: 'modifiers', types: ['export'] },
//         name: 'fn0',
//         body: [
//             `console.log('test')`
//         ]
//     })
//     assertEquals(actual, `export function fn0(){\nconsole.log('test')\n}`)
// })

// Deno.test('Test method-decl', () => {
//     const gen = new TsGenerator()

//     let actual = gen.emit({
//         kind: 'method-decl',
//         modifiers: { kind: 'modifiers', types: ['private'] },
//         name: 'fn0',
//     })
//     assertEquals(actual, `private fn0()`)


//     actual = gen.emit({
//         kind: 'method-decl',
//         modifiers: { kind: 'modifiers', types: ['private'] },
//         name: 'fn0',
//         body: [
//             `console.log('test')`
//         ]
//     })
//     assertEquals(actual, `private fn0(){\nconsole.log('test')\n}`)
// })

Deno.test('Test class-decl', () => {
    const gen = new TsGenerator()

    let actual = gen.emit({
        kind: 'class-decl',
        modifiers: { kind: 'modifiers', types: ['export'] },
        name: 'class0',
        methods: [
            {
                kind: 'method-decl',
                name: 'constructor',
            }
        ]
    })
    assert.assertNotEquals(actual, '')

    actual = gen.emit({
        kind: 'class-decl',
        modifiers: { kind: 'modifiers', types: ['export'] },
        name: 'class0',
        properties: [
            {
                kind: 'property',
                name: 'arg0',
                type: gen.Builder.type('string'),
                array: true,
                optional: true
            }
        ],
        methods: [
            {
                kind: 'method-decl',
                name: 'constructor',
                body: [
                    `console.log('test')`
                ]
            },
            {
                kind: 'method-decl',
                name: 'validate',
                body: [
                    `console.log('test')`
                ]
            }
        ]
    })
    assert.assertNotEquals(actual, '')
    su.test.save('./gen-output/ts-gen/class-decl.ts', actual)
})


// import * as path from "jsr:@std/path";
// import * as fs from 'node:fs';
// import {
//     GeneratorStatement, TsBuiltIn, GeneratorClass, type GeneratorOptions
// } from './gen-model.ts'

// export { type GeneratorStatement, TsBuiltIn as GeneratorBuiltIn }

// interface GeneratorProduceOutput {
//     name: string
//     lines: string[]
// }



// interface GeneratorGenerateOutput {
//     name: string
//     lines: string
// }



// export class Manager {
//     options: GeneratorOptions

//     private line = ""
//     private importMap: { [name: string]: string } = {}

//     constructor(options?: Partial<GeneratorOptions>) {
//         this.options = {
//             tabIndex: 0,
//             outputDir: ".",
//             tabPrefix: "    ",
//             ...options
//         }
//     }

//     save(items: GeneratorStatement[]) {
//         const outputs = this.generate(items)
//         if (!fs.existsSync(this.options.outputDir))
//             Deno.mkdir(this.options.outputDir, { recursive: true })
//         const encoder = new TextEncoder();
//         for (const output of outputs) {
//             Deno.writeFileSync(output.name, encoder.encode(output.lines))
//         }
//     }

//     generate(items: GeneratorStatement[]): GeneratorGenerateOutput[] {
//         const outputs = this.produce(items)
//         return outputs.map(output => {
//             return {
//                 name: output.name,
//                 lines: output.lines.join("\n")
//             }
//         })
//     }

//     produce(items: GeneratorStatement[]): GeneratorProduceOutput[] {
//         const outputs: GeneratorProduceOutput[] = []

//         for (const item of items) {
//             if (typeof item == 'object' && item.kind == 'Class') {
//                 const manager = new GeneratorProduceManager(
//                     this.options,
//                     path.join(this.options.outputDir, `${item.name}.ts`),
//                 )

//                 this.produceTypeInterface(manager, item)
//                 manager.newline()
//                 this.produceClass(manager, item)
//                 manager.newline()
//                 this.produceClassObject(manager, item)
//                 manager.newline()
//                 // this.produceClassJson(item, manager)

//                 outputs.push(manager.output())
//             }
//         }


//         return outputs
//     }




//     private calculateImports(item: GeneratorClass) {
//         if (item.properties) {
//             if (!item.imports)
//                 item.imports = []

//             for (const property of item.properties) {
//                 switch (property.type) {
//                     case TsBuiltIn.String:
//                     case TsBuiltIn.Number:
//                         break
//                     default:
//                         if (property.type) {
//                             item.imports.push(property.type)
//                             item.imports.push(`${property.type}Object`)
//                         }
//                 }
//             }
//         }
//     }


//     private produceTypeInterface(item: GeneratorClass,) {
//         const name = `${item.name}Interface`
//         const generator = new Generator2()

//         generator.generate({
//             kind: 'Interface',
//             name,
//             properties: item.properties
//         },)
//     }

//     private produceImports(item: GeneratorClass) {
//         // if (item.imports) {
//         //     for (const imp of item.imports) {
//         //         const found = this.importMap[imp]
//         //         const importPath = path.relative(item.name, found)

//         //         if (found) {
//         //             manager.writeLine(`import { ${imp} } from '${importPath}'`)
//         //         }
//         //     }

//         //     manager.newline()
//         // }
//     }


//     produceClassObject(item: GeneratorClass,) {
//         const optionsName = `${item.name}ObjectOptions`
//         const objectName = `${item.name}Object`
//         const convertName = `${item.name}ConvertType`
//         const interfaceName = `${item.name}Interface`
//         const generator = new Generator2()

//         generator.generate({
//             kind: 'Enum',
//             name: `${convertName}`,
//             values: ['Full', 'Embedded']
//         })

//         generator.newline()

//         generator.generate({
//             kind: 'Interface',
//             name: optionsName,
//             exported: true,
//             properties: [
//                 { kind: 'Property', name: 'marshalType', type: convertName }
//             ]
//         })

//         generator.newline()

//         generator.generate({
//             kind: 'Class',
//             name: objectName,
//             exported: true,

//         })

//         generator.generate({
//             kind: 'Class',
//             name: objectName,
//             exported: true

//         })

//         generator.generate({
//             'kind': 'Function',
//             name: 'constructor',
//             inputs: [{
//                 kind: 'Property',
//                 name: 'options',
//                 type: optionsName,
//                 partial: true
//             }],
//             body: [
//                 'this.options = options'
//             ]
//         },)

//         generator.newline()

//         generator.generate({
//             'kind': 'Function',
//             name: 'marshal',
//             inputs: [{
//                 kind: 'Property',
//                 name: 'item',
//                 type: item.name
//             }],
//             output: 'object',
//             body: [
//                 {
//                     kind: 'Assignment',
//                     name: 'obj',
//                     variableType: GeneratorVariableType.Const,
//                     statement: {
//                         kind: 'Object',
//                         properties: item.properties?.map(property => {
//                             return {
//                                 kind: 'Property',
//                                 key: `${property.name}`,
//                                 name: `item.${property.name}`
//                             }

//                         })
//                     }
//                 },
//                 'return obj'
//             ]
//         },)


//         generator.newline()


//         generator.generate({
//             'kind': 'Function',
//             name: 'unmarshal',
//             inputs: [{
//                 kind: 'Property',
//                 name: 'item',
//                 type: interfaceName
//             }],
//             output: item.name,
//             body: [
//                 {
//                     kind: 'Return',
//                     statement: {
//                         kind: 'New',
//                         name: item.name,
//                         inputs: [
//                             'item'
//                         ]
//                     }
//                 }
//             ]
//         })
//     }

//     // produceClassJson(item: GeneratorClass) {
//     //     this.importMap[`${item.name}Json`] = manager.name
//     //     if (item.exported) manager.write("export ")
//     //     manager.write(`class ${item.name}Json {`)
//     //     manager.newline()

//     //     {
//     //         manager.indent()
//     //         manager.writeLine(`marshal(item: ${item.name}): string {`)
//     //         manager.indent()

//     //         if (item.properties) {
//     //             for (const property of item.properties) {
//     //                 switch (property.type) {
//     //                     case GeneratorBuiltIn.String:
//     //                     case GeneratorBuiltIn.Number:
//     //                         break
//     //                     default:
//     //                         manager.writeLine(`const ${property.name}Object = new ${property.type}Object()`)

//     //                 }
//     //             }
//     //         }

//     //         manager.writeLine(`const obj = {`)
//     //         manager.indent()

//     //         // manager.writeLine(`kind: '${item.typeKind}',`)

//     //         if (item.properties) {
//     //             for (const property of item.properties) {
//     //                 manager.write(`${property.name}: `)
//     //                 switch (property.type) {
//     //                     case GeneratorBuiltIn.String:
//     //                     case GeneratorBuiltIn.Number:
//     //                         manager.write(`item.${property.name}`)
//     //                         break
//     //                     default:
//     //                         manager.write(`item.${property.name}`)
//     //                         if (property.optional)
//     //                             manager.write("?")
//     //                         manager.write(`.map(${property.name} => ${property.name}Object.marshal(${property.name}))`)
//     //                 }

//     //                 manager.writeLine(',')
//     //             }
//     //         }

//     //         manager.dedent()
//     //         manager.writeLine(`}`)

//     //         manager.newline()
//     //         manager.writeLine(`return JSON.stringify(obj)`)
//     //         manager.dedent()
//     //         manager.writeLine("}")
//     //     }

//     //     manager.newline()

//     //     {
//     //         manager.writeLine(`unmarshal(item: string): ${item.name} {`)
//     //         manager.indent()

//     //         manager.writeLine('const obj = JSON.parse(item)')
//     //         manager.writeLine(`return new ${item.name}(obj)`)

//     //         manager.dedent()
//     //         manager.writeLine(`}`)
//     //     }


//     //     manager.dedent()
//     //     manager.writeLine("}")
//     // }
// }


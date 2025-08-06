// import { Manager, GeneratorStatement, GeneratorBuiltIn as B } from './gen.ts'

// const generator = new Manager({
//   outputDir: './gen-output'
// })
// const inputs: GeneratorStatement[] = [
//   {
//     kind: "Class",
//     name: "NamePrefix",
//     exported: true,
//     constants: ['name-prefix'],
//     properties: [
//       { kind: "Property", name: "id", type: B.String },
//       { kind: "Property", name: "name", type: B.String },
//       { kind: "Property", name: "abbr", type: B.String, optional: true, array: true },
//       { kind: "Property", name: "description", type: B.String, optional: true },
//     ],
//   }, {
//     kind: 'Class',
//     name: "NameSuffix",
//     exported: true,
//     constants: ['name-suffix'],
//     properties: [
//       { kind: "Property", name: "id", type: B.String },
//       { kind: "Property", name: "name", type: B.String },
//       { kind: "Property", name: "abbr", type: B.String, optional: true, array: true },
//       { kind: "Property", name: "description", type: B.String, optional: true },
//     ],
//   },
//   // {
//   //   kind: 'Class',
//   //   name: 'Name',
//   //   exported: true,
//   //   constants: ['name'],
//   //   properties: [
//   //     { kind: 'Property', name: "id", type: B.String },
//   //     { kind: 'Property', name: 'prefix', type: 'NamePrefix', optional: true, array: true },
//   //     { kind: 'Property', name: 'first', type: B.String, optional: true, array: true },
//   //     { kind: 'Property', name: 'middle', type: B.String, optional: true, array: true },
//   //     { kind: 'Property', name: 'last', type: B.String, optional: true, array: true },
//   //     { kind: 'Property', name: 'suffix', type: 'NameSuffix', optional: true, array: true },
//   //   ]
//   // }
// ]

// generator.save(inputs)
import * as BI from './bounded-model.ts'
import { BoundedGenerator } from './bounded-gen.ts'
import { BoundedSqliteGenerator } from './bounded-sqlite-gen.ts'
import * as TS from './ts-gen.ts'
import { TsNamer } from './namer.ts'
import * as npath from "jsr:@std/path";

type EmitValidationOptions = {
    name?: string
}
export class BoundedTsGenerator {
    bounded: BoundedGenerator
    namer = new TsNamer()
    tsGen = new TS.TsGenerator()

    constructor(bounded: BoundedGenerator) {
        this.bounded = bounded
    }

    private invalidTypeErr(type: BI.BoundedType): Error {
        return new Error(`invalid token: ${JSON.stringify(type)}`)
    }

    emit(type: BI.BoundedType): string {
        let result = ''

        switch (type.kind) {
            case 'object':
                result = this.tsGen.emit({
                    kind: 'class-decl',
                    name: this.emitTsClassName(type),
                    properties: type.properties.map(prop => this.emitProperty(prop)),
                    methods: [
                        this.emitConstructor(type),
                        this.emitValidator(type)
                    ]
                })

                break
            default:
                throw this.invalidTypeErr(type)
        }
        return result
    }

    private emitProperty(type: BI.BoundedType): TS.TsProperty {
        let result: TS.TsProperty

        switch (type.kind) {
            case 'string':
            case 'number':
            case 'boolean':
                result = {
                    kind: 'property',
                    name: this.emitTsVarName(type),
                    type: this.tsGen.Builder.type(type.kind),
                    optional: type.optional,
                }
                break
            case 'integer':
                result = {
                    kind: 'property',
                    name: this.emitTsVarName(type),
                    type: this.tsGen.Builder.type('number'),
                    optional: type.optional,
                }
                break
            case 'array':
                result = this.emitProperty(type.items)
                result.array = true
                break
            case 'ref':
                {
                    const resolved = this.bounded.getType(type)
                    result = {
                        kind: 'property',
                        name: this.emitTsVarName(type),
                        type: {
                            kind: 'type',
                            type: {
                                kind: 'raw',
                                value: this.emitTsClassName(resolved)
                            }
                        }
                    }
                }
                break
            default:
                throw this.invalidTypeErr(type)
        }

        return result
    }

    private emitConstructor(type: BI.BoundedObject): TS.TsMethodDecl {
        const body: string[] = []

        if (type.properties.length > 0)
            body.push('\n')

        for (const prop of type.properties) {
            const name = this.emitTsVarName(prop)
            body.push(`this.${name} = args.${name}\n`)
        }

        if (type.properties.length > 0)
            body.push('\n')

        const ctor: TS.TsMethodDecl = {
            kind: 'method-decl',
            name: 'constructor',
            args: {
                kind: 'arguments',
                sep: ', ',
                items: [
                    {
                        kind: 'property',
                        name: 'args',
                        type: {
                            kind: 'type',
                            type: {
                                kind: 'object-type',
                                sep: ',\n',
                                properties: type.properties.map(prop => this.emitProperty(prop)),
                            }
                        }
                    }
                ],
            },
            body
        }

        return ctor
    }

    private emitValidator(type: BI.BoundedObject): TS.TsMethodDecl {
        const body: string[] = this.emitValidation(type)

        if (body.length > 0)
            body.unshift('\n')

        const ctor: TS.TsMethodDecl = {
            kind: 'method-decl',
            name: 'validate',
            return: 'Error[]',
            body
        }

        return ctor
    }

    private emitValidation(type: BI.BoundedType, options?: Partial<EmitValidationOptions>): string[] {
        const lines: string[] = []
        const name = options?.name || this.emitTsVarName(type)
        const opt = type.optional ? `this.${name} && ` : ''

        const emitErr = (text: string): string => {
            text = text.replaceAll(`'`, `\\'`)
            const result = `   errors.push(new Error('${name} ${text}'))`
            return result
        }


        const required = () => {
            if (!type.optional) {
                lines.push(`if(this.${name}) {`)
                lines.push(emitErr('is required'))
                lines.push(`}`)
                lines.push(`\n`)
            }
        }

        const emitValues = (items: TS.TsPrimitiveLiterals[]) => this.tsGen.Builder
            .list(items)
            .map(item => this.tsGen.emit(item))

        switch (type.kind) {
            case 'object':
                lines.push(`const errors: Error[] = []`)
                lines.push(`\n`)

                for (const prop of type.properties)
                    lines.push(...this.emitValidation(prop))

                lines.push(`\n`)
                lines.push(`return errors`)
                break
            case 'string':
                required()

                if (type.min) {
                    lines.push(`if(${opt}this.${name}.length < ${type.min}) {`)
                    lines.push(emitErr(`length is less than ${type.min}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }

                if (type.max) {
                    lines.push(`if(${opt}this.${name}.length > ${type.max}) {`)
                    lines.push(emitErr(`length is greater than ${type.max}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }

                if (type.endsWith) {
                    lines.push(`if(${opt}this.${name}.endsWith('${type.endsWith})') {`)
                    lines.push(emitErr(`does not end with ${type.endsWith}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }

                if (type.startsWith) {
                    lines.push(`if(${opt}this.${name}.startsWith('${type.startsWith})') {`)
                    lines.push(emitErr(`does not start with ${type.startsWith}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }

                if (type.anyOf) {
                    const values = emitValues(type.anyOf)
                    lines.push(`if(${opt}![${values}].includes(this.${name})) {`)
                    lines.push(emitErr(`is not any of ${values}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }

                if (type.noneOf) {
                    const values = emitValues(type.noneOf)
                    lines.push(`if(${opt}[${values}].includes(this.${name})) {`)
                    lines.push(emitErr(`is any of ${values}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }
                break
            case 'number':
            case 'integer':
                required()

                if (type.min) {
                    lines.push(`if(${opt}this.${name} < ${type.min}) {`)
                    lines.push(emitErr(`is less than ${type.min}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }

                if (type.max) {
                    lines.push(`if(${opt}this.${name} > ${type.max}) {`)
                    lines.push(emitErr(`is greater than ${type.max}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }

                if (type.anyOf) {
                    const values = emitValues(type.anyOf)
                    lines.push(`if(${opt}![${values}].includes(this.${name})) {`)
                    lines.push(emitErr(`is not any of ${values}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }

                if (type.noneOf) {
                    const values = emitValues(type.noneOf)
                    lines.push(`if(${opt}[${values}].includes(this.${name})) {`)
                    lines.push(emitErr(`is any of ${values}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }
                break
            case 'array':
                required()

                if (type.min) {
                    lines.push(`if(${opt}this.${name}.length < ${type.min}) {`)
                    lines.push(emitErr(`length is less than ${type.min}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }

                if (type.max) {
                    lines.push(`if(${opt}this.${name}.length > ${type.max}) {`)
                    lines.push(emitErr(`length is greater than ${type.max}`))
                    lines.push(`}`)
                    lines.push(`\n`)
                }

                {
                    const items = this.emitValidation(type.items, { name: 'item' })
                    if (items.length > 0) {
                        lines.push(`for(const item of ${name}) {`)
                        lines.push(...items)
                        lines.push(`}`)
                    }
                }
                break
            case 'ref':
                required()
                lines.push(`if(${opt}) {`)
                lines.push(`    errors.push(...this.${name}.validate())`)
                lines.push(`}`)
                break
            default:
                throw this.invalidTypeErr(type)
        }

        return lines
    }

    emitFilename(type: BI.BoundedType, ...path: string[]): string {
        let name = this.bounded.resolveType(type)
        name = this.namer.tsClass(name)
        name = `${name}.gen.ts`
        name = npath.join(...path, name)
        return name
    }

    private emitTsClassName(type: BI.BoundedType): string {
        let name = this.bounded.resolveName(type)

        if (name === '')
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`)

        name = this.namer.tsClass(name)
        return name
    }

    private emitTsVarName(type: BI.BoundedType): string {
        let name = this.bounded.resolveName(type)

        if (name === '')
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`)

        name = this.namer.tsVar(name)
        return name
    }
}
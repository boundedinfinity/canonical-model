import { TsClass, TsGenerator, TsPropertyType, TsGen } from './ts-gen-model.ts'
import { TsNamer } from './namer.ts'

type BoundedType = BoundedObject | BoundedRef | BoundedArray | BoundedString

interface BoundedObject {
    kind: 'object'
    name: string
    properties: BoundedType[]
    optional?: boolean
}

interface BoundedRef {
    kind: 'ref'
    name?: string
    optional?: boolean
    ref: string
}

interface BoundedArray {
    kind: 'array'
    name?: string
    items: BoundedType
    optional?: boolean
    min?: number
    max?: number
}

interface BoundedString {
    kind: 'string'
    name: string
    optional?: boolean
    min?: number
    max?: number
    includes?: string
    excludes?: string
    startsWith?: string
    endsWith?: string
}

export class BoundedGenerator {
    private tsNamer = new TsNamer()
    private tsGen = new TsGenerator()
    private typeMap: { [name: string]: BoundedType } = {}

    private resolveType(type: BoundedType): string {
        let name: string | undefined

        switch (type.kind) {
            case 'object':
                name = type.name
                break
            case 'array':
                name = this.resolveType(type.items)
                break
            case 'ref':
                if (type.ref in this.typeMap) {
                    name = this.resolveType(this.typeMap[type.ref])
                }
                break
            case 'string':
                name = type.kind
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        if (!name) {
            throw new Error(`cannot resolve type for ${JSON.stringify(type)}`)
        }

        return name
    }

    private resolveName(type: BoundedType): string {
        let name: string

        switch (type.kind) {
            case 'object':
                name = this.tsNamer.tsClass(type.name)
                break
            case 'array':
                name = type.name ?? this.resolveName(type.items)
                break
            case 'ref':
                name = type.name ?? this.resolveType(type)
                break
            case 'string':
                name = type.name
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        return name
    }

    private tsClass(type: BoundedType): string {
        let name = this.resolveName(type)
        name = this.tsNamer.tsClass(name)
        return name
    }

    private tsVar(type: BoundedType): string {
        let name = this.resolveName(type)
        name = this.tsNamer.tsVar(name)
        return name
    }

    gen(...types: BoundedType[]): string {
        for (const type of types) {
            if (type.name && !(type.name in this.typeMap)) {
                this.typeMap[type.name] = type
            }
        }

        const results: TsGen[] = []

        for (const type of types) {
            const output = this.genType(type)
            results.push(output)
        }

        const text = this.tsGen.gens(results)

        return text
    }

    private genType(type: BoundedType): TsGen {
        let result: TsGen

        switch (type.kind) {
            case 'object':
                result = this.genObject(type)
                break
            // case 'array':
            //     this.genType(type.items)
            //     break
            case 'string':
                result = this.genStringPropery(type)
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        return result
    }


    private genObject(type: BoundedObject): TsClass {
        const result: TsClass = {
            kind: 'class',
            name: this.tsClass(type),
            properties: [],
            methods: []
        }

        for (const property of type.properties) {
            let tsProperty: TsPropertyType

            switch (property.kind) {
                case 'object':
                    tsProperty = {
                        kind: 'property',
                        key: this.tsVar(property),
                        optional: property.optional,
                        value: {
                            kind: 'custom-type',
                            name: this.tsClass(property),
                        }
                    }
                    break
                case 'string':
                    tsProperty = this.genStringPropery(property)
                    break
                default:
                    throw new Error(`invalid type ${JSON.stringify(property)}`)
            }

            result.properties?.push(tsProperty)
        }

        const validations: string[] = []

        for (const property of type.properties) {
            switch (property.kind) {
                case 'object':
                    break
                case 'string':
                    validations.push(this.genStringValidation(property))
                    break
                default:
                    throw new Error(`invalid type ${JSON.stringify(property)}`)
            }
        }

        result.methods?.push({
            kind: 'method',
            name: 'validate',
            body: validations
        })



        return result
    }

    private genStringPropery(type: BoundedString): TsPropertyType {
        return {
            kind: 'property',
            key: this.tsNamer.tsVar(type.name),
            optional: type.optional,
            value: {
                kind: 'string-type'
            }
        }
    }

    private genStringValidation(type: BoundedString): string {
        let result = ``
        const optional = type.optional ? `${this.tsVar(type)} && ` : ''

        if (type.max) {
            result += `
                if(${optional}this.${this.tsVar(type)}.length > ${type.max}) {
                    throw new Error('${this.tsVar(type)} is greater than ${type.max} characters in length')
                }
            `
        }

        if (type.min) {
            result += `
                if(${optional}this.${this.tsVar(type)}.length < ${type.min}) {
                    throw new Error('${this.tsVar(type)} is less than ${type.min} characters in length')
                }
            `
        }

        if (type.includes) {
            result += `
                ${optional}if(!this.${this.tsVar(type)}.includes('${type.includes}')) {
                    throw new Error('${this.tsVar(type)} does not include "${type.includes}"')
                }
            `
        }

        if (type.excludes) {
            result += `
                ${optional}if(this.${this.tsVar(type)}.includes(${type.excludes})) {
                    throw new Error('${this.tsVar(type)} includes "${type.excludes}"')
                }
            `
        }

        if (type.startsWith) {
            result += `
                ${optional}if(!this.${this.tsVar(type)}.startsWith(${type.startsWith})) {
                    throw new Error('${this.tsVar(type)} does not start with "${type.startsWith}"')
                }
            `
        }

        if (type.endsWith) {
            result += `
                ${optional}if(!this.${this.tsVar(type)}.endsWith(${type.endsWith})) {
                    throw new Error('${this.tsVar(type)} does not end with "${type.endsWith}"')
                }
            `
        }

        return result
    }
}


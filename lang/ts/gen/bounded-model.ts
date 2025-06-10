import * as TS from './ts-gen-model.ts'
import { TsNamer } from './namer.ts'


type Prettify<T> = {
    [K in keyof T]: T[K];
} & {};

type BoundedType = Prettify<BoundedObject | BoundedRef | BoundedArray | BoundedString | BoundedNumber>

interface BoundedObject {
    kind: 'object'
    name?: string
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
    name?: string
    optional?: boolean
    min?: number
    max?: number
    anyOf?: string[]
    oneOf?: string[]
    allOf?: string[]
    noneOf?: string[]
    startsWith?: string
    endsWith?: string
    matches?: string
}

interface BoundedNumber {
    kind: 'number'
    name?: string
    optional?: boolean
    min?: number
    max?: number
    anyOf?: number[]
    allOf?: number[]
    noneOf?: number[]
    multipleOf?: number
}

const NAME_NOT_FOUND = '______NAME__NOT__FOUND______'

export class BoundedGenerator {
    private tsNamer = new TsNamer()
    private tsGen = new TS.TsGenerator()
    private typeMap: { [name: string]: BoundedType } = {}

    gen(...types: BoundedType[]): string {
        for (const type of types) {
            if (type.name && !(type.name in this.typeMap)) {
                this.typeMap[type.name] = type
            }
        }

        const tsGens: TS.TsToken[] = []
        for (const type of types)
            tsGens.push(this.genTs(type))

        const result = this.tsGen.gen(...tsGens)

        return result
    }

    private genTs(type: BoundedType, parent?: BoundedType): TS.TsToken {
        let result: TS.TsToken | undefined
        const array = parent && parent.kind === 'array' || false

        switch (type.kind) {
            case 'object':
                result = this.genObject(type)
                break
            case 'array':
                this.genTs(type.items, type)
                break
            case 'string':
                result = { kind: 'string-type', array }
                break
            case 'ref':
                result = {
                    kind: 'custom-type',
                    name: this.resolveTsClass(type),
                    array,
                }
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        if (!result)
            throw new Error(`invalid type ${JSON.stringify(type)}`)

        return result
    }

    private genObject(type: BoundedObject): TS.TsClass {
        const result: TS.TsClass = {
            kind: 'class',
            name: this.resolveTsClass(type),
            properties: type.properties.map(prop => {
                const token: TS.TsToken = {
                    kind: 'property',
                    key: this.resolveTsVar(prop),
                    optional: prop.optional,
                    value: this.genTs(prop)
                }
                return token
            }),
            methods: []
        }

        const validations: string[] = []

        for (const property of type.properties) {
            switch (property.kind) {
                case 'object':
                    break
                case 'string':
                    validations.push(this.genStringValidation(property))
                    break
                case 'array':

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
                name = type.name ?? NAME_NOT_FOUND
                break
            case 'array':
                name = type.name ?? this.resolveName(type.items)
                break
            case 'ref':
                name = type.name ?? this.resolveType(type)
                break
            case 'string':
                name = type.name ?? NAME_NOT_FOUND
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        return name
    }

    private resolveTsClass(type: BoundedType): string {
        let name = this.resolveName(type)

        if (name == NAME_NOT_FOUND)
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`)

        name = this.tsNamer.tsClass(name)
        return name
    }

    private resolveTsVar(type: BoundedType): string {
        let name = this.resolveName(type)

        if (name == NAME_NOT_FOUND)
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`)

        name = this.tsNamer.tsVar(name)

        if ('array' in type && type.array)
            name += '[]'

        return name
    }



    private genStringValidation(type: BoundedString): string {
        let result = ``
        // const optional = type.optional ? `${this.tsVar(type)} && ` : ''

        // if (type.max) {
        //     result += `
        //         if(${optional}this.${this.tsVar(type)}.length > ${type.max}) {
        //             throw new Error('${this.tsVar(type)} is greater than ${type.max} characters in length')
        //         }
        //     `
        // }

        // if (type.min) {
        //     result += `
        //         if(${optional}this.${this.tsVar(type)}.length < ${type.min}) {
        //             throw new Error('${this.tsVar(type)} is less than ${type.min} characters in length')
        //         }
        //     `
        // }

        // if (type.includes) {
        //     result += `
        //         ${optional}if(!this.${this.tsVar(type)}.includes('${type.includes}')) {
        //             throw new Error('${this.tsVar(type)} does not include "${type.includes}"')
        //         }
        //     `
        // }

        // if (type.excludes) {
        //     result += `
        //         ${optional}if(this.${this.tsVar(type)}.includes(${type.excludes})) {
        //             throw new Error('${this.tsVar(type)} includes "${type.excludes}"')
        //         }
        //     `
        // }

        // if (type.startsWith) {
        //     result += `
        //         ${optional}if(!this.${this.tsVar(type)}.startsWith(${type.startsWith})) {
        //             throw new Error('${this.tsVar(type)} does not start with "${type.startsWith}"')
        //         }
        //     `
        // }

        // if (type.endsWith) {
        //     result += `
        //         ${optional}if(!this.${this.tsVar(type)}.endsWith(${type.endsWith})) {
        //             throw new Error('${this.tsVar(type)} does not end with "${type.endsWith}"')
        //         }
        //     `
        // }

        return result
    }
}


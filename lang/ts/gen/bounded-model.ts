import { TsGenerator } from './ts-gen-model.ts'


type Prettify<T> = {
    [K in keyof T]: T[K];
} & {};

export type BoundedType = Prettify<
    BoundedObject |
    BoundedRef |
    BoundedArray |
    BoundedString |
    BoundedNumber |
    BoundedInteger |
    BoundedBoolean
>

export interface BoundedObject {
    kind: 'object'
    parent?: BoundedType
    name?: string
    properties: BoundedType[]
    optional?: boolean
}

export interface BoundedRef {
    kind: 'ref'
    parent?: BoundedType
    name?: string
    optional?: boolean
    ref: string
}

export interface BoundedArray {
    kind: 'array'
    parent?: BoundedType
    name?: string
    items: BoundedType
    optional?: boolean
    min?: number
    max?: number
}

export interface BoundedString {
    kind: 'string'
    parent?: BoundedType
    name?: string
    optional?: boolean
    unique?: boolean
    primaryKey?: boolean
    indexed?: boolean
    min?: number
    max?: number
    anyOf?: string[]
    oneOf?: string[]
    allOf?: string[]
    noneOf?: string[]
    startsWith?: string
    endsWith?: string
    matches?: string
    searchable?: boolean

    // validation?: {
    //     unique?: boolean
    //     optional?: boolean
    //     min?: number
    //     max?: number
    //     anyOf?: string[]
    //     oneOf?: string[]
    //     allOf?: string[]
    //     noneOf?: string[]
    //     startsWith?: string
    //     endsWith?: string
    //     matches?: string
    // },
    // query?: {
    //     primaryKey?: boolean
    //     indexed?: boolean
    //     equal?: boolean
    //     contains?: boolean
    // }
}

export interface BoundedInteger {
    kind: 'integer'
    parent?: BoundedType
    name?: string
    optional?: boolean
    primaryKey?: boolean
    unique?: boolean
    indexed?: boolean
    min?: number
    max?: number
    anyOf?: number[]
    allOf?: number[]
    noneOf?: number[]
    multipleOf?: number
    searchable?: boolean
}

export interface BoundedNumber {
    kind: 'number'
    parent?: BoundedType
    name?: string
    optional?: boolean
    primaryKey?: boolean
    unique?: boolean
    indexed?: boolean
    min?: number
    max?: number
    anyOf?: number[]
    allOf?: number[]
    noneOf?: number[]
    multipleOf?: number
    searchable?: boolean
}

export interface BoundedBoolean {
    kind: 'boolean'
    parent?: BoundedType
    name?: string
    optional?: boolean
    searchable?: boolean
}

type Files = { [path: string]: string }

export class BoundedGenerator {
    private tsGen = new TsGenerator(this)
    private typeMap: { [name: string]: BoundedType } = {}

    gen(...types: BoundedType[]): Files {
        for (const type of types) {
            if (type.name && !(type.name in this.typeMap)) {
                this.typeMap[type.name] = type
            }
        }

        for (const type of Object.values(this.typeMap)) {
            const process = (t: BoundedType, p?: BoundedType) => {
                t.parent = p

                switch (t.kind) {
                    case 'ref':
                        if (!(t.ref in this.typeMap))
                            throw new Error(`invalid type ref ${JSON.stringify(type)}`)
                        break
                    case 'array':
                        process(t.items, t)
                        break
                    case 'boolean':
                    case 'integer':
                    case 'number':
                    case 'string':
                        // Nothing to do
                        break
                    case 'object':
                        for (const prop of t.properties)
                            process(prop, t)
                        break
                    default:
                        throw new Error(`invalid type ${JSON.stringify(type)}`)
                }
            }

            process(type)
        }

        let results: Files = {}

        for (const type of types) {
            results = { ...results, ...this.tsGen.gen(type) }
        }

        return results
    }

    getType(type: BoundedType): BoundedType {
        let found: BoundedType

        switch (type.kind) {
            case 'ref':
                found = this.typeMap[type.ref]
                break
            case 'array':
                found = type.items
                break
            default:
                found = type
        }

        return found
    }

    resolveIsRef(type: BoundedType): boolean {
        let found = false

        switch (type.kind) {
            case 'ref':
                found = true
                break
            case 'array':
                found = this.resolveIsRef(type.items)
                break
        }

        return found
    }

    resolveIsArray(type: BoundedType): boolean {
        let found = false

        switch (type.kind) {
            case 'ref':
                {
                    const rtype = this.getType(type)
                    found = this.resolveIsArray(rtype)
                }
                break
            case 'array':
                found = true
                break
        }

        return found
    }

    resolveType(type: BoundedType): string {
        let name: string | undefined

        switch (type.kind) {
            case 'object':
                name = type.name
                break
            case 'array':
                name = this.resolveType(type.items)
                break
            case 'ref':
                {
                    const rtype = this.getType(type)
                    name = this.resolveType(rtype)
                }
                break
            case 'boolean':
            case 'number':
            case 'integer':
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

    resolveName(type: BoundedType): string {
        let name: string | undefined

        switch (type.kind) {
            case 'object':
                name = type.name
                break
            case 'array':
                name = type.name ?? this.resolveName(type.items)
                break
            case 'ref':
                name = type.name ?? this.resolveType(type)
                break
            case 'boolean':
            case 'number':
            case 'integer':
            case 'string':
                name = type.name
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        if (!name && type.parent)
            name = type.parent.name

        if (!name || name === '')
            throw new Error(`invalid type ${JSON.stringify(type)}`)

        return name
    }
}


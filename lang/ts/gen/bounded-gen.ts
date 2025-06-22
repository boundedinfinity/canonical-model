import { BoundedType } from './bounded-model.ts'
import { BoundedTsGenerator } from './bounded-ts-gen.ts'

type Files = { [path: string]: string }

export class BoundedGenerator {
    private tsGen = new BoundedTsGenerator(this)
    private typeMap: { [name: string]: BoundedType } = {}

    emit(...types: BoundedType[]): Files {
        for (const type of types) {
            if (type.name && !(type.name in this.typeMap)) {
                this.typeMap[type.name] = type
            }
        }

        const process = (t: BoundedType, p?: BoundedType) => {
            console.log(t)
            switch (t.kind) {
                case 'ref':
                    if (!(t.ref in this.typeMap))
                        throw this.invalidTypeErr(t)
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
                    throw this.invalidTypeErr(t)
            }
            t.parent = p
        }


        for (const type of Object.values(this.typeMap)) {
            process(type)
        }

        const files: Files = {}

        for (const type of types) {
            const content = this.tsGen.emit(type)
            const path = this.tsGen.emitFilename(type, './gen-output/ts-gen')
            files[path] = content
        }

        return files
    }

    private invalidTypeErr(type: BoundedType) {
        new Error(`invalid token: ${JSON.stringify(type)}`)
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
import * as BI from './bounded-model.ts'
import { TsNamer } from './namer.ts'
import { BoundedSqliteGenerator } from './bounded-sqlite-gen-model.ts'
import { stringUtils, StringBuffer, indenter } from './utils.ts'

type Files = { [path: string]: string }

export class TsGenerator {
    bounded: BI.BoundedGenerator
    namer = new TsNamer()

    constructor(bounded: BI.BoundedGenerator) {
        this.bounded = bounded
    }

    gen(type: BI.BoundedType): Files {
        const files: Files = {}

        switch (type.kind) {
            case 'object':
                {
                    // const sb1 = new StringBuffer()
                    // this.genTypeObject(type, sb1, false)
                    // files[this.resolveClientPath(type)] = sb1.toString()

                    const sb2 = new StringBuffer()
                    this.genTypeObject(type, sb2, true)
                    files[this.resolveServerPath(type)] = sb2.toString()
                }
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        return files
    }

    private genTypeObject(type: BI.BoundedObject, sb: StringBuffer, isServer: boolean) {
        const indent = indenter(sb)
        const name = this.resolveTsClass(type)

        for (const prop of type.properties) {
            if (this.bounded.resolveIsRef(prop)) {
                const fname = this.resolveClientPath(prop)
                const name = this.resolveType(prop, true)
                sb.l(`import {${name}} from './${fname}'`)
            }
        }

        if (sb.lines.length > 0) {
            sb.nl()
        }

        sb.l(`export class ${name} {`)

        indent(() => {
            indent(() => {
                for (const prop of type.properties) {
                    const pname = this.resolveTsVar(prop)
                    const ptype = this.resolveType(prop)
                    sb.l(`${pname}: ${ptype}`)
                }
            })

            sb.nl()

            indent(() => {
                sb.l(`constructor(args: {`)
                indent(() => {
                    for (const prop of type.properties) {
                        const pname = this.resolveTsVar(prop)
                        const ptype = this.resolveType(prop)
                        sb.l(`${pname}: ${ptype}`)
                    }
                })

                sb.l(`}) {`)

                indent(() => {
                    for (const prop of type.properties) {
                        const pname = this.resolveTsVar(prop)
                        sb.l(`this.${pname} = args.${pname}`)
                    }
                })

                sb.l(`}`)
            })

            sb.nl()

            indent(() => {
                sb.l(`validate(): Error[] {`)

                indent(() => {
                    sb.l(`const errors: Error[] = []`)
                    sb.nl()

                    for (const prop of type.properties) {
                        new TsValidators(this, sb).validate(prop, { vprefix: "this." })
                    }

                    sb.l(`return errors`)
                })

                sb.l(`}`)
            })

            if (isServer) {
                sb.nl()

                indent(() => {
                    sb.l(`sqlCreate(): string {`)

                    indent(() => {
                        sb.l('return `')

                        indent(() => {
                            new BoundedSqliteGenerator(this.bounded, sb).genTable(type)
                        })

                        sb.l('`')
                    })

                    sb.l(`}`)
                })

                indent(() => {
                    sb.l(`sqlUpdate(): string {`)

                    indent(() => {
                        sb.l('return `')

                        indent(() => {
                            new BoundedSqliteGenerator(this.bounded, sb).genUpdate(type)
                        })

                        sb.l('`')
                    })

                    sb.l(`}`)
                })

                indent(() => {
                    sb.l(`sqlSelect(): string {`)

                    indent(() => {
                        sb.l('return `')

                        indent(() => {
                            new BoundedSqliteGenerator(this.bounded, sb).genSelect(type)
                        })

                        sb.l('`')
                    })

                    sb.l(`}`)
                })
            }
        })
        sb.l(`}`)
    }

    resolveClientPath(type: BI.BoundedType): string {
        let name = this.resolveType(type)
        name = this.namer.tsClass(name)
        name = `client/${name}.gen.ts`
        return name
    }

    resolveServerPath(type: BI.BoundedType): string {
        let name = this.resolveType(type)
        name = this.namer.tsClass(name)
        name = `server/${name}.gen.ts`
        return name
    }

    resolveTsClass(type: BI.BoundedType): string {
        let name = this.bounded.resolveName(type)

        if (name === '')
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`)

        name = this.namer.tsClass(name)
        return name
    }

    resolveTsVar(type: BI.BoundedType): string {
        let name = this.bounded.resolveName(type)

        if (name === '')
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`)

        name = this.namer.tsVar(name)

        return name
    }

    resolveType(type: BI.BoundedType, typeOnly?: boolean): string {
        let name: string | undefined
        const rtype = this.bounded.resolveType(type)

        switch (rtype) {
            case 'string':
                name = `string`
                break
            case 'integer':
            case 'number':
                name = `number`
                break
            case 'boolean':
                name = `boolean`
                break
            default:
                name = rtype
        }

        if (!name) {
            throw new Error(`cannot resolve type for ${JSON.stringify(type)}`)
        }

        if (this.bounded.resolveIsArray(type) && !typeOnly)
            name += '[]'

        return name
    }
}

type ValidatorOptions = {
    vprefix?: string
    vname?: string
    dprefix?: string
    dname?: string
}

export class TsValidators {
    bounded: TsGenerator
    sb: StringBuffer
    indent: ReturnType<typeof indenter>
    private sv: TsStringValidators
    private ov: TsObjectValidators
    private av: TsArrayValidators
    private rv: TsRefValidators
    private nv: TsNumberValidators

    constructor(bounded: TsGenerator, sb: StringBuffer) {
        this.bounded = bounded
        this.sb = sb
        this.indent = indenter(sb)
        this.sv = new TsStringValidators(this, bounded, sb)
        this.ov = new TsObjectValidators(this, bounded, sb)
        this.av = new TsArrayValidators(this, bounded, sb)
        this.rv = new TsRefValidators(this, bounded, sb)
        this.nv = new TsNumberValidators(this, bounded, sb)
    }

    validate(type: BI.BoundedType, o?: ValidatorOptions) {
        let results = ''

        switch (type.kind) {
            case 'object':
                this.optional(type, o)
                this.ov.sub(type, o)
                break
            case 'string':
                results += this.optional(type, o)
                results += this.sv.max(type, o)
                results += this.sv.min(type, o)
                results += this.sv.anyOf(type, o)
                break
            case 'array':
                results += this.optional(type, o)
                results += this.av.min(type, o)
                results += this.av.max(type, o)
                results += this.av.items(type, o)
                break
            case 'ref':
                results += this.optional(type, o)
                results += this.rv.validate(type, o)
                break
            case 'integer':
            case 'number':
                results += this.optional(type, o)
                results += this.nv.min(type, o)
                results += this.nv.max(type, o)
                results += this.nv.anyOf(type, o)
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        return results
    }

    optional(type: BI.BoundedType, o?: ValidatorOptions) {
        const vname = this.vname(type, o)
        const dname = this.dname(type, o)

        if (!type.optional) {
            this.sb.l(`if(${vname} === undefined || ${vname} === null) {`)
            {
                this.sb.indent()
                this.sb.l(`errors.push(new Error('${dname} is required'))`)
                this.sb.dedent()
            }
            this.sb.l(`}`)
            this.sb.nl()
        }
    }

    vname(type: BI.BoundedType, o?: ValidatorOptions): string {
        let name = o?.vname ?? this.bounded.resolveTsVar(type)
        if (o?.vprefix) name = o.vprefix + name
        return name
    }

    dname(type: BI.BoundedType, o?: ValidatorOptions): string {
        let name = o?.dname ?? this.bounded.resolveTsVar(type)
        if (o?.dprefix) name = o.dprefix + name
        return name
    }
}

class TsArrayValidators {
    parent: TsValidators
    bounded: TsGenerator
    sb: StringBuffer
    indent: ReturnType<typeof indenter>

    constructor(parent: TsValidators, bounded: TsGenerator, sb: StringBuffer) {
        this.parent = parent
        this.bounded = bounded
        this.sb = sb
        this.indent = indenter(sb)
    }

    min(type: BI.BoundedArray, o?: ValidatorOptions) {
        if (type.min) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                this.sb.l(`if(${vname} && ${vname}.length < ${type.min}) {`)
            else
                this.sb.l(`if(${vname}.length < ${type.min}) {`)
            this.sb.indent()
            this.sb.l(`errors.push(new Error('${dname} length is less than ${type.min} items'))`)
            this.sb.dedent()
            this.sb.l(`}`)
            this.sb.nl()
        }
    }

    max(type: BI.BoundedArray, o?: ValidatorOptions) {
        if (type.max) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                this.sb.l(`if(${vname} && ${vname}.length < ${type.max}) {`)
            else
                this.sb.l(`if(${vname}.length < ${type.max}) {`)
            {
                this.sb.indent()
                this.sb.l(`errors.push(new Error('${dname} length is greater than ${type.max} items'))`)
                this.sb.dedent()
            }
            this.sb.l(`}`)
            this.sb.nl()
        }
    }

    items(type: BI.BoundedArray, o?: ValidatorOptions) {
        const vname = this.parent.vname(type, o)

        if (type.optional)
            this.sb.l(`if(${vname}) {`)
        else
            this.sb.l(`{`)

        {
            this.sb.indent()
            this.sb.l(`for(const item of ${vname}) {`)
            {
                this.sb.indent()
                this.parent.validate(type.items, { vname: 'item' })
                this.sb.dedent()
            }
            this.sb.l(`}`)
            this.sb.dedent()
            this.sb.nl()
        }
        this.sb.l(`}`)
    }
}

class TsObjectValidators {
    parent: TsValidators
    bounded: TsGenerator
    sb: StringBuffer
    indent: ReturnType<typeof indenter>

    constructor(parent: TsValidators, bounded: TsGenerator, sb: StringBuffer) {
        this.parent = parent
        this.bounded = bounded
        this.sb = sb
        this.indent = indenter(sb)
    }

    sub(type: BI.BoundedObject, o?: ValidatorOptions) {
        const vname = this.parent.vname(type, o)
        // const dname = this.parent.dname(type, o)

        if (type.optional)
            this.sb.a(`if(${vname}) `)

        this.sb.a('{')
        this.sb.nl()
        {
            this.sb.indent()
            this.sb.l(`const ${name}Errors = ${vname}.validate()`)
            this.sb.l(`errors.push(...${name}Errors)`)
            this.sb.dedent()
        }
        this.sb.l(`}`)
        this.sb.nl()
    }
}

class TsRefValidators {
    parent: TsValidators
    bounded: TsGenerator
    sb: StringBuffer
    indent: ReturnType<typeof indenter>

    constructor(parent: TsValidators, bounded: TsGenerator, sb: StringBuffer) {
        this.parent = parent
        this.bounded = bounded
        this.sb = sb
        this.indent = indenter(sb)
    }

    validate(type: BI.BoundedRef, o?: ValidatorOptions) {
        const vname = this.parent.vname(type, o)
        // const dname = this.parent.dname(type, o)

        if (type.optional)
            this.sb.a(`if(${vname}) `)

        this.sb.a('{')
        this.sb.nl()
        {
            this.sb.indent()
            this.sb.l(`const ${name}Errors = ${vname}.validate()`)
            this.sb.l(`errors.push(...${name}Errors)`)
            this.sb.dedent()
        }
        this.sb.l(`}`)
        this.sb.nl()
    }
}

class TsNumberValidators {
    parent: TsValidators
    bounded: TsGenerator
    sb: StringBuffer
    indent: ReturnType<typeof indenter>

    constructor(parent: TsValidators, bounded: TsGenerator, sb: StringBuffer) {
        this.parent = parent
        this.bounded = bounded
        this.sb = sb
        this.indent = indenter(sb)
    }

    min(type: BI.BoundedNumber | BI.BoundedInteger, o?: ValidatorOptions) {
        if (type.min || type.min === 0) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                this.sb.l(`if(${vname} && ${vname} < ${type.min}) {`)
            else
                this.sb.l(`if(${vname} < ${type.min}) { `)

            this.indent(() => {
                this.sb.l(`    errors.push(new Error('${dname} length is less than ${type.min}'))`)
            })

            this.sb.l(`}`)
            this.sb.nl()
        }
    }

    max(type: BI.BoundedNumber | BI.BoundedInteger, o?: ValidatorOptions) {
        if (type.max) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                this.sb.l(`if(${vname} && ${vname} > ${type.max}) {`)
            else
                this.sb.l(`if(${vname} > ${type.max}) {`)

            this.indent(() => {
                this.sb.l(`errors.push(new Error('${dname} length is less than ${type.max}'))`)
            })

            this.sb.l(`}`)
            this.sb.nl()
        }
    }

    anyOf(type: BI.BoundedNumber | BI.BoundedInteger, o?: ValidatorOptions) {
        if (type.anyOf) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)
            const values = type.anyOf.join(', ')

            if (type.optional)
                this.sb.l(`if(${name}) {`)
            else
                this.sb.l(`{`)

            this.indent(() => {
                this.sb.l(`let found = false`)

                this.sb.l(`for(const v of [${values}]) {`)
                this.indent(() => {
                    this.sb.l(`if(${vname} == v) {`)

                    this.indent(() => {
                        this.sb.l(`found = true`)
                        this.sb.l(`break`)
                    })

                    this.sb.l(`}`)
                })
                this.sb.l(`}`)

                this.sb.nl()
                this.sb.l(`if(!found) {`)
                this.indent(() => {
                    this.sb.l(`errors.push(new Error('${dname} does not include one of ${values}}'))`)
                })
                this.sb.l(`    }`)
            })
            this.sb.l(`}`)
            this.sb.nl()
        }
    }
}

class TsStringValidators {
    parent: TsValidators
    bounded: TsGenerator
    sb: StringBuffer
    indent: ReturnType<typeof indenter>

    constructor(parent: TsValidators, bounded: TsGenerator, sb: StringBuffer) {
        this.parent = parent
        this.bounded = bounded
        this.sb = sb
        this.indent = indenter(sb)
    }

    min(type: BI.BoundedString, o?: ValidatorOptions) {
        if (type.min) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                this.sb.l(`if(${vname} && ${vname}.length < ${type.min}) {`)
            else
                this.sb.l(`if(${vname}.length < ${type.min}) { `)

            this.indent(() => {
                this.sb.l(`errors.push(new Error('${dname} length is less than ${type.min} characters'))`)
            })

            this.sb.l(`}`)
            this.sb.nl()
        }
    }

    max(type: BI.BoundedString, o?: ValidatorOptions) {
        if (type.max) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                this.sb.l(`if(${vname} && ${vname}.length > ${type.max}) {`)
            else
                this.sb.l(`if(${vname}.length > ${type.max}) {`)

            this.indent(() => {
                this.sb.l(`errors.push(new Error('${dname} length is less than ${type.max} characters'))`)
            })

            this.sb.l(`}`)
            this.sb.nl()
        }
    }

    anyOf(type: BI.BoundedString, o?: ValidatorOptions) {
        if (type.anyOf) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)
            const values = type.anyOf.map(x => stringUtils.doubleQuote(x)).join(', ')

            if (type.optional)
                this.sb.l(`if(${name}) `)

            this.sb.l(`{`)

            this.indent(() => {
                this.sb.l(`let found = false`)
                this.sb.l(`for(const v of [${values}])`)

                this.indent(() => {
                    this.sb.l(`if(${vname}.includes(v)) {`)
                    this.indent(() => {
                        this.sb.l(`found = true`)
                        this.sb.l(`break`)
                    })
                    this.sb.l(`}`)
                })

                this.sb.nl()

                this.sb.l(`if(!found) {`)
                this.indent(() => {
                    this.sb.l(`errors.push(new Error('${dname} does not include one of ${values}}'))`)
                })
                this.sb.l(`}`)
            })

            this.sb.l(`}`)
            this.sb.nl()
        }
    }

    oneOf(type: BI.BoundedString, o?: ValidatorOptions) {
        if (type.oneOf) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)
            const values = type.oneOf.map(x => stringUtils.doubleQuote(x)).join(', ')

            if (type.optional)
                this.sb.l(`if(${name})`)

            this.sb.a(` {`)
            this.sb.nl()

            this.sb.l(`    let found: string[] = []`)
            this.sb.l(`    for(const v of [${values}])`)
            this.sb.l(`    if(${vname}.includes(v)) {`)
            this.sb.l(`        found.push(v)`)
            this.sb.l(`    }`)
            this.sb.l(``)
            this.sb.l(`    switch(found.length) {`)
            this.sb.l(`       case 0:`)
            this.sb.l(`           errors.push(new Error('${dname} does not include one of ${values}}'))`)
            this.sb.l(`           break`)
            this.sb.l(`       case 1:`)
            this.sb.l(`           break`)
            this.sb.l(`       default:`)
            this.sb.l(`           const founds = found.join(',')`)
            this.sb.l(`           errors.push(new Error('${dname} includes ' + founds + ', more than one of ${values}}'))`)
            this.sb.l(`    }`)
            this.sb.l(`}`)
            this.sb.nl()
        }
    }
}
import * as BI from './bounded-model.ts'
import { TsNamer } from './namer.ts'
import { BoundedSqliteGenerator } from './bounded-sqlite-gen-model.ts'
import { stringUtils, StringBuffer } from './utils.ts'

type Files = { [path: string]: string }


function fnwrap(pre: () => void, post: () => void, run: () => void) {
    try {
        pre()
        run()
    } finally {
        post()
    }
}

function indenter(sb: StringBuffer): (run: () => void) => void {
    return function (run: () => void) {
        fnwrap(
            () => sb.indent(),
            () => sb.dedent(),
            run
        )
    }
}

export class TsGenerator {
    bounded: BI.BoundedGenerator
    namer = new TsNamer()
    validator: TsValidators

    constructor(bounded: BI.BoundedGenerator) {
        this.bounded = bounded
        this.validator = new TsValidators(this)
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
        const iw = indenter(sb)
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

        iw(() => {
            iw(() => {
                for (const prop of type.properties) {
                    const pname = this.resolveTsVar(prop)
                    const ptype = this.resolveType(prop)
                    sb.l(`${pname}: ${ptype}`)
                }
            })

            sb.nl()

            iw(() => {
                sb.l(`constructor(args: {`)
                iw(() => {
                    for (const prop of type.properties) {
                        const pname = this.resolveTsVar(prop)
                        const ptype = this.resolveType(prop)
                        sb.l(`${pname}: ${ptype}`)
                    }
                })

                sb.l(`}) {`)

                iw(() => {
                    for (const prop of type.properties) {
                        const pname = this.resolveTsVar(prop)
                        sb.l(`this.${pname} = args.${pname}`)
                    }
                })

                sb.l(`}`)
            })

            sb.nl()

            iw(() => {
                sb.l(`validate(): Error[] {`)

                iw(() => {
                    sb.l(`const errors: Error[] = []`)
                    sb.nl()

                    for (const prop of type.properties) {
                        this.validator.validate(prop, sb, { vprefix: "this." })
                    }

                    sb.nl()
                    sb.l(`return errors`)
                })

                sb.l(`}`)
            })

            if (isServer) {
                sb.nl()

                iw(() => {
                    sb.l(`sqlCreate(): string {`)

                    iw(() => {
                        sb.l('return `')

                        iw(() => {
                            new BoundedSqliteGenerator(this.bounded, sb).genTable(type)
                        })

                        sb.l('`')
                    })

                    sb.l(`}`)
                })

                iw(() => {
                    sb.l(`sqlSelect(): string {`)

                    iw(() => {
                        sb.l('return `')

                        iw(() => {
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
    private sv: TsStringValidators
    private ov: TsObjectValidators
    private av: TsArrayValidators
    private rv: TsRefValidators
    private nv: TsNumberValidators

    constructor(bounded: TsGenerator) {
        this.bounded = bounded
        this.sv = new TsStringValidators(this, bounded)
        this.ov = new TsObjectValidators(this, bounded)
        this.av = new TsArrayValidators(this, bounded)
        this.rv = new TsRefValidators(this, bounded)
        this.nv = new TsNumberValidators(this, bounded)
    }

    validate(type: BI.BoundedType, sb: StringBuffer, o?: ValidatorOptions) {
        let results = ''

        switch (type.kind) {
            case 'object':
                this.optional(type, sb, o)
                this.ov.sub(type, sb, o)
                break
            case 'string':
                results += this.optional(type, sb, o)
                results += this.sv.max(type, sb, o)
                results += this.sv.min(type, sb, o)
                results += this.sv.anyOf(type, sb, o)
                break
            case 'array':
                results += this.optional(type, sb, o)
                results += this.av.min(type, sb, o)
                results += this.av.max(type, sb, o)
                results += this.av.items(type, sb, o)
                break
            case 'ref':
                results += this.optional(type, sb, o)
                results += this.rv.validate(type, sb, o)
                break
            case 'integer':
            case 'number':
                results += this.optional(type, sb, o)
                results += this.nv.min(type, sb, o)
                results += this.nv.max(type, sb, o)
                results += this.nv.anyOf(type, sb, o)
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        return results
    }

    optional(type: BI.BoundedType, sb: StringBuffer, o?: ValidatorOptions) {
        const vname = this.vname(type, o)
        const dname = this.dname(type, o)

        if (!type.optional) {
            sb.l(`if(${vname} === undefined || ${vname} === null) {`)
            {
                sb.indent()
                sb.l(`errors.push(new Error('${dname} is required'))`)
                sb.dedent()
            }
            sb.l(`}`)
            sb.nl()
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

    constructor(parent: TsValidators, bounded: TsGenerator) {
        this.parent = parent
        this.bounded = bounded
    }

    min(type: BI.BoundedArray, sb: StringBuffer, o?: ValidatorOptions) {
        if (type.min) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                sb.l(`if(${vname} && ${vname}.length < ${type.min}) {`)
            else
                sb.l(`if(${vname}.length < ${type.min}) {`)
            sb.indent()
            sb.l(`errors.push(new Error('${dname} length is less than ${type.min} items'))`)
            sb.dedent()
            sb.l(`}`)
            sb.nl()
        }
    }

    max(type: BI.BoundedArray, sb: StringBuffer, o?: ValidatorOptions) {
        if (type.max) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                sb.l(`if(${vname} && ${vname}.length < ${type.max}) {`)
            else
                sb.l(`if(${vname}.length < ${type.max}) {`)
            {
                sb.indent()
                sb.l(`errors.push(new Error('${dname} length is greater than ${type.max} items'))`)
                sb.dedent()
            }
            sb.l(`}`)
            sb.nl()
        }
    }

    items(type: BI.BoundedArray, sb: StringBuffer, o?: ValidatorOptions) {
        const vname = this.parent.vname(type, o)

        if (type.optional)
            sb.l(`if(${vname}) {`)
        else
            sb.l(`{`)

        {
            sb.indent()
            sb.l(`for(const item of ${vname}) {`)
            {
                sb.indent()
                this.parent.validate(type.items, sb, { vname: 'item' })
                sb.dedent()
            }
            sb.l(`}`)
            sb.dedent()
            sb.nl()
        }
        sb.l(`}`)
    }
}

class TsObjectValidators {
    parent: TsValidators
    bounded: TsGenerator

    constructor(parent: TsValidators, bounded: TsGenerator) {
        this.parent = parent
        this.bounded = bounded
    }

    sub(type: BI.BoundedObject, sb: StringBuffer, o?: ValidatorOptions) {
        const vname = this.parent.vname(type, o)
        // const dname = this.parent.dname(type, o)

        if (type.optional)
            sb.a(`if(${vname})`)

        sb.l('{')
        {
            sb.indent()
            sb.l(`const ${name}Errors = ${vname}.validate()`)
            sb.l(`errors.push(...${name}Errors)`)
            sb.dedent()
        }
        sb.l(`}`)
        sb.nl()
    }
}

class TsRefValidators {
    parent: TsValidators
    bounded: TsGenerator

    constructor(parent: TsValidators, bounded: TsGenerator) {
        this.parent = parent
        this.bounded = bounded
    }

    validate(type: BI.BoundedRef, sb: StringBuffer, o?: ValidatorOptions) {
        const vname = this.parent.vname(type, o)
        // const dname = this.parent.dname(type, o)

        if (type.optional)
            sb.a(`if(${vname})`)

        sb.l('{')
        {
            sb.indent()
            sb.l(`const ${name}Errors = ${vname}.validate()`)
            sb.l(`errors.push(...${name}Errors)`)
            sb.dedent()
        }
        sb.l(`}`)
        sb.nl()
    }
}

class TsNumberValidators {
    parent: TsValidators
    bounded: TsGenerator

    constructor(parent: TsValidators, bounded: TsGenerator) {
        this.parent = parent
        this.bounded = bounded
    }

    min(type: BI.BoundedNumber | BI.BoundedInteger, sb: StringBuffer, o?: ValidatorOptions) {
        if (type.min || type.min === 0) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                sb.l(`if(${vname} && ${vname} < ${type.min}) {`)
            else
                sb.l(`if(${vname} < ${type.min}) { `)
            {
                sb.indent()
                sb.l(`    errors.push(new Error('${dname} length is less than ${type.min}'))`)
                sb.dedent()
            }
            sb.l(`}`)
            sb.nl()
        }
    }

    max(type: BI.BoundedNumber | BI.BoundedInteger, sb: StringBuffer, o?: ValidatorOptions) {
        if (type.max) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                sb.l(`if(${vname} && ${vname} > ${type.max}) {`)
            else
                sb.l(`if(${vname} > ${type.max}) {`)

            {
                sb.indent()
                sb.l(`errors.push(new Error('${dname} length is less than ${type.max}'))`)
                sb.dedent()
            }
            sb.l(`}`)
            sb.nl()
        }
    }

    anyOf(type: BI.BoundedNumber | BI.BoundedInteger, sb: StringBuffer, o?: ValidatorOptions) {
        if (type.anyOf) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)
            const values = type.anyOf.join(', ')

            if (type.optional)
                sb.l(`if(${name})`)

            sb.l(` {`)
            {
                sb.indent()
                sb.l(`let found = false`)
                sb.l(`for(const v of [${values}]) {`)
                {
                    sb.indent()
                    sb.l(`       if(${vname} == v) {`)
                    {
                        sb.indent()
                        sb.l(`           found = true`)
                        sb.l(`           break`)
                        sb.dedent()
                    }
                    sb.l(`       }`)
                    sb.l(`    }`)
                    sb.dedent()
                }
                sb.l(``)
                sb.l(`    if(!found) {`)
                sb.l(`        errors.push(new Error('${dname} does not include one of ${values}}'))`)
                sb.l(`    }`)
                sb.dedent()
            }
            sb.l(`}`)
            sb.nl()
        }
    }
}

class TsStringValidators {
    parent: TsValidators
    bounded: TsGenerator

    constructor(parent: TsValidators, bounded: TsGenerator) {
        this.parent = parent
        this.bounded = bounded
    }

    min(type: BI.BoundedString, sb: StringBuffer, o?: ValidatorOptions) {


        if (type.min) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                sb.l(`if(${vname} && ${vname}.length < ${type.min}) {`)
            else
                sb.l(`if(${vname}.length < ${type.min}) { `)
            sb.l(`    errors.push(new Error('${dname} length is less than ${type.min} characters'))`)
            sb.l(`}`)
            sb.nl()
        }
    }

    max(type: BI.BoundedString, sb: StringBuffer, o?: ValidatorOptions) {
        if (type.max) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                sb.l(`if(${vname} && ${vname}.length > ${type.max}) {`)
            else
                sb.l(`if(${vname}.length > ${type.max}) {`)

            sb.l(`    errors.push(new Error('${dname} length is less than ${type.max} characters'))`)
            sb.l(`}`)
            sb.nl()
        }
    }

    anyOf(type: BI.BoundedString, sb: StringBuffer, o?: ValidatorOptions) {


        if (type.anyOf) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)
            const values = type.anyOf.map(x => stringUtils.doubleQuote(x)).join(', ')

            if (type.optional)
                sb.l(`if(${name})`)

            sb.l(` {`)

            sb.l(`    let found = false`)
            sb.l(`    for(const v of [${values}])`)
            sb.l(`    if(${vname}.includes(v)) {`)
            sb.l(`        found = true`)
            sb.l(`        break`)
            sb.l(`    }`)
            sb.l(``)
            sb.l(`    if(!found) {`)
            sb.l(`        errors.push(new Error('${dname} does not include one of ${values}}'))`)
            sb.l(`    }`)
            sb.l(`}`)
            sb.nl()
        }
    }

    oneOf(type: BI.BoundedString, sb: StringBuffer, o?: ValidatorOptions) {


        if (type.oneOf) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)
            const values = type.oneOf.map(x => stringUtils.doubleQuote(x)).join(', ')

            if (type.optional)
                sb.l(`if(${name})`)

            sb.l(` {`)

            sb.l(`    let found: string[] = []`)
            sb.l(`    for(const v of [${values}])`)
            sb.l(`    if(${vname}.includes(v)) {`)
            sb.l(`        found.push(v)`)
            sb.l(`    }`)
            sb.l(``)
            sb.l(`    switch(found.length) {`)
            sb.l(`       case 0:`)
            sb.l(`           errors.push(new Error('${dname} does not include one of ${values}}'))`)
            sb.l(`           break`)
            sb.l(`       case 1:`)
            sb.l(`           break`)
            sb.l(`       default:`)
            sb.l(`           const founds = found.join(',')`)
            sb.l(`           errors.push(new Error('${dname} includes ' + founds + ', more than one of ${values}}'))`)
            sb.l(`    }`)
            sb.l(`}`)
            sb.nl()
        }
    }
}
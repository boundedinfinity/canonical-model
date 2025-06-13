import * as BI from './bounded-model.ts'
import { TsNamer } from './namer.ts'
import { stringUtils } from './utils.ts'

type Files = { [path: string]: string }

export class TsGenerator2 {
    bounded: BI.BoundedGenerator
    namer = new TsNamer()
    validator: TsValidators

    constructor(bounded: BI.BoundedGenerator) {
        this.bounded = bounded
        this.validator = new TsValidators(this)
    }

    gen(type: BI.BoundedType): Files {
        const files: Files = {}
        let content = ''

        switch (type.kind) {
            case 'object':
                content = this.genTypeObject(type)
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        files[this.resolvePath(type)] = content

        return files
    }

    private genTypeObject(type: BI.BoundedObject) {
        const name = this.resolveTsClass(type)
        let content = ''

        for (const prop of type.properties) {
            if (this.bounded.resolveIsRef(prop)) {
                const fname = this.resolvePath(prop)
                const name = this.resolveType(prop, true)
                content += `import {${name}} from './${fname}'\n`
            }
        }

        if (content !== '') {
            content += '\n'
        }

        content += `export class ${name} {\n`

        for (const prop of type.properties) {
            const pname = this.resolveTsVar(prop)
            const ptype = this.resolveType(prop)
            content += `    ${pname}: ${ptype}\n`
        }
        content += `\n`

        content += `    constructor(args: {\n`
        for (const prop of type.properties) {
            const pname = this.resolveTsVar(prop)
            const ptype = this.resolveType(prop)
            content += `        ${pname}: ${ptype}\n`
        }

        content += `    }) {`
        for (const prop of type.properties) {
            const pname = this.resolveTsVar(prop)
            content += `        this.${pname} = args.${pname}\n`
        }
        content += `    }\n\n`


        content += `    validate(): Error[] {\n`
        content += `        const errors: Error[] = []\n\n`
        for (const prop of type.properties) {
            content += this.validator.validate(prop, { vprefix: "this." })
            content += `\n`
        }

        content += `\n`
        content += `        return errors\n`
        content += `    }`

        content += `}`

        return content
    }


    resolvePath(type: BI.BoundedType): string {
        let name = this.resolveType(type)
        name = this.namer.tsClass(name)
        name = `${name}.gen.ts`
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
    bounded: TsGenerator2
    private sv: TsStringValidators
    private ov: TsObjectValidators
    private av: TsArrayValidators
    private rv: TsRefValidators
    private nv: TsNumberValidators

    constructor(bounded: TsGenerator2) {
        this.bounded = bounded
        this.sv = new TsStringValidators(this, bounded)
        this.ov = new TsObjectValidators(this, bounded)
        this.av = new TsArrayValidators(this, bounded)
        this.rv = new TsRefValidators(this, bounded)
        this.nv = new TsNumberValidators(this, bounded)
    }

    validate(type: BI.BoundedType, o?: ValidatorOptions): string {
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

    optional(type: BI.BoundedType, o?: ValidatorOptions): string {
        const vname = this.vname(type, o)
        const dname = this.dname(type, o)
        let text = ''


        if (!type.optional) {
            text += `if(${vname} === undefined || ${vname} === null) {\n`
            text += `    errors.push(new Error('${dname} is required'))\n`
            text += `}\n\n`
        }

        return text
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
    bounded: TsGenerator2

    constructor(parent: TsValidators, bounded: TsGenerator2) {
        this.parent = parent
        this.bounded = bounded
    }

    min(type: BI.BoundedArray, o?: ValidatorOptions): string {
        let text = ''

        if (type.min) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                text += `if(${vname} && ${vname}.length < ${type.min}) {\n`
            else
                text += `if(${vname}.length < ${type.min}) { \n`
            text += `    errors.push(new Error('${dname} length is less than ${type.min} items'))\n`
            text += `}\n\n`
        }

        return text
    }

    max(type: BI.BoundedArray, o?: ValidatorOptions): string {
        let text = ''

        if (type.max) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                text += `if(${vname} && ${vname}.length < ${type.max}) {\n`
            else
                text += `if(${vname}.length < ${type.max}) { \n`
            text += `    errors.push(new Error('${dname} length is greater than ${type.max} items'))\n`
            text += `}\n\n`
        }

        return text
    }

    items(type: BI.BoundedArray, o?: ValidatorOptions): string {
        let text = ''
        const vname = this.parent.vname(type, o)


        if (type.optional)
            text += `if(${vname})`

        text += `{\n`
        text += `   for(const item of ${vname}) {`
        text += this.parent.validate(type.items, { vname: 'item' })
        text += `   }`

        text += `}\n\n`

        return text

    }
}


class TsObjectValidators {
    parent: TsValidators
    bounded: TsGenerator2

    constructor(parent: TsValidators, bounded: TsGenerator2) {
        this.parent = parent
        this.bounded = bounded
    }

    sub(type: BI.BoundedObject, o?: ValidatorOptions): string {
        let text = ''

        const vname = this.parent.vname(type, o)
        const dname = this.parent.dname(type, o)

        if (type.optional)
            text += `if(${vname})`

        text += '{\n'
        text += `    const ${name}Errors = ${vname}.validate()\n`
        text += `    errors.push(...${name}Errors)\n`
        text += `}\n\n`


        return text
    }
}

class TsRefValidators {
    parent: TsValidators
    bounded: TsGenerator2

    constructor(parent: TsValidators, bounded: TsGenerator2) {
        this.parent = parent
        this.bounded = bounded
    }

    validate(type: BI.BoundedRef, o?: ValidatorOptions): string {
        let text = ''

        const vname = this.parent.vname(type, o)
        const dname = this.parent.dname(type, o)

        if (type.optional)
            text += `if(${vname})`

        text += '{\n'
        text += `    const ${name}Errors = ${vname}.validate()\n`
        text += `    errors.push(...${name}Errors)\n`
        text += `}\n`


        return text
    }
}

class TsNumberValidators {
    parent: TsValidators
    bounded: TsGenerator2

    constructor(parent: TsValidators, bounded: TsGenerator2) {
        this.parent = parent
        this.bounded = bounded
    }

    min(type: BI.BoundedNumber, o?: ValidatorOptions): string {
        let text = ''

        if (type.min || type.min === 0) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                text += `if(${vname} && ${vname} < ${type.min}) {\n`
            else
                text += `if(${vname} < ${type.min}) { \n`
            text += `    errors.push(new Error('${dname} length is less than ${type.min}'))\n`
            text += `}\n\n`
        }

        return text
    }

    max(type: BI.BoundedNumber, o?: ValidatorOptions): string {
        let text = ''

        if (type.max) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                text += `if(${vname} && ${vname} > ${type.max}) {\n`
            else
                text += `if(${vname} > ${type.max}) {\n`

            text += `    errors.push(new Error('${dname} length is less than ${type.max}'))\n`
            text += `}\n\n`
        }

        return text
    }

    anyOf(type: BI.BoundedNumber, o?: ValidatorOptions): string {
        let text = ''

        if (type.anyOf) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)
            const values = type.anyOf.join(', ')

            if (type.optional)
                text += `if(${name})`

            text += ` {\n`

            text += `    let found = false\n`
            text += `    for(const v of [${values}]) {\n`
            text += `       if(${vname} == v) {\n`
            text += `           found = true\n`
            text += `           break\n`
            text += `       }\n`
            text += `    }\n`
            text += `\n`
            text += `    if(!found) {\n`
            text += `        errors.push(new Error('${dname} does not include one of ${values}}'))\n`
            text += `    }\n`
            text += `}\n\n`
        }

        return text
    }
}


class TsStringValidators {
    parent: TsValidators
    bounded: TsGenerator2

    constructor(parent: TsValidators, bounded: TsGenerator2) {
        this.parent = parent
        this.bounded = bounded
    }

    min(type: BI.BoundedString, o?: ValidatorOptions): string {
        let text = ''

        if (type.min) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                text += `if(${vname} && ${vname}.length < ${type.min}) {\n`
            else
                text += `if(${vname}.length < ${type.min}) { \n`
            text += `    errors.push(new Error('${dname} length is less than ${type.min} characters'))\n`
            text += `}\n\n`
        }

        return text
    }

    max(type: BI.BoundedString, o?: ValidatorOptions): string {
        let text = ''

        if (type.max) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)

            if (type.optional)
                text += `if(${vname} && ${vname}.length > ${type.max}) {\n`
            else
                text += `if(${vname}.length > ${type.max}) {\n`

            text += `    errors.push(new Error('${dname} length is less than ${type.max} characters'))\n`
            text += `}\n\n`
        }

        return text
    }

    anyOf(type: BI.BoundedString, o?: ValidatorOptions): string {
        let text = ''

        if (type.anyOf) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)
            const values = type.anyOf.map(x => stringUtils.doubleQuote(x)).join(', ')

            if (type.optional)
                text += `if(${name})`

            text += ` {\n`

            text += `    let found = false\n`
            text += `    for(const v of [${values}])\n`
            text += `    if(${vname}.includes(v)) {\n`
            text += `        found = true\n`
            text += `        break\n`
            text += `    }\n`
            text += `\n`
            text += `    if(!found) {\n`
            text += `        errors.push(new Error('${dname} does not include one of ${values}}'))\n`
            text += `    }\n`
            text += `}\n\n`
        }

        return text
    }

    oneOf(type: BI.BoundedString, o?: ValidatorOptions): string {
        let text = ''

        if (type.oneOf) {
            const vname = this.parent.vname(type, o)
            const dname = this.parent.dname(type, o)
            const values = type.oneOf.map(x => stringUtils.doubleQuote(x)).join(', ')

            if (type.optional)
                text += `if(${name})`

            text += ` {\n`

            text += `    let found: string[] = []\n`
            text += `    for(const v of [${values}])\n`
            text += `    if(${vname}.includes(v)) {\n`
            text += `        found.push(v)\n`
            text += `    }\n`
            text += `\n`
            text += `    switch(found.length) {\n`
            text += `       case 0:\n`
            text += `           errors.push(new Error('${dname} does not include one of ${values}}'))\n`
            text += `           break\n`
            text += `       case 1:\n`
            text += `           break\n`
            text += `       default:\n`
            text += `           const founds = found.join(',')`
            text += `           errors.push(new Error('${dname} includes ' + founds + ', more than one of ${values}}'))\n`
            text += `    }\n`
            text += `}\n\n`
        }

        return text
    }
}
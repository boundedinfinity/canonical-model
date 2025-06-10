import _ from 'lodash'
import { Prettify, isKey } from './utils.ts'

export type TsToken = Prettify<
    TsStringType |
    TsNumberType |
    TsBooleanType |
    TsEnumType |
    TsInterfaceType |
    TsUnionType |
    TsCustomType |
    TsTypeUndefined |
    TsTypeAny |
    TsTypeNull |
    TsObjectLiteral |
    TsStringLiteral |
    TsNumberLiteral |
    TsBooleanLiteral |
    TsEnumLiteral |
    TsTypeUndefined |
    TsTypeAny |
    TsTypeNull |
    TsMethodCall |
    TsPropertyType |
    TsClass |
    TsMethod |
    TsVariable |
    TsMethodCall |
    TsIf
>


export type TsAccessControl = 'public' | 'private' | 'protected' | 'readonly'
export type TsVariableScope = 'var' | 'const' | 'let'

export type TsVariable = {
    kind: 'variable'
    scope: TsVariableScope
    name: string
    anotation?: TsToken
    value?: TsToken
}

export type TsOpType = '=' | '<' | '<=' | '>' | '>=' | '!='
export type TsSep = ',' | ':' | '\n' | '|'

export type TsUnionType = {
    kind: 'union'
    values: TsToken[]
    array?: boolean
}

export type TsTypeUndefined = { kind: 'undefined' }
export type TsTypeNull = { kind: 'null' }
export type TsTypeAny = { kind: 'any' }

export type TsStringType = {
    kind: 'string-type'
    array?: boolean
}

export type TsStringLiteral = {
    kind: 'string-literal'
    value: string
}

export type TsNumberType = {
    kind: 'number-type'
    array?: boolean
}

export type TsNumberLiteral = {
    kind: 'number-literal'
    value: number
}

export type TsCustomType = {
    kind: 'custom-type'
    name: string
    array?: boolean
}

export type TsBooleanType = {
    kind: 'boolean-type'
    array?: boolean
}

export type TsBooleanLiteral = {
    kind: 'boolean-literal'
    value: boolean
}

export type TsPropertyType = {
    kind: 'property'
    key: string
    value?: TsToken
    optional?: boolean
}

export type TsInterfaceType = {
    kind: 'interface-type'
    name: string
    properties?: TsPropertyType[]
    partial?: boolean
    exported?: boolean
}

export type TsObjectLiteral = {
    kind: 'object-literal'
    properties?: TsPropertyType[]
}

export type TsEnumType = {
    kind: 'enum-type'
    name: string
    exported?: boolean
    values?: {
        name: string
        value?: string | boolean | number
    }[]
}

export type TsEnumLiteral = {
    kind: 'enum-literal'
    enum: string
    value: string
}

export type TsMethodCall = {
    kind: 'method-call'
    name: string
    method: string
    args?: TsPropertyType[],
}

export type TsMethod = {
    kind: 'method'
    name: string
    accessControl?: TsAccessControl
    args?: TsPropertyType[],
    body?: string[]
}

export type TsClass = Prettify<{
    kind: 'class'
    name: string
    properties?: TsPropertyType[]
    exported?: boolean
    methods?: TsMethod[]
}>

export type TsIf = {
    kind: 'if'
    conditions: string[]
    body: string[]
}

export class TsGenerator {
    gen(...tokens: TsToken[]): string {
        const lines: string[] = []

        for (const token of tokens) {
            let text = ''

            switch (token.kind) {
                case 'method-call':
                    {
                        text = `${token.name}.${token.method}(`
                        text += this.genProperties(token.args).join(',')
                        text += `)`
                    }
                    break
                case 'any':
                    text = 'any' + this.genArray(token)
                    break
                case 'null':
                    text = 'null' + this.genArray(token)
                    break
                case 'undefined':
                    text = 'undefined' + this.genArray(token)
                    break
                case 'variable':
                    text = `${token.scope} ${token.name}`
                    if (token.anotation)
                        text += `: ${this.gen(token.anotation)}`
                    if (token.value)
                        text += ` = ${this.gen(token.value)}`
                    break
                case 'custom-type':
                    text = token.name + this.genArray(token)
                    break
                case 'string-type':
                    text = 'string' + this.genArray(token)
                    break
                case 'boolean-type':
                    text = 'boolean' + this.genArray(token)
                    break
                case 'number-type':
                    text = 'number' + this.genArray(token)
                    break
                case 'string-literal':
                    text = `'${token.value}'`
                    break
                case 'number-literal':
                case 'boolean-literal':
                    text = `${token.value}`
                    break
                case 'enum-type':
                    text = this.genEnum(token)
                    break
                case 'enum-literal':
                    text = `${token.enum}.${token.value}`
                    break
                case 'interface-type':
                    text = this.genInterface(token)
                    break
                case 'object-literal':
                    text = this.genObjectLiteral(token)
                    break
                case 'method':
                    text = this.genMethod(token)
                    break
                case 'class':
                    text = this.genClass(token)
                    break
                case 'property':
                    text = this.genProperty(token)
                    break
                case 'union':
                    lines.push(...this.genUnion(token))
                    break
                case 'if':
                    lines.push(...this.genIf(token))
                    break
                default:
                    throw new Error(`invalid kind ${JSON.stringify(token)}`)
            }
        }


        return lines.join('\n')
    }

    private genIf(token: TsIf): string[] {
        const lines: string[] = []

        lines.push(`if( ${token.conditions.join(' & ')}) {`)
        lines.push(token.body.join('\n'))
        lines.push(`}`)

        return lines
    }

    private genUnion(token: TsUnionType): string[] {
        const unions: string[] = []

        for (const union of token.values) {
            const text = this.gen(union)
            unions.push(text)
        }

        const text = unions.join(" | ") + this.genArray(token)
        return [text]
    }

    private genObjectLiteral(input: TsObjectLiteral): string {
        const properties = this.genProperties(input.properties)

        const text = `{
            ${properties.join(',\n')}
        }
        `
        return text
    }

    private genProperties(tokens?: TsPropertyType[]): string[] {
        const properties: string[] = []

        if (tokens) {
            for (const property of tokens) {
                const output = this.genProperty(property)
                properties.push(output)
            }
        }

        return properties
    }

    private genProperty(token: TsPropertyType): string {
        let text: string = `${token.key}` + this.genOptional(token)

        if (token.value)
            text += `: ${this.gen(token.value)}`

        return text
    }

    private genInterface(token: TsInterfaceType): string {
        const properties = this.genProperties(token.properties)

        const text = `
            ${this.genExported(token)}interface ${token.name} {
                ${properties.join('\n')}
            }
        `

        return text
    }

    private genMethods(methods?: TsMethod[]) {
        const text: string[] = []

        if (methods) {
            for (const method of methods) {
                const output = this.genMethod(method)
                text.push(output)
            }
        }

        return text
    }

    private genMethod(method: TsMethod) {
        const accessControl = this.genAccessControl(method.accessControl)
        const args = this.genProperties(method.args)
        const body = method.body ?? []

        const text = `
            ${accessControl}${method.name}(
                ${args}
            ) {
                ${body.join('\n')}
            }
        `

        return text
    }

    private genClass(token: TsClass): string {
        const properties = this.genProperties(token.properties)
        const methods = this.genMethods(token.methods)

        const result = `
            ${this.genExported(token)}class ${token.name} {
                ${properties.join('\n')}
                ${methods.join('\n')}
            }
        `
        return result
    }

    private genEnum(token: TsEnumType): string {
        const values = token.values?.map(value => {
            const v = value.value ? ` = '${value.value}'` : ''
            return `${value.name}${v}`
        }) ?? []

        const result = `
            ${this.genExported(token)}enum ${token.name} {
                ${values.join(',\n')}
            }
        `

        return result
    }

    private isArray(input: TsToken): boolean {
        return isKey('array', input)
    }

    private genArray(input: TsToken): string {
        return this.isArray(input) ? '[]' : ''
    }

    private isExported(input: TsToken): boolean {
        return isKey('exported', input)
    }

    private genExported(token: TsToken): string {
        return this.isExported(token) ? 'exported ' : ''
    }

    private isOptional(input: TsToken): boolean {
        return isKey('optional', input)
    }

    private genOptional(token: TsToken): string {
        return this.isOptional(token) ? '? ' : ''
    }

    private genAccessControl(input?: TsAccessControl): string {
        let text = ""

        if (input) {
            switch (input) {
                case "public":
                    break
                case "private":
                case "protected":
                case 'readonly':
                    text = `${input} `
                    break
                default:
                    throw new Error(`invalid access control ${input}`)
            }
        }

        return text
    }
}


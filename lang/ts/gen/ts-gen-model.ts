import _ from 'lodash'

type TsAccessControl = 'public' | 'private' | 'protected' | 'readonly'
type TsVariableScope = 'var' | 'const' | 'let'

interface TsVariable {
    kind: 'variable'
    scope: TsVariableScope
    name: string
    anotation?: TsType
    value?: TsType
}

export type TsType =
    TsStringType |
    TsNamedType |
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
    TsPropertyType

export type TsGen =
    TsType |
    TsClass |
    TsMethod |
    TsVariable |
    TsMethodCall

interface TsUnionType {
    kind: 'union'
    values: TsType[]
    array?: boolean
}

interface TsTypeUndefined {
    kind: 'undefined'
}

interface TsTypeNull {
    kind: 'null'
}

interface TsTypeAny {
    kind: 'any'
}

interface TsStringType {
    kind: 'string-type'
    array?: boolean
}

interface TsNamedType {
    kind: 'named-type'
    value: string
    array?: boolean
}

interface TsStringLiteral {
    kind: 'string-literal'
    value: string
}

interface TsNumberType {
    kind: 'number-type'
    array?: boolean
}

interface TsNumberLiteral {
    kind: 'number-literal'
    value: number
}

interface TsCustomType {
    kind: 'custom-type'
    name: string
    array?: boolean
}

interface TsBooleanType {
    kind: 'boolean-type'
    array?: boolean
}

interface TsBooleanLiteral {
    kind: 'boolean-literal'
    value: boolean
}

export interface TsPropertyType {
    kind: 'property'
    key: string
    value?: TsType
    optional?: boolean
}

interface TsInterfaceType {
    kind: 'interface-type'
    name: string
    properties?: TsPropertyType[]
    partial?: boolean
    exported?: boolean
}

interface TsObjectLiteral {
    kind: 'object-literal'
    properties?: TsPropertyType[]
}

export interface TsEnumType {
    kind: 'enum-type'
    name: string
    exported?: boolean
    values?: {
        name: string
        value?: string | boolean | number
    }[]
}

interface TsEnumLiteral {
    kind: 'enum-literal'
    enum: string
    value: string
}

interface TsMethodCall {
    kind: 'method-call'
    name: string
    method: string
    args?: TsPropertyType[],
}

interface TsMethod {
    kind: 'method'
    name: string
    accessControl?: TsAccessControl
    args?: TsPropertyType[],
    body?: string[]
}

export interface TsClass {
    kind: 'class'
    name: string
    properties?: TsPropertyType[]
    exported?: boolean
    methods?: TsMethod[]
}

export class TsGenerator {
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

    gens(inputs?: TsGen[]): string {
        const outputs: string[] = []

        if (inputs) {
            for (const input of inputs) {
                const output = this.gen(input)
                outputs.push(output)
            }
        }

        return outputs.join('\n')
    }

    gen(input?: TsGen): string {
        let text: string = ""

        function markArray(input: TsType, text: string) {
            if ('array' in input && input.array) {
                switch (input.kind) {
                    case 'union':
                        return `$(${text})[]`
                    default:
                        return `${text}[]`
                }
            }

            return text
        }

        if (input) {
            switch (input.kind) {
                case 'method-call':
                    {
                        text = `${input.name}.${input.method}(`
                        text += this.genProperties(input.args).join(',')
                        text += `)`
                    }
                    break
                case 'any':
                    text = markArray(input, 'any')
                    break
                case 'null':
                    text = markArray(input, 'null')
                    break
                case 'undefined':
                    text = markArray(input, 'undefined')
                    break
                case 'variable':
                    text = `${input.scope} ${input.name}`
                    if (input.anotation)
                        text += `: ${this.gen(input.anotation)}`
                    if (input.value)
                        text += ` = ${this.gen(input.value)}`
                    break
                case 'custom-type':
                    text = markArray(input, input.name)
                    break
                case 'string-type':
                    text = markArray(input, 'string')
                    break
                case 'named-type':
                    text = markArray(input, `${input.value}`)
                    break
                case 'boolean-type':
                    text = markArray(input, 'boolean')
                    break
                case 'number-type':
                    text = markArray(input, 'number')
                    break
                case 'string-literal':
                    text = `'${input.value}'`
                    break
                case 'number-literal':
                case 'boolean-literal':
                    text = `${input.value}`
                    break
                case 'enum-type':
                    text = this.genEnum(input)
                    break
                case 'enum-literal':
                    text = `${input.enum}.${input.value}`
                    break
                case 'interface-type':
                    text = this.genInterface(input)
                    break
                case 'object-literal':
                    text = this.genObjectLiteral(input)
                    break
                case 'method':
                    text = this.genMethod(input)
                    break
                case 'class':
                    text = this.genClass(input)
                    break
                case 'property':
                    text = this.genProperty(input)
                    break
                case 'union':
                    {
                        const unions: string[] = []

                        for (const union of input.values) {
                            const text = this.gen(union)
                            unions.push(text)
                        }

                        text = markArray(input, unions.join(" | "))
                    }
                    break
                default:
                    throw new Error(`invalid kind ${JSON.stringify(input)}`)
            }
        }

        return text
    }

    private genObjectLiteral(input: TsObjectLiteral): string {
        const properties: string[] = []

        if (input.properties) {
            for (const property of input.properties) {
                const output = this.genProperty(property)
                properties.push(output)
            }
        }

        const text = `{
            ${properties.join(',\n')}
        }
        `
        return text
    }

    private genProperties(inputs?: TsPropertyType[]): string[] {
        const properties: string[] = []

        if (inputs) {
            for (const property of inputs) {
                const output = this.genProperty(property)
                properties.push(output)
            }
        }

        return properties
    }

    private genProperty(input: TsPropertyType): string {
        let text: string = `${input.key}`

        if (input.optional)
            text += "?"

        if (input.value)
            text += `: ${this.gen(input.value)}`

        return text
    }

    private genInterface(input: TsInterfaceType): string {
        const properties = this.genProperties(input.properties)

        const text = `
            ${input.exported ? 'export ' : ''}interface ${input.name} {
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

    private genClass(klass: TsClass): string {
        const properties = this.genProperties(klass.properties)
        const methods = this.genMethods(klass.methods)

        const result = `
            ${klass.exported ? 'export ' : ''}class ${klass.name} {
                ${properties.join('\n')}
                ${methods.join('\n')}
            }
        `
        return result
    }

    private genEnum(theEnum: TsEnumType): string {
        const values = theEnum.values?.map(value => {
            const v = value.value ? ` = '${value.value}'` : ''
            return `${value.name}${v}`
        }) ?? []

        const result = `
            ${theEnum.exported ? 'export ' : ''}enum ${theEnum.name} {
                ${values.join(',\n')}
            }
        `

        return result
    }
}

type Prettify<T> = {
    [K in keyof T]: T[K];
} & unknown

type TsProperty2 = {
    name: string
    type: string
    optional?: boolean
    array?: boolean
    partial?: boolean
}

export class TsGenerator2 {
    private options: {
        tabTemp: string
        indent: number
    } = {
            tabTemp: '__TAB__',
            indent: 0
        }

    constructor(options?: Partial<typeof this.options>) {
        _.merge(this.options, options)
    }

    objectLit(args: {
        array?: boolean,
        properties?: TsProperty2[]
    }): string {
        const properties = args.properties?.map(arg => {
            const optional = arg.optional ? '?' : ''
            const type = arg.type ? ` : ${arg.type}` : ''
            const array = arg.array ? `[]` : ''
            const result = `${arg.name}${optional}${type}${array}`
            return result
        }) ?? []

        const array = args.array ? '[]' : ''

        return `{ ${properties.join(',\n')} }${array}`
    }

    if(args: {
        compares: {
            lhs: {
                name: string
                optional?: string
                array?: boolean
            }
            op: 'equal' | 'not equal' | 'greater than' | 'less than'
            rhs: {
                name: string
                optional?: string
                array?: boolean
            }
        }
        body: string[]
    }): string {
        const ps = (p: {
            name: string
            optional?: string
            array?: boolean
        }) => {
            return `${name}${p.optional ? '?' : ''}`
        }

        const compares = args.compares.map(c => `${ps(c.lhs)} ${c.op} ${ps(c.rhs)}`)

        return `if(${compares.join(', ')}) {
            ${args.body.join('\n')}
        }`
    }

    class(args: {
        name: string,
        exported?: boolean,
        properties?: TsProperty2[],
        methods?: {
            name: string,
            accessControl?: TsAccessControl,
            method?: boolean
            exported?: boolean
            args?: TsProperty2[],
            return?: string,
            body: string[]
        }[]
    }): string {
        const exported = args.exported ? 'exported ' : ''

        const properties = args.properties?.map(arg => {
            const optional = arg.optional ? '?' : ''
            const array = arg.array ? '[]' : ''
            const type = arg.type ? `: ${arg.type}` : ''
            const result = `${arg.name}${optional}${type}${array}`
            return result
        }) ?? []

        const methods = args.methods?.map(m => this.function({
            ...m, method: true
        })) ?? []

        const result = `${exported}class ${args.name}{
            ${properties.join('\n')}

            ${methods.join('\n')}
        }`

        return result
    }

    genericeCall(args: {
        generic: string,
        name: string
    }) {
        return `${args.generic}<${args.name}>`
    }

    function(args: {
        name: string,
        accessControl?: TsAccessControl,
        method?: boolean
        exported?: boolean
        args?: TsProperty2[],
        return?: string,
        body: string[]
    }): string {
        const margs = args.args?.map(arg => {
            const optional = arg.optional ? '?' : ''
            const array = arg.array ? '[]' : ''
            let type = ''

            if (arg.type) {
                type = arg.type

                if (arg.partial)
                    type = this.genericeCall({
                        generic: 'Partial',
                        name: type,
                    })

                type = ` : ${type}`
            }

            const result = `${arg.name}${optional}${type}${array}`
            return result
        }) ?? []

        const exported = args.exported ? 'exported ' : ''
        const fn = args.method ? '' : 'function '
        let ac = ''

        switch (args.accessControl) {
            case 'public':
                // nothing todo
                break
            case 'private':
            case 'protected':
            case 'readonly':
                ac = `${args.accessControl}`
                break
            default:
            // nothing todo
        }

        const ret = args.return ? ` : ${args.return}` : ''

        const result = `
            ${ac}${exported}${fn}${args.name} (${margs.join(', ')})${ret}{
                ${args.body.join('\n')}
            }   
        `

        return result
    }
}
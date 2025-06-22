export type TsSep = ', ' | '\n' | ',\n' | '\n\n'

export type TsRaw = {
    kind: 'raw',
    value: string
}

export type TsNewLine = {
    kind: 'newline'
}

export type TsPrimitiveLiterals = string | number | boolean | symbol
export type TsPrimitiveLiteral = {
    kind: 'primitive'
    value: TsPrimitiveLiterals
}

export type TsLiteral = {
    kind: 'literal'
    value: TsPrimitiveLiteral | TsObjectLiteral
}

export type TsObjectLiteral = {
    kind: 'literal-object'
    items: {
        name: string
        value: TsLiteral
    }[]
    sep: TsSep
}

export type TsLiteralArray = {
    kind: 'literal-array'
    items: TsLiteral[]
    sep: TsSep
}

export type TsPrimitiveTypes = 'string' | 'number' | 'boolean' | 'symbol'
export type TsPrimitiveType = {
    kind: 'primitive-type',
    type: TsPrimitiveTypes
}

export type TsObjectType = {
    kind: 'object-type'
    properties: TsProperty[]
    sep: TsSep
}

export type TsType = {
    kind: 'type'
    type: TsPrimitiveType | TsObjectType | TsRaw
}

export type TsModifers = {
    kind: 'modifiers'
    types: ('const' | 'export' | 'public' | 'private' | 'static' | 'default')[]
}

export type TsProperty = {
    kind: 'property'
    comments?: TsComments
    name: string
    type: TsType
    array?: boolean
    variadic?: boolean
    optional?: boolean
}

export type TsArguments = {
    kind: 'arguments'
    items?: TsProperty[]
    sep: TsSep
}

export type TsParmeterObject = {
    kind: 'parameter-object',
    items: TsProperty[]
    sep: TsSep
}

export type TsComments = {
    kind: 'comment'
    lines?: string[]
    noNlewline?: boolean
}

export type TsMethodDecl = {
    kind: 'method-decl'
    comments?: TsComments
    modifiers?: TsModifers
    name: string
    args?: TsArguments
    body?: string[]
    return?: string
}

export type TsFunctionDecl = {
    kind: 'function-decl'
    comments?: TsComments
    modifiers?: TsModifers
    name: string
    args?: TsArguments
    body?: string[]
    return?: string
}

export type TsClassDecl = {
    kind: 'class-decl'
    comments?: TsComments
    modifiers?: TsModifers
    name: string
    properties?: TsProperty[]
    methods?: TsMethodDecl[]
}

export type TsEnumValue = {
    kind: 'enum-value'
    comments?: TsComments
    name: string
    value: TsPrimitiveLiteral
}

export type TsEnumDecl = {
    kind: 'enum-decl'
    comments?: TsComments
    modifiers?: TsModifers
    name: string
    values: TsEnumValue[]
}

export type TsImport = {
    kind: 'import'
    symbols?: TsImportSymbol[]
    star?: string
    from: string
}

export type TsImportSymbol = {
    kind: 'import-symbol'
    symbol: string
    as?: string
}

export type TsToken = TsPrimitiveLiteral | TsType | TsProperty | TsArguments | TsParmeterObject |
    TsObjectLiteral | TsModifers | TsLiteral | TsMethodDecl | TsRaw | TsNewLine |
    TsFunctionDecl | TsClassDecl | TsEnumValue | TsComments | TsEnumDecl | TsLiteralArray |
    TsPrimitiveType | TsObjectType | TsImport | TsImportSymbol

export const TokenBuilder = {
    newline: (): TsNewLine => { return { kind: 'newline' } },
    raw: (v: string): TsRaw => { return { kind: 'raw', value: v } },
    type: (v: TsPrimitiveTypes): TsType => {
        return {
            kind: 'type',
            type: { kind: 'primitive-type', type: v }
        }
    },
    primitive: (v: TsPrimitiveLiterals): TsPrimitiveLiteral => { return { kind: 'primitive', value: v } },
    list: (vs: TsPrimitiveLiterals[]): TsLiteral[] => vs.map(v => {
        return {
            kind: 'literal',
            value: TokenBuilder.primitive(v)
        }
    })
}

const utils = {
    isKeyword: (s: string): boolean => TsKeywords.find(x => x == s) !== undefined
}

export type TsGeneratorOptions = {
    keywordPrefix: string
    keywordPrefixIgnores: TsKeyword[]
    numericPrefix: string
}

export class TsGenerator {
    Builder = TokenBuilder
    options: TsGeneratorOptions

    constructor(options?: Partial<TsGeneratorOptions>) {
        this.options = {
            keywordPrefix: '_',
            keywordPrefixIgnores: [],
            numericPrefix: '_',
            ...options
        }
    }

    private wrap(before: string, after: string, sep: TsSep, tokens?: TsToken[]): string {
        const results: string[] = []

        if (tokens) {
            for (const token of tokens) {
                const result = this.emit(token)
                results.push(result)
            }
        }

        const result = `${before}${results.join(sep)}${after}`
        return result
    }

    private invalidTokenErr(token: TsToken) {
        new Error(`invalid token: ${JSON.stringify(token)}`)
    }

    private e = {
        raw: (v: string): string => { return this.emit(TokenBuilder.raw(v)) }
    }

    emit(token?: TsToken): string {
        let result = ''

        if (!token)
            return result

        switch (token.kind) {
            case 'type':
                result = this.emit(token.type)
                break
            case 'primitive-type':
                result = token.type
                break
            case 'object-type':
                result = this.wrap('{', '}', token.sep, token.properties)
                break
            case 'raw':
                result = token.value
                break
            case 'newline':
                result = '\n'
                break
            case 'primitive':
                switch (typeof token.value) {
                    case 'string':
                        result = `'${token.value}'`
                        break
                    case 'number':
                    case 'boolean':
                    case 'symbol':
                        result = String(token.value)
                        break
                    default:
                        throw this.invalidTokenErr(token)
                }
                break
            case 'modifiers':
                result = token.types.join(' ')
                if (result.length > 0)
                    result += " "
                break
            case 'property':
                result += this.emit(token.comments)
                if (token.variadic)
                    result += '...'
                result += token.name
                if (token.optional)
                    result += '?'
                result += `: `
                result += this.emit(token.type)
                if (token.array || token.variadic)
                    result += '[]'
                break
            case 'arguments':
                result += this.wrap('(', ')', token.sep, token.items)
                break
            case 'parameter-object':
                result += this.wrap('{', '}', token.sep, token.items)
                break
            case 'literal':
                result = this.emit(token.value)
                break
            case 'literal-object':
                {
                    const lresults: string[] = []
                    for (const item of token.items) {
                        let lresult = item.name
                        lresult += `: `
                        lresult += this.emit(item.value)
                        lresults.push(lresult)
                    }
                    result += '{' + lresults.join(token.sep) + '}'
                }
                break
            case 'literal-array':
                result += this.wrap('[', ']', ', ', token.items)
                break
            case 'method-decl':
                result += this.emit({ ...token, kind: 'function-decl' })
                result = result.replaceAll(`function `, '')
                break
            case 'function-decl':
                result += this.emit(token.comments)
                result += this.emit(token.modifiers)
                result += `function `
                result += token.name

                if (token.args)
                    result += this.emit(token.args)
                else
                    result += this.emit({ kind: 'arguments', sep: ', ' })

                if (token.return)
                    result += `: ${token.return}`

                {
                    const body: TsRaw[] = token.body?.map(line => {
                        return { kind: 'raw', value: line }
                    }) ?? []
                    result += this.wrap('{', '}', '\n', body)
                }
                break
            case 'class-decl':
                {
                    result += this.emit(token.comments)
                    result += this.emit(token.modifiers)
                    result += `class ${token.name} `

                    const properties: TsToken[] = token.properties
                        ? [...token.properties, this.Builder.newline()]
                        : []


                    const methods: TsToken[] = token.methods
                        ? token.methods.map(m => [m, this.Builder.newline()]).flat()
                        : []


                    result += this.wrap('{', '}', '\n', [...properties, ...methods])
                }
                break
            case 'comment':
                result += token.lines?.map(line => `// ${line}`).join('\n')
                if (!token.noNlewline)
                    result += '\n'
                break
            case 'enum-value':
                result += this.emit(token.comments)
                result += `${token.name} = ${this.emit(token.value)}`
                break
            case 'enum-decl':
                result += this.emit(token.comments)
                result += this.emit(token.modifiers)
                result += token.name
                result += this.wrap('{', '}', ',\n', token.values)
                break
            case 'import-symbol':
                result = token.symbol
                if (token.as)
                    result += ` as ${token.as}`
                break
            case 'import':
                result += 'import '
                if (token.star)
                    result += `* as ${token.star}`
                if (token.symbols)
                    result += this.wrap('{', '}', ', ', token.symbols)
                result += ` from ${token.from}`
                break
            default:
                throw this.invalidTokenErr(token)
        }

        return result
    }
}

export type TsKeyword = typeof TsKeywords[number];

const TsKeywords = [
    "abstract",
    "any",
    "as",
    "asserts",
    "async",
    "await",
    "boolean",
    "break",
    "case",
    "catch",
    "class",
    "const",
    "constructor",
    "continue",
    "debugger",
    "declare",
    "default",
    "delete",
    "do",
    "else",
    "enum",
    "export",
    "extends",
    "false",
    "finally",
    "for",
    "from",
    "function",
    "get",
    "if",
    "implements",
    "import",
    "in",
    "infer",
    "instanceof",
    "interface",
    "is",
    "keyof",
    "let",
    "module",
    "namespace",
    "never",
    "new",
    "null",
    "number",
    "object",
    "package",
    "private",
    "protected",
    "public",
    "readonly",
    "require",
    "global",
    "return",
    "set",
    "static",
    "string",
    "super",
    "switch",
    "symbol",
    "this",
    "throw",
    "true",
    "try",
    "type",
    "typeof",
    "undefined",
    "unique",
    "unknown",
    "var",
    "void",
    "while",
    "with",
    "yield"
] as const

export type TsPunctuation = typeof TsPunctuations[number];
const TsPunctuations = [
    "{",
    "}",
    "(",
    ")",
    "[",
    "]",
    ".",
    "...",
    ";",
    ",",
    "<",
    ">",
    "<=",
    ">=",
    "==",
    "!=",
    "===",
    "!==",
    "=",
    "+",
    "-",
    "*",
    "**",
    "/",
    "%",
    "++",
    "--",
    "<<",
    ">>",
    ">>>",
    "&",
    "|",
    "^",
    "!",
    "~",
    "&&",
    "||",
    "?",
    ":",
    "@",
    "?.",
    "??",
    "&=",
    "|=",
    "^=",
    "+=",
    "-=",
    "*=",
    "/=",
    "%=",
    "<<=",
    ">>=",
    ">>>=",
    "**=",
    "=>",
    "`",
    "${"
] as const;


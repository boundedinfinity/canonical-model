import { StringBuffer } from "./utils.ts";

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

type TsModifierValue = 'const' | 'export' | 'public' | 'private' | 'static' | 'default'
export type TsModifers = {
    kind: 'modifiers'
    types: TsModifierValue[]
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
    }),
    modifiers: (vs: TsModifierValue[]): TsModifers => {
        return {
            kind: 'modifiers',
            types: vs
        }
    }
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

    private wrap(sb: StringBuffer, before: string, after: string, sep: TsSep, tokens?: TsToken[]) {
        sb.a(before)

        if (tokens) {
            for (let i = 0; i < tokens.length; i++) {
                const token = tokens[i]
                this.emit(sb, token)

                if (i < tokens.length - 1) {
                    switch (sep) {
                        case ', ':
                            sb.a(sep)
                            break
                        case ',\n':
                            sb.a(',')
                            sb.nl()
                            break
                        case '\n':
                            sb.nl()
                            break
                        case '\n\n':
                            sb.nl()
                            sb.nl()
                            break
                    }

                }
            }
        }

        sb.a(after)
    }

    private invalidTokenErr(token: TsToken) {
        new Error(`invalid token: ${JSON.stringify(token)}`)
    }

    emit(sb: StringBuffer, token?: TsToken) {
        if (!token)
            return

        switch (token.kind) {
            case 'type':
                this.emit(sb, token.type)
                break
            case 'primitive-type':
                sb.a(token.type)
                break
            case 'object-type':
                this.wrap(sb, '{', '}', token.sep, token.properties)
                break
            case 'raw':
                sb.a(token.value)
                break
            case 'newline':
                sb.a('\n')
                break
            case 'primitive':
                switch (typeof token.value) {
                    case 'string':
                        sb.a(`'${token.value}'`)
                        break
                    case 'number':
                    case 'boolean':
                    case 'symbol':
                        sb.a(String(token.value))
                        break
                    default:
                        throw this.invalidTokenErr(token)
                }
                break
            case 'modifiers':
                for (const type of token.types)
                    sb.a(`${type} `)
                break
            case 'property':
                this.emit(sb, token.comments)
                if (token.variadic)
                    sb.a('...')
                sb.a(token.name)
                if (token.optional)
                    sb.a('?')
                sb.a(`: `)
                this.emit(sb, token.type)
                if (token.array || token.variadic)
                    sb.a('[]')
                break
            case 'arguments':
                this.wrap(sb, '(', ')', token.sep, token.items)
                break
            case 'parameter-object':
                switch (token.sep) {
                    case ',\n':
                    case '\n':
                    case '\n\n':
                        sb.i(() => {
                            this.wrap(sb, '{', '}', token.sep, token.items)
                        })
                        break
                    default:
                        this.wrap(sb, '{', '}', token.sep, token.items)
                }
                break
            case 'literal':
                this.emit(sb, token.value)
                break
            case 'literal-object':
                {
                    const fn = () => {
                        for (let i = 0; i < token.items.length; i++) {
                            const item = token.items[i]
                            sb.a(`${item.name}: `)
                            this.emit(sb, item.value)

                            if (i < token.items.length - 1) {
                                switch (token.sep) {
                                    case ', ':
                                        sb.a(token.sep)
                                        break
                                    case ',\n':
                                        sb.a(',')
                                        sb.nl()
                                        break
                                    case '\n':
                                        sb.nl()
                                        break
                                    case '\n\n':
                                        sb.nl()
                                        sb.nl()
                                        break
                                }
                            }
                        }
                    }

                    sb.a(`{`)
                    switch (token.sep) {
                        case ',\n':
                        case '\n':
                        case '\n\n':
                            sb.i(() => {
                                fn()
                            })
                            break
                        default:
                            fn()
                    }

                    sb.a(`}`)
                }
                break
            case 'literal-array':
                this.wrap(sb, '[', ']', ', ', token.items)
                break
            case 'method-decl':
                this.emit(sb, token.comments)
                this.emit(sb, token.modifiers)

                sb.a(token.name)

                if (token.args)
                    this.emit(sb, token.args)
                else
                    this.emit(sb, { kind: 'arguments', sep: ', ' })

                if (token.return)
                    sb.a(`: ${token.return}`)

                {
                    const body: TsRaw[] = token.body?.map(line => {
                        return { kind: 'raw', value: line }
                    }) ?? []
                    this.wrap(sb, '{', '}', '\n', body)
                }
                break
            case 'function-decl':
                sb.a(`function `)
                this.emit(sb, { ...token, kind: 'method-decl' })
                break
            case 'class-decl':
                {
                    this.emit(sb, token.comments)
                    this.emit(sb, token.modifiers)
                    sb.a(`class ${token.name} `)

                    const properties: TsToken[] = token.properties
                        ? [...token.properties, this.Builder.newline()]
                        : []

                    const methods: TsToken[] = token.methods
                        ? token.methods.map(m => [m, this.Builder.newline()]).flat()
                        : []

                    const items = [...properties, ...methods]

                    if (items.length > 0)
                        sb.nl()

                    this.wrap(sb, '{', '}', '\n', items)
                }
                break
            case 'comment':
                if (token.lines) {
                    for (const line of token.lines)
                        sb.l(`// ${line}`)

                    if (!token.noNlewline)
                        sb.nl()
                }
                break
            case 'enum-value':
                this.emit(sb, token.comments)
                sb.a(`${token.name} = `)
                this.emit(sb, token.value)
                break
            case 'enum-decl':
                this.emit(sb, token.comments)
                this.emit(sb, token.modifiers)
                sb.a(token.name)
                this.wrap(sb, '{', '}', ',\n', token.values)
                break
            case 'import-symbol':
                sb.a(token.symbol)
                if (token.as)
                    sb.a(` as ${token.as}`)
                break
            case 'import':
                sb.a('import ')
                if (token.star)
                    sb.a(`* as ${token.star}`)
                if (token.symbols)
                    this.wrap(sb, '{', '}', ', ', token.symbols)
                sb.a(` from '${token.from}'`)
                sb.nl(2)
                break
            default:
                throw this.invalidTokenErr(token)
        }
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


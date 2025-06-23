import _ from 'lodash'

export function ce<T>(s: string, type: T): Error {
    return new Error(`${s}: ${JSON.stringify(type)}`)
}

export type Prettify<T> = {
    [K in keyof T]: T[K];
    // deno-lint-ignore ban-types
} & {};


export function fnwrap(pre: () => void, post: () => void, run: () => void) {
    try {
        pre()
        run()
    } finally {
        post()
    }
}

export type StringBufferOptions = {
    indent: number
    tabSize: number
}

export class StringBuffer {
    lines: string[] = []
    options: StringBufferOptions

    constructor(options?: Partial<StringBufferOptions>) {
        this.options = {
            indent: 0,
            tabSize: 4,
            ...options
        }
    }

    i(run: () => void) {
        fnwrap(
            () => this.indent(),
            () => this.dedent(),
            run
        )
    }

    merge(sb: StringBuffer | string) {
        if (typeof sb == 'string')
            this.lines.push(...sb.split('\n'))
        else
            this.lines.push(...sb.lines)
    }

    indent(n?: number) {
        const r = n ? n : 1
        this.options.indent += r
    }

    dedent(n?: number) {
        const r = n ? n : 1
        this.options.indent -= r
        if (this.options.indent < 0)
            this.options.indent = 0
    }

    redent(n?: number) {
        this.indent(n)
        for (let i = 0; i < this.lines.length; i++) {
            this.lines[i] = this.prefix(this.lines[i])
        }
        this.dedent(n)
    }

    private prefix(s: string): string {
        const size = this.options.indent * this.options.tabSize
        const prefix = " ".repeat(size)
        let line = s.trimStart()
        line = prefix + line
        return line
    }

    a(s: string) {
        if (this.lines.length == 0)
            this.nl()

        const index = this.lines.length - 1
        let line = this.lines[index]

        if (line == '')
            line = this.prefix(line)

        line += s
        this.lines[index] = line
    }

    l(s: string) {
        const current = this.prefix(s)
        this.lines.push(current)
    }

    j(s: string[], sep: string) {
        const last = s.length - 1
        for (let i = 0; i < s.length; i++) {
            if (i != last)
                this.a(s[i] + sep)
            else
                this.a(s[i])
        }
    }

    nl(n?: number) {
        const c = n ?? 1

        for (let i = 0; i < c; i++)
            this.lines.push('')
    }

    toString(): string {
        return this.lines.join('\n')
    }
}

export function isKey<T extends object>(name: string, token: T): boolean {
    for (const [key, val] of Object.entries(token)) {
        if (key == name && typeof val === 'boolean')
            return val
    }

    return false
}

function ensureSingleChar(c: string) {
    if (c.length != 1)
        throw new Error(`${c} must be single character`)
}

const numbers = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']
const symbols = ['~', '`', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '/', '\\', "'", '"', ',', '.', '?', '<', '>', '|', '[', ']']
const spaces = [' ', '\t']

export const stringUtils = {
    numbers, spaces, symbols,
    pipe(s: string, ...fns: ((s: string) => string)[]) {
        let result = s

        for (const fn of fns) {
            result = fn(result)
        }

        return result
    },
    splitThen(sep: string, fn: (s: string) => string): (s: string) => string {
        return function (s: string): string {
            const results: string[] = []

            for (const comp of s.split(sep)) {
                results.push(fn(comp))
            }

            return results.join(sep)
        }
    },
    isLetter: function (c: string): boolean {
        ensureSingleChar(c)
        return c.toLocaleLowerCase() != c.toLocaleUpperCase()
    },
    isNumber: function (c: string): boolean {
        ensureSingleChar(c)
        return numbers.includes(c)
    },
    isSymbol: function (c: string): boolean {
        ensureSingleChar(c)
        return symbols.includes(c)
    },
    isSpace: function (c: string): boolean {
        ensureSingleChar(c)
        return spaces.includes(c)
    },
    onlySpace(s: string): boolean {
        for (const c of s) {
            if (!this.isSpace(c))
                return false
        }
        return true
    },
    isUpperCase: function (c: string): boolean {
        return stringUtils.isLetter(c) && c == c.toLocaleUpperCase()
    },
    isLowerCase: function (c: string): boolean {
        return stringUtils.isLetter(c) && c == c.toLocaleLowerCase()
    },
    uppercase(s: string): string {
        return s.toLocaleUpperCase()
    },
    uppperCaseFirst(s: string): string {
        return s.slice(0, 1).toLocaleUpperCase() + s.slice(1)
    },
    lowerCase(s: string): string {
        return s.toLocaleLowerCase()
    },
    lowerCaseFirst(s: string): string {
        return s.slice(0, 1).toLocaleLowerCase() + s.slice(1)
    },
    replaceWith(s: string, from: string[], to: string): string {
        const results: string[] = []

        for (const c of s) {
            if (from.includes(c)) {
                results.push(to)
                continue
            }

            results.push(c)
        }

        return results.join('')
    },
    singleQuote: function (s: string): string {
        let o = s
        o = o.replaceAll(`'`, `\\'`)
        o = `'` + o + `'`
        return o
    },
    doubleQuote: function (s: string): string {
        let o = s
        o = o.replaceAll(`"`, `\"`)
        o = `"` + o + `"`
        return o
    },
    replaceWithFn: function (from: string[], to: string): (s: string) => string {
        return function (s: string): string {
            return stringUtils.replaceWith(s, from, to)
        }
    },
    compactSpaces: function (s: string): string {
        let result = s
        result = result.trim()
        result = result.replace(/ +/, ' ')
        return result
    },
    unique: function (strs: string[]): string[] {
        let result = [...strs]
        result = [...new Set(result)]
        return result
    },
    test: {
        normal: function (text: string): string {
            return text
                .trim()
                .split('\n')
                .map(line => line.trim())
                .filter(line => line !== '')
                .join('\n')
                .trim()
        },
        save: function (file: string, text: string) {
            const encoder = new TextEncoder();
            Deno.writeFileSync(file, encoder.encode(text))
        }
    }
}


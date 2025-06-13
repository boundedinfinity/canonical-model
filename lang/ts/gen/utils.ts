import _ from 'lodash'

export type Prettify<T> = {
    [K in keyof T]: T[K];
    // deno-lint-ignore ban-types
} & {};


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
    }
}

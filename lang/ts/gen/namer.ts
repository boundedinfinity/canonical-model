import _ from 'lodash'
import { Caser, type CaseType } from './caser.ts'

export class TsNamer {
    private caser = new Caser()

    options: {
        safeChar: string,
        firstChar: string,
        fromCase: CaseType
    } = {
            safeChar: '_',
            firstChar: '_',
            fromCase: 'phrase'
        }

    constructor(options?: Partial<typeof this.options>) {
        _.merge(this.options, options)
    }

    tsVar(s: string): string {
        return this.caser.phrase2Camel(s)
    }

    tsClass(s: string): string {
        return this.caser.phrase2Pascal(s)
    }

    sqlTable(s: string): string {
        return this.caser.phrase2SnakeLower(s)
    }

    sqlColumn(s: string): string {
        return this.caser.phrase2SnakeLower(s)
    }
}
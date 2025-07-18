import _ from 'lodash'

import { Label } from './label'
import { NamePrefix } from './name-prefix'
import { NameSuffix } from './name-suffix'

export class Name {
    kind = 'label'
    id: string
    prefix?: NamePrefix[]
    first?: string[]
    middle?: string[]
    last?: string[]
    suffix?: NameSuffix[]
    labels?: Label[]

    constructor(args: {
        id: string,
        prefix?: NamePrefix[],
        first?: string[],
        middle?: string[],
        last?: string[],
        suffix?: NameSuffix[],
        labels?: Label[]
    }) {
        this.id = args.id
        this.prefix = args.prefix
        this.first = args.first
        this.middle = args.middle
        this.last = args.last
        this.suffix = args.suffix
        this.labels = args.labels
    }

    validate(options?: { prefix?: string }): Error[] {
        const prefix = options?.prefix ? `${options.prefix}.` : '' + `label.`
        const errors: Error[] = []
        const oneOf = [this.middle, this.middle, this.middle]

        if (oneOf.filter(prop => prop !== undefined).length < 1) {
            errors.push(new Error(`must include one of ${prefix}first, ${prefix}middle, ${prefix}last`))
        }

        if (this.id === undefined || this.id === null) {
            errors.push(new Error(`${prefix}id is required`))
        } else {
            if (this.id.length < 36) {
                errors.push(new Error(`${prefix}id is less than 36 characters`))
            }

            if (this.id.length > 36) {
                errors.push(new Error(`${prefix}id is greater than 36 characters`))
            }
        }

        if (this.prefix && this.prefix.length > 0) {
            for (let i = 0; i < this.prefix.length; i++) {
                const label = this.prefix[i]
                errors.push(...label.validate({ prefix: 'name.' }))
            }
        }

        if (this.first && this.first.length > 1) {
            for (let i = 0; i < this.first.length; i++) {
                const first = this.first[i]

                if (first.length < 2) {
                    errors.push(new Error(`${prefix}first[${i}] is less than 2 characters`))
                }

                if (first.length < 25) {
                    errors.push(new Error(`${prefix}first[${i}] is greater than 25 characters`))
                }
            }
        }

        if (this.middle && this.middle.length > 1) {
            for (let i = 0; i < this.middle.length; i++) {
                const first = this.middle[i]

                if (first.length < 2) {
                    errors.push(new Error(`${prefix}middle[${i}] is less than 2 characters`))
                }

                if (first.length < 25) {
                    errors.push(new Error(`${prefix}middle[${i}] is greater than 25 characters`))
                }
            }
        }

        if (this.last && this.last.length > 1) {
            for (let i = 0; i < this.last.length; i++) {
                const first = this.last[i]

                if (first.length < 2) {
                    errors.push(new Error(`${prefix}last[${i}] is less than 2 characters`))
                }

                if (first.length < 25) {
                    errors.push(new Error(`${prefix}last[${i}] is greater than 25 characters`))
                }
            }
        }

        if (this.suffix && this.suffix.length > 0) {
            for (let i = 0; i < this.suffix.length; i++) {
                const label = this.suffix[i]
                errors.push(...label.validate({ prefix: 'name.' }))
            }
        }

        if (this.labels && this.labels.length > 0) {
            for (let i = 0; i < this.labels.length; i++) {
                const label = this.labels[i]
                errors.push(...label.validate({ prefix: 'name.' }))
            }
        }

        return errors
    }
}

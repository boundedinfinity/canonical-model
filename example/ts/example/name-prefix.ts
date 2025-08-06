import { Label } from './label'

export class NamePrefix {
    kind = 'name-prefix'
    id: string
    name: string
    abbr?: string[]
    description?: string
    labels?: Label[]

    constructor(args: {
        id: string,
        name: string,
        abbr?: string[]
        description?: string
        labels?: Label[]
    }) {
        this.id = args.id
        this.name = args.name
        this.abbr = args.abbr
        this.labels = args.labels
    }

    validate(options?: { prefix?: string }) {
        const prefix = options?.prefix ? `${options.prefix}.` : '' + 'namePrefix.'
        const errors: Error[] = []

        if (!this.id) {
            errors.push(new Error(`${prefix}id is required`))
        }

        if (!this.name) {
            errors.push(new Error(`${prefix}name is required`))
        } else {
            if (this.name.length < 2) {
                errors.push(new Error(`${prefix}name is less than 2 characters`))
            }

            if (this.name.length > 50) {
                errors.push(new Error(`${prefix}name is more than 50 characters`))
            }
        }

        if (this.abbr && this.abbr.length > 1) {
            for (let i = 0; i < this.abbr.length; i++) {
                const abbr = this.abbr[i]

                if (abbr.length < 2) {
                    errors.push(new Error(`${prefix}abbr[${i}] is less than 2 characters`))
                }

                if (abbr.length < 25) {
                    errors.push(new Error(`${prefix}abbr[${i}] is greater than 25 characters`))
                }
            }
        }

        if (this.description) {
            if (this.description.length < 2) {
                errors.push(new Error(`${prefix}description is less than 2 characters`))
            }

            if (this.description.length > 200) {
                errors.push(new Error(`${prefix}description is more than 200 characters`))
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

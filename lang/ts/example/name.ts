import _ from 'lodash'
import { Label, LabelDtoFull, LabelDtoId } from './label'
import { NamePrefix, NamePrefixDtoId, NamePrefixJson } from './name-prefix'
import { NameSuffix, NameSuffixDtoId, NameSuffixJson } from './name-suffix'

interface NameDtoId {
    kind: 'name-id'
    id: string
}

interface NameDtoRef {
    kind: 'name-ref'
    id: string
    prefix?: NamePrefixDtoId[]
    first?: string[]
    middle?: string[]
    last?: string[]
    suffix?: NameSuffixDtoId[]
    labels?: LabelDtoId[]
}

interface NameDtoFull {
    kind: 'name-full'
    id: string
    prefix?: string[]
    first?: string[]
    middle?: string[]
    last?: string[]
    suffix?: string[]
    labels?: LabelDtoFull[]
}

export class LabelJson {
    config: {
        kind: 'id' | 'ref' | 'full'
    } = {
            kind: 'full'
        }

    constructor(config?: Partial<typeof this.config>) {
        Object.assign(this.config, config)
    }

    serialize(obj: Name, config?: Partial<typeof this.config>): NameDtoId | NameDtoRef | NameDtoFull {
        const rconfig: typeof this.config = _.merge(this.config, config)
        const prefix = new NamePrefixJson()
        const suffix = new NameSuffixJson()
        const label = new LabelJson()

        switch (rconfig.kind) {
            case 'id':
                return { kind: 'name-id', id: obj.id } as NameDtoId
            case 'ref':
                return {
                    kind: 'name-ref',
                    id: obj.id,
                    prefix: obj.prefix?.map(p => prefix.serialize(p, { kind: 'id' })),
                    first: obj.first?.map(n => n),
                    middle: obj.middle?.map(n => n),
                    last: obj.last?.map(n => n),
                    suffix: obj.suffix?.map(s => suffix.serialize(s, { kind: 'id' })),
                    labels: obj.labels?.map(l => label.serialize(l, { kind: 'id' })),
                } as NameDtoRef
            default:
                return {
                    kind: 'name-full',
                    id: obj.id,
                    prefix: obj.prefix?.map(p => prefix.serialize(p)),
                    first: obj.first?.map(n => n),
                    middle: obj.middle?.map(n => n),
                    last: obj.last?.map(n => n),
                    suffix: obj.suffix?.map(s => suffix.serialize(s)),
                    labels: obj.labels?.map(l => label.serialize(l)),
                } as NameDtoFull
        }
    }
}

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

import _ from 'lodash'

export interface LabelDtoId {
    kind: 'label-id'
    id: string
}

export interface LabelDtoFull {
    kind: 'label-full'
    id: string
    name: string
    description?: string | null
}

export class LabelJson {
    config: {
        kind: 'id' | 'full'
    } = {
            kind: 'full'
        }

    constructor(config?: Partial<typeof this.config>) {
        Object.assign(this.config, config)
    }

    serialize(obj: Label, config?: Partial<typeof this.config>): LabelDtoId | LabelDtoFull {
        const rconfig: typeof this.config = _.merge(this.config, config)

        switch (rconfig.kind) {
            case 'id':
                return { kind: 'label-id', id: obj.id } as LabelDtoId
            default:
                return {
                    kind: 'label-full',
                    id: obj.id,
                    name: obj.name,
                    description: obj.description
                } as LabelDtoFull
        }
    }
}

export class Label {
    kind = 'label'
    id: string
    name: string
    description?: string | null

    constructor(args: Omit<LabelDtoFull, 'kind'>) {
        this.id = args.id
        this.name = args.name
        this.description = args.description
    }

    validate(options?: { prefix?: string }): Error[] {
        const prefix = options?.prefix ? `${options.prefix}.` : '' + `label.`
        const errors: Error[] = []

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

        if (this.name === undefined || this.name === null) {
            errors.push(new Error(`${prefix}name is required`))
        } else {
            if (this.name.length < 2) {
                errors.push(new Error(`${prefix}name is less than 2 characters`))
            }

            if (this.name.length > 50) {
                errors.push(new Error(`${prefix}name is greater than 50 characters`))
            }
        }

        if (this.description) {
            if (this.description.length < 2) {
                errors.push(new Error(`${prefix}description is less than 2 characters`))
            }

            if (this.description.length > 50) {
                errors.push(new Error(`${prefix}description is greater than 500 characters`))
            }
        }

        return errors
    }
}



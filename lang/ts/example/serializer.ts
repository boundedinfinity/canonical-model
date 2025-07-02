import _ from 'lodash'
import { Label } from './label'
import { Name } from './name'
import { NamePrefix } from './name-prefix'
import { NameSuffix } from './name-suffix'

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

export class Serializer {

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


export interface NamePrefixDtoId {
    kind: 'name-prefix-id'
    id: string
}

export interface NamePrefixDtoRef {
    kind: 'name-prefix-ref'
    id: string
    name: string
    abbr?: string[]
    description?: string | null
    labels: LabelDtoId[]
}

export interface NamePrefixDtoFull {
    kind: 'name-prefix-full'
    id: string
    name: string
    abbr?: string[]
    description?: string | null
    labels: LabelDtoFull[]
}

export class NamePrefixJson {
    config: {
        kind: 'id' | 'ref' | 'full'
    } = {
            kind: 'full'
        }

    constructor(config?: Partial<typeof this.config>) {
        Object.assign(this.config, config)
    }

    serialize(obj: NamePrefix, config?: Partial<typeof this.config>): NamePrefixDtoId | NamePrefixDtoRef | NamePrefixDtoFull {
        const rconfig: typeof this.config = _.merge(this.config, config)
        const labelJson = new LabelJson()

        switch (rconfig.kind) {
            case 'id':
                return { kind: 'name-prefix-id', id: obj.id } as NamePrefixDtoId
            case 'ref':

                return {
                    kind: 'name-prefix-ref',
                    id: obj.id,
                    name: obj.name,
                    abbr: obj.abbr?.map(a => a),
                    description: obj.description,
                    labels: obj.labels?.map(label => labelJson.serialize(label, { kind: 'id' })),
                } as NamePrefixDtoRef
            default:
                return {
                    kind: 'name-prefix-full',
                    id: obj.id,
                    name: obj.name,
                    abbr: obj.abbr?.map(a => a),
                    description: obj.description,
                    labels: obj.labels?.map(label => labelJson.serialize(label)),
                } as NamePrefixDtoFull
        }
    }
}


export interface NameSuffixDtoId {
    kind: 'name-suffix-id'
    id: string
}

export interface NameSuffixDtoRef {
    kind: 'name-suffix-ref'
    id: string
    name: string
    abbr?: string[]
    description?: string | null
    labels: LabelDtoId[]
}

export interface NameSuffixDtoFull {
    kind: 'name-suffix-full'
    id: string
    name: string
    abbr?: string[]
    description?: string | null
    labels: Label[]
}

export class NameSuffixJson {
    config: {
        kind: 'id' | 'ref' | 'full'
    } = {
            kind: 'full'
        }

    constructor(config?: Partial<typeof this.config>) {
        Object.assign(this.config, config)
    }

    serialize(obj: NameSuffix, config?: Partial<typeof this.config>): NameSuffixDtoId | NameSuffixDtoRef | NameSuffixDtoFull {
        const rconfig: typeof this.config = _.merge(this.config, config)
        const label = new LabelJson()

        switch (rconfig.kind) {
            case 'id':
                return { kind: 'name-suffix-id', id: obj.id } as NameSuffixDtoId
            case 'ref':
                return {
                    kind: 'name-suffix-ref',
                    id: obj.id,
                    name: obj.name,
                    abbr: obj.abbr?.map(a => a),
                    description: obj.description,
                    labels: obj.labels?.map(l => label.serialize(l, { kind: 'id' })),
                } as NameSuffixDtoRef
            default:
                return {
                    kind: 'name-suffix-full',
                    id: obj.id,
                    name: obj.name,
                    abbr: obj.abbr?.map(a => a),
                    description: obj.description,
                    labels: obj.labels?.map(l => label.serialize(l)),
                } as NameSuffixDtoFull
        }
    }
}


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

export class NameJson {
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

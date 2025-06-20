type LabelJsonOptions = {
    kind: 'full' | 'id',
    indent: number
}

type LabelSelectArgs = {
    id?: string,
    columns?: ('id' | 'name' | 'description')[]
    name?: {
        equals?: string,
        contains?: string
    },
    description?: {
        contains?: string
    },
    term: {
        equals?: string
        contains?: string
    }
}

class Label {
    id: string
    name: string
    description?: string

    constructor(args: {
        id: string,
        name: string,
        description?: string
    }) {
        this.id = args.id
        this.name = args.name
        this.description = args.description
    }

    toJson(options?: Partial<LabelJsonOptions>): string {
        const roptions: LabelJsonOptions = {
            ...{ kind: 'full', indent: 4 },
            ...options
        }

        const obj: object = { kind: 'label' }

        switch (roptions.kind) {
            case 'id':
                Object.assign(obj, { id: this.id })
                break
            case 'full':
                Object.assign(obj, {
                    id: this.id,
                    name: this.name,
                    description: this.description
                })
                break
        }

        return JSON.stringify(obj, null, roptions.indent)
    }

    static fromJson(json: string): Label {
        const obj = JSON.parse(json)
        return new Label(obj)
    }

    static canJson(json: string): boolean {
        const obj: { kind: string } = JSON.parse(json)
        return obj.kind === 'label'
    }

    sqlCreate(): string {
        return `
        CREATE TABLE label (
            id: TEXT,
            name: TEXT NOT NULL,
            description: TEXT,
            PRIMARY KEY (id),
        );`
    }

    sqlInsert(): string {
        const columns = ['id', 'name']
        if (this.description) columns.push('description')

        const values = [this.id, this.name]
        if (this.description) values.push(this.description)


        return `
            INSERT INTO label (
                ${columns.join(', ')}
            ) VALUES (
                ${values.map(v => `'${v}'`).join(', ')}
            );
        `
    }

    sqlSelect(options?: LabelSelectArgs) {
        const columns: string[] = []
        const wheres: string[] = []

        options?.columns?.forEach(c => columns.push(c)) ?? ['*']

        if (options?.id)
            wheres.push(`id = ${this.id}`)

        let sql = `
            SELECT ${columns.join(', ')}
            FROM labels
            WHERE
        `
        return ``
    }

    validate(): Error[] {
        const errors: Error[] = []

        if (this.id === undefined || this.id === null) {
            errors.push(new Error('id is required'))
        } else {
            if (this.id.length < 36) {
                errors.push(new Error('id is less than 36 characters'))
            }

            if (this.id.length > 36) {
                errors.push(new Error('id is greater than 36 characters'))
            }
        }

        if (this.name === undefined || this.name === null) {
            errors.push(new Error('name is required'))
        } else {
            if (this.name.length < 2) {
                errors.push(new Error('name is less than 2 characters'))
            }

            if (this.name.length > 50) {
                errors.push(new Error('name is greater than 50 characters'))
            }
        }

        if (this.description) {
            if (this.description.length < 2) {
                errors.push(new Error('description is less than 2 characters'))
            }

            if (this.description.length > 50) {
                errors.push(new Error('description is greater than 500 characters'))
            }
        }


        return errors
    }
}


interface NamePrefixData {

}

export class NamePrefix {
    id: string
    name: string
    abbr?: string[]
    description?: string

    constructor(data: {
        id: string
        name: string
        abbr?: string[]
        description?: string
    }) {
        this.id = data.id
        this.name = data.name
        this.abbr = data.abbr
        this.description = data.description
    }

    validate() {
        if (this.id.length < 36)
            throw new Error(`id is less than 36`)

        if (this.id.length > 36)
            throw new Error(`id is greater than 36`)

        if (!this.id.includes('-'))
            throw new Error(`id does not contain '-'`)
    }
}

export class NamePrefixTextMarshaller {
    marshal(value: NamePrefix): string {
        if (value.data.abbr && value.data.abbr.length > 0)
            return value.data.abbr[0]

        return value.data.name
    }
}

interface NamePrefixMarshalOptions {
    indent: number
    marshalType: 'full' | 'id'
}

export class NamePrefixObjectMarshaller {
    options: NamePrefixMarshalOptions = {
        marshalType: 'full',
        indent: 4
    }

    constructor(options: Partial<NamePrefixMarshalOptions>) {
        this.options = { ...this.options, ...options }
    }

    marshal(item: NamePrefix,): object {
        let obj: object
        switch (this.options.marshalType) {
            case 'full':
                obj = item.data
                break
            case 'id':
                obj = { 'id': item.data.id }
                break
        }

        return obj
    }

    unmarshal(item: NamePrefixData,): NamePrefix {
        return new NamePrefix(item)
    }
}

export class NamePrefixJsonMarshaller {
    options: NamePrefixMarshalOptions = {
        marshalType: 'full',
        indent: 4
    }

    constructor(options: Partial<NamePrefixMarshalOptions>) {
        this.options = { ...this.options, ...options }
    }

    marshal(item: NamePrefix,): string {
        const obj = new NamePrefixObjectMarshaller(this.options).marshal(item)
        const json = JSON.stringify(obj, null, this.options.indent)
        return json
    }

    unmarshal(item: string): NamePrefix {
        const obj: NamePrefixData = JSON.parse(item)
        return new NamePrefix(obj)
    }
}

interface NamePrefixSqlMarshallerOptions {
    indent: number
    marshalType: 'full' | 'id'
}

export class NamePrefixSqlMarshaller {
    options: NamePrefixSqlMarshallerOptions = {
        marshalType: 'full',
        indent: 4
    }

    constructor(options: Partial<NamePrefixSqlMarshallerOptions>) {
        this.options = { ...this.options, ...options }
    }

    table(): string[] {
        return [`
            CREATE TABLE name_prefix (
                id TEXT,
                name TEXT,
                description TEXT
            )
        `, `
            CREATE TABLE name_prefix__abbr (
                name_prefix_id TEXT,
                index INTEGER,
                abbr TEXT
            )
        `]
    }

    insert(input: NamePrefix): string[] {
        const abbrs = input.data.abbr?.map((abbr, index) => {
            return `
                INSERT INTO name_prefix__abbr (
                    name_prefix_id,
                    index,
                    abbr
                ) VALUES (
                    ${input.data.id},
                    ${index},
                    ${abbr}
                )
            `
        }) || []

        return [`
            INSERT INTO name_prefix (
                id, 
                name, 
                description
            ) VALUES (
                ${input.data.id}, 
                ${input.data.name}, 
                ${input.data.description}
            )
        `, ...abbrs]
    }

    update(input: NamePrefix): string[] {
        const abbrs = input.data.abbr?.map((abbr, index) => {
            return `
                INSERT INTO name_prefix__abbr (
                    name_prefix_id,
                    index,
                    abbr
                ) VALUES (
                    '${input.data.id}',
                    ${index},
                    '${abbr}'
                );
            `
        }) || []


        return [`
                DELETE FROM name_prefix__abbr
                WHERE id = '${input.data.id}';
            `, `
            INSERT INTO name_prefix (
                id, 
                name, 
                description
            ) VALUES (
                '${input.data.id}', 
                '${input.data.name}', 
                '${input.data.description}'
            );
        `, ...abbrs]
    }

    get(id: string): string[] {
        const queries = [
            `
            SELECT id, name, description
                FROM name_prefix
                WHERE id = '${id}';
            `,
            `
           SELECT name_prefix_id AS id, abbr
                FROM name_prefix__abbr
                WHERE id = '${id}'
                ORDER BY index ASC;
            `,
        ]


        return queries
    }

    find(options?: NamePrefixQueryOptions): string[] {
        const queries: string[] = []
        const prefixClauses: string[] = []
        const abbrClauses: string[] = []


        sqlStringClauses(prefixClauses, 'name', options?.name, options?.modifiers)
        sqlStringClauses(prefixClauses, 'name', options?.term, options?.modifiers)

        sqlStringClauses(abbrClauses, 'abbr', options?.abbr, options?.modifiers)
        sqlStringClauses(abbrClauses, 'abbr', options?.term, options?.modifiers)


        if (prefixClauses.length > 0) {
            queries.push(`
                SELECT id, name, description
                FROM name_prefix
                WHERE ${prefixClauses.join(' OR ')};
            `)
        }

        if (abbrClauses.length > 0) {
            queries.push(`
                SELECT name_prefix_id AS id, abbr
                FROM name_prefix__abbr
                WHERE ${abbrClauses.join(' OR ')}
                ORDER BY index ASC;
            `)
        }

        return queries
    }
}

function sqlStringClauses(
    clauses: string[] = [],
    field: string,
    query?: SqlStringQuery,
    modifiers?: SqlStringQueryModifer,
) {


    query?.equals?.forEach(value => {
        sqlStringClause(clauses, {
            field: field,
            op: 'equals',
            modifiers,
            value,
        })
    })

    query?.not?.forEach(value => {
        sqlStringClause(clauses, {
            field: field,
            op: 'not',
            modifiers,
            value,
        })
    })

    query?.endsWith?.forEach(value => {
        sqlStringClause(clauses, {
            field: field,
            op: 'endsWith',
            modifiers,
            value,
        })
    })

    query?.startsWith?.forEach(value => {
        sqlStringClause(clauses, {
            field: field,
            op: 'startsWith',
            modifiers,
            value,
        })
    })

    query?.includes?.forEach(value => {
        sqlStringClause(clauses, {
            field: field,
            op: 'includes',
            modifiers,
            value,
        })
    })

    query?.excludes?.forEach(value => {
        sqlStringClause(clauses, {
            field: field,
            op: 'excludes',
            modifiers,
            value,
        })
    })
}


function sqlStringClause(clauses: string[] = [], inputs: {
    field: string,
    op: 'equals' | 'not' | 'startsWith' | 'endsWith' | 'includes' | 'excludes',
    modifiers?: SqlStringQueryModifer,
    value: string,
}) {
    const sqlField = inputs.modifiers?.caseSensitive
        ? inputs.field
        : `UPPER(${inputs.field})`
    const sqlValue = inputs.modifiers?.caseSensitive
        ? inputs.value
        : inputs.value.toUpperCase()
    let clause: string

    switch (inputs.op) {
        case 'equals':
            clause = `${sqlField} = '${sqlValue}'`
            break
        case 'not':
            clause = `${sqlField} != '${sqlValue}'`
            break;
        case 'startsWith':
            clause = `${sqlField} LIKE '${sqlValue}%'`
            break;
        case 'endsWith':
            clause = `${sqlField} LIKE '%${sqlValue}'`
            break;
        case 'includes':
            clause = `${sqlField} LIKE '%${sqlValue}%'`
            break;
        case 'excludes':
            clause = `${sqlField} NOT LIKE '%${sqlValue}%'`
            break
        default:
            throw new Error(`invalid operation ${inputs.op}`)
    }

    clauses.push(clause)
}

interface SqlStringQuery {
    equals?: string[]
    not?: string[]
    includes?: string[]
    excludes?: string[]
    startsWith?: string[]
    endsWith?: string[]
}

interface SqlStringQueryModifer {
    caseSensitive?: boolean

}

interface NamePrefixQueryOptions {
    modifiers?: SqlStringQueryModifer
    name?: SqlStringQuery
    abbr?: SqlStringQuery
    term?: SqlStringQuery
}
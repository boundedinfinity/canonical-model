interface NamePrefixData {
    id: string
    name: string
    abbr?: string[]
    description?: string
}

export class NamePrefix {
    data: NamePrefixData

    constructor(data: NamePrefixData) {
        this.data = data
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
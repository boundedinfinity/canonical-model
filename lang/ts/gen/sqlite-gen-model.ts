import { StringBuffer } from "./utils.ts";

export class SqliteGenerator {
    sb: StringBuffer

    constructor(sb: StringBuffer) {
        this.sb = sb
    }

    private value(item: SqlValue) {
        function escape(s: string): string {
            return s.replaceAll("'", "''")
        }

        switch (item.kind) {
            case 'integer':
            case 'numeric':
            case 'real':
                this.sb.a(`${item.value}`)
                break
            case 'text':
                this.sb.a(`'${escape(item.value)}'`)
                break
            case 'ts-interpolation-numeric':
                this.sb.a('${' + item.name + '}')
                break
            case 'ts-interpolation-text':
                this.sb.a("'${" + item.name + "}'")
                break
            case 'sql-parameter':
                if (item.index)
                    this.sb.a(`?${item.index}`)
                else if (item.name)
                    this.sb.a(`:${item.name}`)
                else
                    this.sb.a('?')
                break
            case 'boolean':
                this.sb.a(`${item.value}`)
                break
            case 'date':
                switch (item.format) {
                    case 'epoch':
                        this.value({ kind: 'integer', value: item.value.getTime() })
                        break
                    case 'ISO-8601':
                        this.sb.a("TODO: ISO-8601")
                        break
                    case 'RFC-3339':
                        this.sb.a("TODO: RFC-3339")
                        break
                    default:
                        throw new Error(`invalid sql date format ${JSON.stringify(item)}`)
                }
                break
            default:
                throw new Error(`invalid sql value ${JSON.stringify(item)}`)
        }
    }

    private refd(item: SqlRference): string { return `${item.table}.${item.column}` }

    private refp(item: SqlRference): string { return `${item.table}(${item.column})` }

    expression(item: SqlExpr) {
        const handleList = (es: SqlExpr[], sep: string) => {
            for (let i = 0; i < es.length; i++) {
                this.expression(es[i])

                if (i < es.length - 1)
                    this.sb.a(sep)
            }
        }

        switch (item.kind) {
            case 'equal':
                this.sb.a(`${item.column} = `)
                this.value(item.value)
                break
            case 'not-equal':
                this.sb.a(`${item.column} != `)
                this.value(item.value)
                break
            case 'is':
                this.sb.a(`${item.column} IS `)
                this.value(item.value)
                break
            case 'is-not':
                this.sb.a(`${item.column} IS NOT `)
                this.value(item.value)
                break
            case 'greater-than':
                this.sb.a(`${item.column} > `)
                this.value(item.value)
                break
            case 'greater-than-or-equal':
                this.sb.a(`${item.column} >= `)
                this.value(item.value)
                break
            case 'less-than':
                this.sb.a(`${item.column} > `)
                this.value(item.value)
                break
            case 'less-than-or-equal':
                this.sb.a(`${item.column} <= `)
                this.value(item.value)
                break
            case 'between':
                this.sb.a(`${item.column} BETWEEN `)
                this.value(item.lower)
                this.sb.a(` AND `)
                this.value(item.upper)
                break
            case 'not':
                this.sb.a(`NOT `)
                this.expression(item.expression)
                break
            case 'and':
                handleList(item.expressions, ' AND ')
                break
            case 'or':
                handleList(item.expressions, ' OR ')
                break
            case 'literal-value':
                this.value(item.value)
                break
            case 'in':
                {
                    this.sb.a(`${item.column} IN (`)

                    for (let i = 0; i < item.values.length; i++) {
                        const value = item.values[i]
                        this.value(value)

                        if (i < item.values.length - 1)
                            this.sb.a(',')
                    }

                    this.sb.a(')')
                }
                break
            case 'bind-parameter':
                // TODO
                break
            default:
                throw new Error(`invalid expression ${JSON.stringify(item)}`)
        }
    }

    updates(items: SqlUpdate[]) {
        if (items) {
            for (const item of items) {
                this.update(item)
                this.sb.nl()
            }
        }
    }

    update(item: SqlUpdate) {
        this.sb.l(`UPDATE ${item.table} SET`)
        this.sb.nl()

        {
            this.sb.indent()

            for (let i = 0; i < item.values.length; i++) {
                const value = item.values[i]
                this.sb.a(`${value.column} = `)
                this.value(value.value)

                if (i < item.values.length - 1) {
                    this.sb.a(', ')
                }

                this.sb.nl()
            }

            this.sb.dedent()
        }

        if (item.where) {
            this.sb.a(` WHERE `)
            this.expression(item.where)
        }

        this.sb.l(';')
    }

    selects(items: SqlSelect[]) {
        if (items) {
            for (const item of items) {
                this.select(item)
                this.sb.nl()
            }
        }
    }

    select(item: SqlSelect) {
        this.sb.a(`SELECT `)
        this.sb.j(item.columns, ', ')
        this.sb.nl()

        {
            this.sb.indent()
            this.sb.a(`FROM ${item.table}`)
            this.sb.dedent()
        }

        if (item.where) {
            this.sb.indent()
            this.sb.nl()
            this.sb.a(`WHERE `)
            this.expression(item.where)
            this.sb.dedent()
        }

        if (item.joins && item.joins.length > 0) {
            this.sb.indent()
            for (const join of item.joins) {
                let jtext = ''

                switch (join.kind) {
                    case 'join-on': {
                        if (join.operator)
                            jtext += ` ${join.operator}`

                        const tableRef: SqlRference = {
                            table: item.table,
                            column: join.column
                        }

                        jtext += `JOIN ${join.reference.table}`
                        jtext += `ON ${this.refd(tableRef)} = ${this.refd(join.reference)}`
                    }
                        break
                    case 'join-using':
                        if (join.operator)
                            jtext += ` ${join.operator}`

                        jtext += ` JOIN ${join.reference.table}`
                        jtext += ` USING (${join.reference.column})`
                        break
                    default:
                        throw new Error(`invalid join type ${JSON.stringify(join)}`)
                }

                this.sb.l(jtext)
            }
            this.sb.dedent()
        }

        if (item.orderBy) {
            this.sb.indent()
            this.sb.nl()
            this.sb.a(`ORDER BY ${item.orderBy.column}`)

            if (item.orderBy.term)
                this.sb.a(` ${item.orderBy.term}`)
            this.sb.dedent()
        }

        if (item.limit) {
            this.sb.indent()
            this.sb.nl()
            this.sb.a(`LIMIT `)
            this.expression(item.limit.expression)

            if (item.limit.offset) {
                this.sb.a(` OFFSET `)
                this.expression(item.limit.offset)
            }

            this.sb.dedent()
            this.sb.nl()
        }


        this.sb.a(';')
        this.sb.nl()
    }

    deletes(items: SqlDelete[]) {
        if (items) {
            for (const item of items) {
                this.delete(item)
            }
        }
    }

    delete(item: SqlDelete) {
        this.sb.a(`DELETE FROM ${item.table}`)

        if (item.where) {
            this.sb.a(` WHERE `)
            this.expression(item.where)
        }

        this.sb.a(';')
        this.sb.nl()
    }

    inserts(items: SqlInsert[]) {
        if (items) {
            for (const item of items) {
                this.insert(item)
                this.sb.nl()
            }
        }
    }

    insert(item: SqlInsert) {
        this.sb.l(`INSERT INTO ${item.table} (`)
        this.sb.nl()

        {
            this.sb.indent()

            for (let i = 0; i < item.values.length; i++) {
                const value = item.values[i]
                this.sb.a(value.column)

                if (i < item.values.length - 1) {
                    this.sb.a(', ')
                }
            }

            this.sb.dedent()
        }

        this.sb.l(`) VALUES (`)
        this.sb.nl()

        {
            this.sb.indent()

            for (let i = 0; i < item.values.length; i++) {
                const value = item.values[i]
                this.value(value.value)

                if (i < item.values.length - 1) {
                    this.sb.a(', ')
                }
            }

            this.sb.dedent()
        }

        this.sb.l(');')
    }

    indexes(items: SqlIndex[]) {
        if (items) {
            for (const item of items) {
                this.index(item)
                this.sb.nl()
            }
        }
    }

    index(item: SqlIndex) {
        let text = 'CREATE'

        if (item.unique)
            text += ' UNIQUE'

        text += ' INDEX'

        if (item.ifNotExists)
            text += ' IF NOT EXISTS'

        text += ` idx_`

        if (item.name && item.name !== '')
            text += `${item.name}`
        else
            text += `${item.table}__${item.column}`

        text += ` ON ${item.table} (${item.column}); `

        text += ';'

        this.sb.l(text)
    }

    tables(items: SqlTable[]) {
        if (items) {
            for (const item of items) {
                this.table(item)
                this.sb.nl()
            }
        }
    }

    table(item: SqlTable) {
        if (!item.columns)
            item.columns = []

        if (!item.constraints)
            item.constraints = []

        if (!item.options)
            item.options = []

        {
            let line = `CREATE TABLE ${item.name}`

            if (item.ifNotExists)
                line += ' IF NOT EXISTS'

            line += ' ('
            this.sb.l(line)
        }

        {

            this.sb.indent()
            const columns: string[] = []


            for (const column of item.columns) {
                const cText = `${column.column} ${column.type}`
                columns.push(cText)
            }

            if (item.constraints) {
                for (const constraint of item.constraints) {
                    let line = ``

                    switch (constraint.kind) {
                        case 'primary-key':
                            line += `PRIMARY KEY (${constraint.columns.join(',')})`
                            break
                        case 'unique-key':
                            line += `UNIQUE (${constraint.columns.join(',')})`
                            break
                        case 'foreign-key':
                            line += `FOREIGN KEY (${constraint.column}) REFERENCES ${this.refp(constraint.reference)}`

                            if (constraint.onDelete)
                                line += ` ON DELETE ${constraint.onDelete}`

                            if (constraint.onUpdate)
                                line += ` ON UPDATE ${constraint.onDelete}`

                            break
                        default:
                            throw new Error(`invalid constraint ${JSON.stringify(constraint)}`)
                    }

                    columns.push(line)
                }
            }

            for (let i = 0; i < columns.length; i++) {
                let line = columns[i]

                if (i < columns.length - 1)
                    line += ','

                this.sb.l(line)
            }

            this.sb.dedent()
        }
        this.sb.l(');')
    }
}

export interface SqlTable {
    name: string
    columns: SqlColumn[]
    constraints?: SqlTableConstraint[]
    options?: SqlTableOptions[]
    ifNotExists?: boolean
}

export interface SqlIndex {
    name?: string
    table: string
    column: string
    unique?: boolean
    ifNotExists?: boolean
}

export enum SqlColumnType {
    Text = 'TEXT',
    Numeric = 'NUMERIC',
    Integer = 'INTERGER',
    Real = 'REAL',
    Blob = 'BLOB',
}

export interface SqlColumn {
    column: string
    type: SqlColumnType
}

export type SqlTableConstraint =
    SqlTableConstraintPrimaryKey |
    SqlTableConstraintForeignKey |
    SqlTableConstraintUniqueKey

export enum SqlTableConstraintForeignKeyOn {
    SetNull = 'SET NULL',
    SetDefault = 'SET DEFAULT',
    Cascade = 'CASCADE',
    Restrict = 'RESTRICT',
    NoAction = 'NO ACTION',
}

export enum SqlTableOptions {
    WithoutRowId = 'WITHOU ROWID',
    Strict = 'STRICT',
}

export interface SqlTableIndexedColumn {
    column: string
    unique?: boolean
    ifNotExists?: boolean
}

export interface SqlTableConstraintPrimaryKey {
    kind: 'primary-key'
    columns: string[]
}

export interface SqlTableConstraintUniqueKey {
    kind: 'unique-key'
    columns: string[]
}

export interface SqlTableConstraintForeignKey {
    kind: 'foreign-key'
    column: string
    reference: SqlRference
    onDelete?: SqlTableConstraintForeignKeyOn
    onUpdate?: SqlTableConstraintForeignKeyOn
}

export interface SqlInsert {
    table: string
    values: {
        column: string
        value: SqlValue
    }[]
}

type SqlValue =
    SqlValueText |
    SqlValueNumeric |
    SqlValueInteger |
    SqlValueReal |
    SqlValueSqlParameter |
    SqlValueTsInterpolationNumeric |
    SqlValueTsInterpolationText |
    SqlValueBoolean |
    SqlValueDate

export interface SqlValueDate {
    kind: 'date'
    value: Date
    format: 'ISO-8601' | 'RFC-3339' | 'epoch'
}

export interface SqlValueBoolean {
    kind: 'boolean'
    value: boolean
}

export interface SqlValueSqlParameter {
    kind: 'sql-parameter'
    name?: string
    index?: number
}

export interface SqlValueTsInterpolationNumeric {
    kind: 'ts-interpolation-numeric'
    name: string
}

export interface SqlValueTsInterpolationText {
    kind: 'ts-interpolation-text'
    name: string
}

export interface SqlValueText {
    kind: 'text'
    value: string
}

export interface SqlValueNumeric {
    kind: 'numeric'
    value: number
}

export interface SqlValueInteger {
    kind: 'integer'
    value: number
}

export interface SqlValueReal {
    kind: 'real'
    value: number
}

export interface SqlUpdate {
    table: string
    values: {
        column: string
        value: SqlValue
    }[]
    where?: SqlExpr
}

export interface SqlDelete {
    table: string
    where?: SqlExpr
}

export interface SqlSelect {
    table: string
    columns: string[]
    where?: SqlExpr
    joins?: SqlJoin[]
    groupBy?: SqlExpr
    orderBy?: {
        column: string
        term?: SqlOrderingTerm
    }
    limit?: {
        expression: SqlExpr
        offset?: SqlExpr
    },
}

export enum SqlOrderingTerm {
    Asc = 'ASC',
    Desc = 'DESC'
}

export interface SqlRference {
    table: string
    column: string
}

export type SqlJoin = SqlJoinOn | SqlJoinUsing

export interface SqlJoinOn {
    kind: 'join-on'
    column: string
    reference: SqlRference
    operator?: SqlJoinOperator
}

export interface SqlJoinUsing {
    kind: 'join-using'
    reference: SqlRference
    operator?: SqlJoinOperator
}

export enum SqlJoinOperator {
    Natrual = 'NATURAL',
    Lelt = 'LEFT',
    Right = 'RIGHT',
    Full = 'FULL',
    Inner = 'INNER',
    Cross = 'CROSS'
}

export type SqlExpr =
    SqlExprLiteralValue |
    SqlExprBindParameter |
    SqlExprEqual |
    SqlExprNotEqual |
    SqlExprIs |
    SqlExprIsNot |
    SqlExprGreaterThan |
    SqlExprGreaterThanOrEqual |
    SqlExprLessThan |
    SqlExprLessThanOrEqual |
    SqlExprAnd |
    SqlExprOr |
    SqlExprBetween |
    SqlExprNot |
    SqlExprIn

export interface SqlExprLiteralValue {
    kind: 'literal-value'
    value: SqlValue
}

export interface SqlExprBindParameter {
    kind: 'bind-parameter'
}

export interface SqlExprEqual {
    kind: 'equal'
    column: string
    value: SqlValue
}

export interface SqlExprNotEqual {
    kind: 'not-equal'
    column: string
    value: SqlValue
}

export interface SqlExprIs {
    kind: 'is'
    column: string
    value: SqlValue
}

export interface SqlExprIsNot {
    kind: 'is-not'
    column: string
    value: SqlValue
}

export interface SqlExprGreaterThan {
    kind: 'greater-than'
    column: string
    value: SqlValue
}

export interface SqlExprGreaterThanOrEqual {
    kind: 'greater-than-or-equal'
    column: string
    value: SqlValue
}

export interface SqlExprLessThan {
    kind: 'less-than'
    column: string
    value: SqlValue
}

export interface SqlExprLessThanOrEqual {
    kind: 'less-than-or-equal'
    column: string
    value: SqlValue
}

export interface SqlExprIn {
    kind: 'in'
    column: string
    values: SqlValue[]
}

export interface SqlExprNot {
    kind: 'not'
    expression: SqlExpr
}

export interface SqlExprAnd {
    kind: 'and'
    expressions: SqlExpr[]
}

export interface SqlExprOr {
    kind: 'or'
    expressions: SqlExpr[]
}

export interface SqlExprBetween {
    kind: 'between'
    column: string
    upper: SqlValue
    lower: SqlValue
}
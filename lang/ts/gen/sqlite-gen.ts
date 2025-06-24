import { StringBuffer } from "./utils.ts";

export class SqliteGenerator {

    private value(sb: StringBuffer, item: SqlValue) {
        function escape(s: string): string {
            return s.replaceAll("'", "''")
        }

        switch (item.kind) {
            case 'integer':
            case 'numeric':
            case 'real':
                sb.a(`${item.value}`)
                break
            case 'text':
                sb.a(`'${escape(item.value)}'`)
                break
            case 'ts-interpolation-numeric':
                sb.a('${' + item.name + '}')
                break
            case 'ts-interpolation-text':
                sb.a("'${" + item.name + "}'")
                break
            case 'sql-parameter':
                if (item.index)
                    sb.a(`?${item.index}`)
                else if (item.name)
                    sb.a(`:${item.name}`)
                else
                    sb.a('?')
                break
            case 'boolean':
                sb.a(`${item.value}`)
                break
            case 'date':
                switch (item.format) {
                    case 'epoch':
                        this.value(sb, { kind: 'integer', value: item.value.getTime() })
                        break
                    case 'ISO-8601':
                        sb.a("TODO: ISO-8601")
                        break
                    case 'RFC-3339':
                        sb.a("TODO: RFC-3339")
                        break
                    default:
                        throw new Error(`invalid sql date format ${JSON.stringify(item)}`)
                }
                break
            default:
                throw new Error(`invalid sql value ${JSON.stringify(item)}`)
        }
    }

    private fqcn(column: SqlColumn): string {
        return `${column.table.name}.${column.name}`
    }

    getPrimaryKeys(table: SqlTable): SqlColumn[] {
        const keys: SqlColumn[] = []

        for (const column of table.columns) {
            if (column.primaryKey)
                keys.push(column)
        }

        if (table.constraints) {
            for (const constraint of table.constraints) {
                if (constraint.kind === 'primary-key')
                    keys.push(...constraint.columns)
            }
        }

        return keys
    }

    getColumnByName(table: SqlTable, name: string): SqlColumn | undefined {
        let found: SqlColumn | undefined

        for (const column of table.columns) {
            if (column.name === name) {
                found = column
                break
            }
        }

        return found
    }

    expression(sb: StringBuffer, item: SqlExpr): string {
        const handleList = (es: SqlExpr[], sep: string) => {
            for (let i = 0; i < es.length; i++) {
                this.expression(sb, es[i])

                if (i < es.length - 1)
                    sb.a(sep)
            }
        }

        switch (item.kind) {
            case 'equal':
                sb.a(`${this.fqcn(item.column)} = `)
                this.value(sb, item.value)
                break
            case 'not-equal':
                sb.a(`${this.fqcn(item.column)} != `)
                this.value(sb, item.value)
                break
            case 'is':
                sb.a(`${this.fqcn(item.column)} IS `)
                this.value(sb, item.value)
                break
            case 'is-not':
                sb.a(`${this.fqcn(item.column)} IS NOT `)
                this.value(sb, item.value)
                break
            case 'greater-than':
                sb.a(`${this.fqcn(item.column)} > `)
                this.value(sb, item.value)
                break
            case 'greater-than-or-equal':
                sb.a(`${this.fqcn(item.column)} >= `)
                this.value(sb, item.value)
                break
            case 'less-than':
                sb.a(`${this.fqcn(item.column)} > `)
                this.value(sb, item.value)
                break
            case 'less-than-or-equal':
                sb.a(`${this.fqcn(item.column)} <= `)
                this.value(sb, item.value)
                break
            case 'between':
                sb.a(`${this.fqcn(item.column)} BETWEEN `)
                this.value(sb, item.lower)
                sb.a(` AND `)
                this.value(sb, item.upper)
                break
            case 'not':
                sb.a(`NOT `)
                this.expression(sb, item.expression)
                break
            case 'and':
                handleList(item.expressions, ' AND ')
                break
            case 'or':
                handleList(item.expressions, ' OR ')
                break
            case 'literal-value':
                this.value(sb, item.value)
                break
            case 'in':
                {
                    sb.a(`${this.fqcn(item.column)} IN (`)

                    for (let i = 0; i < item.values.length; i++) {
                        const value = item.values[i]
                        this.value(sb, value)

                        if (i < item.values.length - 1)
                            sb.a(',')
                    }

                    sb.a(')')
                }
                break
            case 'bind-parameter':
                // TODO
                break
            default:
                throw new Error(`invalid expression ${JSON.stringify(item)}`)
        }

        return sb.toString()
    }

    update(sb: StringBuffer, item: SqlUpdate) {
        const keys = this.getPrimaryKeys(item.table)

        if (keys.length <= 0)
            throw new Error(`missing primary key for ${JSON.stringify(item)}`)

        sb.l(`UPDATE ${item.table.name} SET`)
        sb.nl()

        sb.i(() => {
            for (let i = 0; i < item.table.columns.length; i++) {
                const column = item.table.columns[i]

                if (column.name == keys[0].name)
                    continue

                sb.a(`${column.name} = `)

                this.value(sb, { kind: 'sql-parameter' })

                if (i < item.table.columns.length - 1) {
                    sb.a(', ')
                }

                sb.nl()
            }
        })

        sb.a(` WHERE `)

        this.expression(sb, {
            kind: 'equal',
            column: keys[0],
            value: { kind: 'sql-parameter' }
        })

        if (item.where) {
            sb.a(` AND `)
            this.expression(sb, item.where)
        }

        sb.l(';')
    }

    select(sb: StringBuffer, item: SqlSelect): string {
        sb.l(`SELECT `)
        sb.nl()

        sb.i(() => {
            const columns = item.columns ?? item.table.columns

            for (let i = 0; i < columns.length; i++) {
                sb.a(this.fqcn(columns[i]))

                if (i < columns.length - 1) {
                    sb.a(', ')
                    sb.nl()
                }
            }
        })

        sb.l(`FROM ${item.table.name}`)

        if (item.where) {
            sb.nl()
            sb.a(`WHERE `)

            sb.i(() => {
                if (!item.where) return
                this.expression(sb, item.where)
            })
        }

        if (item.joins && item.joins.length > 0) {
            sb.i(() => {
                for (const join of item.joins!) {
                    let jtext = ''

                    switch (join.kind) {
                        case 'join-on': {
                            if (join.operator)
                                jtext += ` ${join.operator}`

                            jtext += `JOIN ${join.reference.table.name} `
                            jtext += `ON ${this.fqcn(join.column)} = ${this.fqcn(join.reference)}`
                        }
                            break
                        case 'join-using':
                            if (join.operator)
                                jtext += ` ${join.operator}`

                            jtext += ` JOIN ${join.reference.table.name}`
                            jtext += ` USING (${join.reference.name})`
                            break
                        default:
                            throw new Error(`invalid join type ${JSON.stringify(join)}`)
                    }

                    sb.l(jtext)
                }
            })
        }

        if (item.orderBy) {
            sb.i(() => {
                if (!item.orderBy) return

                sb.nl()
                sb.a(`ORDER BY ${item.orderBy.column.name}`)

                if (item.orderBy.term)
                    sb.a(` ${item.orderBy.term}`)
            })
        }

        if (item.limit) {
            sb.nl()

            sb.i(() => {
                if (!item.limit) return
                sb.a(`LIMIT `)
                this.expression(sb, item.limit.expression)

                if (item.limit.offset) {
                    sb.a(` OFFSET `)
                    this.expression(sb, item.limit.offset)
                }
            })

            sb.nl()
        }

        sb.a(';')
        sb.nl()

        return sb.toString()
    }

    delete(sb: StringBuffer, item: SqlDelete) {
        const keys = this.getPrimaryKeys(item.table)

        if (keys.length <= 0)
            throw new Error(`missing primary key for ${JSON.stringify(item)}`)

        sb.a(`DELETE FROM ${item.table.name}`)

        sb.a(` WHERE `)

        const equalId: SqlExprEqual = {
            kind: 'equal',
            column: keys[0],
            value: { kind: 'sql-parameter' }
        }

        let where: SqlExpr

        if (item.where) {
            where = {
                kind: 'and',
                expressions: [equalId, item.where]
            } as SqlExprAnd
        } else {
            where = equalId
        }

        this.expression(sb, where)

        sb.a(';')
    }

    insert(sb: StringBuffer, item: SqlInsert) {
        sb.l(`INSERT INTO ${item.table.name} (`)
        sb.nl()

        if (item.values) {
            sb.i(() => {
                if (!item.values) return

                for (let i = 0; i < item.values.length; i++) {
                    const value = item.values[i]
                    sb.a(value.column.name)

                    if (i < item.values.length - 1) {
                        sb.a(', ')
                    }
                }
            })
            sb.l(`) VALUES (`)
            sb.nl()

            sb.i(() => {
                if (!item.values) return

                for (let i = 0; i < item.values.length; i++) {
                    const value = item.values[i]

                    if (value.value)
                        this.value(sb, value.value)
                    else
                        this.value(sb, { kind: 'sql-parameter' })

                    if (i < item.values.length - 1) {
                        sb.a(', ')
                    }
                }
            })
        } else {
            sb.i(() => {
                for (let i = 0; i < item.table.columns.length; i++) {
                    const column = item.table.columns[i]
                    sb.a(column.name)

                    if (i < item.table.columns.length - 1) {
                        sb.a(', ')
                    }
                }
            })
            sb.l(`) VALUES (`)
            sb.nl()

            sb.i(() => {


                for (let i = 0; i < item.table.columns.length; i++) {
                    this.value(sb, { kind: 'sql-parameter' })

                    if (i < item.table.columns.length - 1) {
                        sb.a(', ')
                    }
                }
            })
        }

        sb.l(');')
    }

    index(sb: StringBuffer, item: SqlIndex) {
        const prefix = item.prefix ?? 'idx'

        const create = (column: SqlColumn) => {
            sb.a('CREATE')

            if (column.unique)
                sb.a(' UNIQUE')

            sb.a(' INDEX')

            if (item.ifNotExists)
                sb.a(' IF NOT EXISTS')

            sb.a(` ${prefix}__${item.table.name}__${column.name}`)

            sb.a(` ON ${item.table.name} (${column.name})`)

            sb.a(';')
        }

        for (const column of item.table.columns) {
            if (column.indexed || column.searchable) {
                create(column)
                sb.nl()
            }
        }

        if (item.columns) {
            for (const column of item.columns) {
                create(column)
                sb.nl()
            }
        }

        return sb.toString()
    }

    table(sb: StringBuffer, item: SqlTable) {
        if (!item.columns)
            item.columns = []

        if (!item.constraints)
            item.constraints = []

        if (!item.options)
            item.options = []

        {
            sb.a(`CREATE TABLE ${item.name}`)

            if (item.ifNotExists)
                sb.a(' IF NOT EXISTS')

            sb.a(' (')
        }

        sb.i(() => {
            const columns: string[] = []

            for (const column of item.columns) {
                let cText = `${column.name} ${column.type}`

                if (!column.optional || column.primaryKey)
                    cText += ` NOT NULL`

                columns.push(cText)
            }

            const constraints: SqlTableConstraint[] = []

            for (const column of item.columns) {
                if (column.primaryKey)
                    constraints.push({ kind: 'primary-key', columns: [column] })

                if (column.unique)
                    constraints.push({ kind: 'unique-key', columns: [column] })
            }

            if (item.constraints)
                constraints.push(...item.constraints)


            for (const constraint of constraints) {
                let line = ``

                switch (constraint.kind) {
                    case 'primary-key':
                        line += `PRIMARY KEY (${constraint.columns.map(c => c.name).join(',')})`
                        break
                    case 'unique-key':
                        line += `UNIQUE (${constraint.columns.map(c => c.name).join(',')})`
                        break
                    case 'foreign-key':
                        line += `FOREIGN KEY (${constraint.column.name}) REFERENCES ${this.fqcn(constraint.reference)}`

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


            for (let i = 0; i < columns.length; i++) {
                let line = columns[i]

                if (i < columns.length - 1)
                    line += ','

                sb.l(line)
            }
        })

        sb.nl()
        sb.a(')')

        if (item.options) {
            for (const option of item.options) {
                sb.a(` ${option}`)
            }
        }

        sb.a(';')
        sb.nl()
    }
}

export interface SqlTable {
    name: string
    columns: SqlColumn[]
    constraints?: SqlTableConstraint[]
    options?: SqlTableOptions[]
    ifNotExists?: boolean
}

export enum SqlTableOptions {
    WithoutRowId = 'WITHOU ROWID',
    Strict = 'STRICT',
}

export interface SqlIndex {
    prefix?: string
    table: SqlTable
    columns?: SqlColumn[]
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
    name: string
    type: SqlColumnType
    optional?: boolean,
    indexed?: boolean
    unique?: boolean
    primaryKey?: boolean
    searchable?: boolean
    table: SqlTable
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

export interface SqlTableIndexedColumn {
    column: string
    unique?: boolean
    ifNotExists?: boolean
}

export interface SqlTableConstraintPrimaryKey {
    kind: 'primary-key'
    columns: SqlColumn[]
}

export interface SqlTableConstraintUniqueKey {
    kind: 'unique-key'
    columns: SqlColumn[]
}

export interface SqlTableConstraintForeignKey {
    kind: 'foreign-key'
    column: SqlColumn
    reference: SqlColumn
    onDelete?: SqlTableConstraintForeignKeyOn
    onUpdate?: SqlTableConstraintForeignKeyOn
}

export interface SqlInsert {
    table: SqlTable
    values?: {
        column: SqlColumn
        value?: SqlValue
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
    table: SqlTable
    where?: SqlExpr
}

export interface SqlDelete {
    table: SqlTable
    where?: SqlExpr
}

export interface SqlSelect {
    table: SqlTable
    columns?: SqlColumn[]
    where?: SqlExpr
    joins?: SqlJoin[]
    groupBy?: SqlExpr
    orderBy?: {
        column: SqlColumn
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

export type SqlJoin = SqlJoinOn | SqlJoinUsing

export interface SqlJoinOn {
    kind: 'join-on'
    column: SqlColumn
    reference: SqlColumn
    operator?: SqlJoinOperator
}

export interface SqlJoinUsing {
    kind: 'join-using'
    reference: SqlColumn
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
    column: SqlColumn
    value: SqlValue
}

export interface SqlExprNotEqual {
    kind: 'not-equal'
    column: SqlColumn
    value: SqlValue
}

export interface SqlExprIs {
    kind: 'is'
    column: SqlColumn
    value: SqlValue
}

export interface SqlExprIsNot {
    kind: 'is-not'
    column: SqlColumn
    value: SqlValue
}

export interface SqlExprGreaterThan {
    kind: 'greater-than'
    column: SqlColumn
    value: SqlValue
}

export interface SqlExprGreaterThanOrEqual {
    kind: 'greater-than-or-equal'
    column: SqlColumn
    value: SqlValue
}

export interface SqlExprLessThan {
    kind: 'less-than'
    column: SqlColumn
    value: SqlValue
}

export interface SqlExprLessThanOrEqual {
    kind: 'less-than-or-equal'
    column: SqlColumn
    value: SqlValue
}

export interface SqlExprIn {
    kind: 'in'
    column: SqlColumn
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
    column: SqlColumn
    upper: SqlValue
    lower: SqlValue
}

const SQL_LITE_KEYWORDS = [
    "ABORT",
    "ACTION",
    "ADD",
    "AFTER",
    "ALL",
    "ALTER",
    "ALWAYS",
    "ANALYZE",
    "AND",
    "AS",
    "ASC",
    "ATTACH",
    "AUTOINCREMENT",
    "BEFORE",
    "BEGIN",
    "BETWEEN",
    "BY",
    "CASCADE",
    "CASE",
    "CAST",
    "CHECK",
    "COLLATE",
    "COLUMN",
    "COMMIT",
    "CONFLICT",
    "CONSTRAINT",
    "CREATE",
    "CROSS",
    "CURRENT",
    "CURRENT_DATE",
    "CURRENT_TIME",
    "CURRENT_TIMESTAMP",
    "DATABASE",
    "DEFAULT",
    "DEFERRABLE",
    "DEFERRED",
    "DELETE",
    "DESC",
    "DETACH",
    "DISTINCT",
    "DO",
    "DROP",
    "EACH",
    "ELSE",
    "END",
    "ESCAPE",
    "EXCEPT",
    "EXCLUDE",
    "EXCLUSIVE",
    "EXISTS",
    "EXPLAIN",
    "FAIL",
    "FILTER",
    "FIRST",
    "FOLLOWING",
    "FOR",
    "FOREIGN",
    "FROM",
    "FULL",
    "GENERATED",
    "GLOB",
    "GROUP",
    "GROUPS",
    "HAVING",
    "IF",
    "IGNORE",
    "IMMEDIATE",
    "IN",
    "INDEX",
    "INDEXED",
    "INITIALLY",
    "INNER",
    "INSERT",
    "INSTEAD",
    "INTERSECT",
    "INTO",
    "IS",
    "ISNULL",
    "JOIN",
    "KEY",
    "LAST",
    "LEFT",
    "LIKE",
    "LIMIT",
    "MATCH",
    "MATERIALIZED",
    "NATURAL",
    "NO",
    "NOT",
    "NOTHING",
    "NOTNULL",
    "NULL",
    "NULLS",
    "OF",
    "OFFSET",
    "ON",
    "OR",
    "ORDER",
    "OTHERS",
    "OUTER",
    "OVER",
    "PARTITION",
    "PLAN",
    "PRAGMA",
    "PRECEDING",
    "PRIMARY",
    "QUERY",
    "RAISE",
    "RANGE",
    "RECURSIVE",
    "REFERENCES",
    "REGEXP",
    "REINDEX",
    "RELEASE",
    "RENAME",
    "REPLACE",
    "RESTRICT",
    "RETURNING",
    "RIGHT",
    "ROLLBACK",
    "ROW",
    "ROWS",
    "SAVEPOINT",
    "SELECT",
    "SET",
    "TABLE",
    "TEMP",
    "TEMPORARY",
    "THEN",
    "TIES",
    "TO",
    "TRANSACTION",
    "TRIGGER",
    "UNBOUNDED",
    "UNION",
    "UNIQUE",
    "UPDATE",
    "USING",
    "VACUUM",
    "VALUES",
    "VIEW",
    "VIRTUAL",
    "WHEN",
    "WHERE",
    "WINDOW",
    "WITH",
    "WITHOUT",
]
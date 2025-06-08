
interface SqlTable {
    name: string
    colunms: SqlColumn[]
    constraints?: SqlTableConstraint[]
    options?: SqlTableOptions[]
    ifNotExists?: boolean
}

interface SqlIndex {
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

interface SqlColumn {
    column: string
    type: SqlColumnType
}

type SqlTableConstraint = SqlTableConstraintPrimaryKey | SqlTableConstraintForeignKey | SqlTableConstraintUniqueKey

enum SqlTableConstraintForeignKeyOn {
    SetNull = 'SET NULL',
    SetDefault = 'SET DEFAULT',
    Cascade = 'CASCADE',
    Restrict = 'RESTRICT',
    NoAction = 'NO ACTION',
}

enum SqlTableOptions {
    WithoutRowId = 'WITHOU ROWID',
    Strict = 'STRICT',
}

interface SqlTableIndexedColumn {
    column: string
    unique?: boolean
    ifNotExists?: boolean
}

interface SqlTableConstraintPrimaryKey {
    kind: 'primary-key'
    columns: string[]
}

interface SqlTableConstraintUniqueKey {
    kind: 'unique-key'
    columns: string[]
}

interface SqlTableConstraintForeignKey {
    kind: 'foreign-key'
    column: string
    reference: SqlRference
    onDelete?: SqlTableConstraintForeignKeyOn
    onUpdate?: SqlTableConstraintForeignKeyOn
}

interface SqlInsert {
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

interface SqlValueDate {
    kind: 'date'
    value: Date
    format: 'ISO-8601' | 'RFC-3339' | 'epoch'
}

interface SqlValueBoolean {
    kind: 'boolean'
    value: boolean
}

interface SqlValueSqlParameter {
    kind: 'sql-parameter'
    name?: string
    index?: number
}

interface SqlValueTsInterpolationNumeric {
    kind: 'ts-interpolation-numeric'
    name: string
}

interface SqlValueTsInterpolationText {
    kind: 'ts-interpolation-text'
    name: string
}

interface SqlValueText {
    kind: 'text'
    value: string
}

interface SqlValueNumeric {
    kind: 'numeric'
    value: number
}

interface SqlValueInteger {
    kind: 'integer'
    value: number
}

interface SqlValueReal {
    kind: 'real'
    value: number
}

interface SqlUpdate {
    table: string
    values: {
        column: string
        value: SqlValue
    }[]
    where?: SqlExpr
}

interface SqlDelete {
    table: string
    where?: SqlExpr
}

interface SqlSelect {
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


interface SqlRference {
    table: string
    column: string
}

type SqlJoin = SqlJoinOn | SqlJoinUsing

interface SqlJoinOn {
    kind: 'join-on'
    column: string
    reference: SqlRference
    operator?: SqlJoinOperator
}

interface SqlJoinUsing {
    kind: 'join-using'
    reference: SqlRference
    operator?: SqlJoinOperator
}

enum SqlJoinOperator {
    Natrual = 'NATURAL',
    Lelt = 'LEFT',
    Right = 'RIGHT',
    Full = 'FULL',
    Inner = 'INNER',
    Cross = 'CROSS'
}

type SqlExpr =
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

interface SqlExprLiteralValue {
    kind: 'literal-value'
    value: SqlValue
}

interface SqlExprBindParameter {
    kind: 'bind-parameter'
}

interface SqlExprEqual {
    kind: 'equal'
    column: string
    value: SqlValue
}

interface SqlExprNotEqual {
    kind: 'not-equal'
    column: string
    value: SqlValue
}

interface SqlExprIs {
    kind: 'is'
    column: string
    value: SqlValue
}

interface SqlExprIsNot {
    kind: 'is-not'
    column: string
    value: SqlValue
}

interface SqlExprGreaterThan {
    kind: 'greater-than'
    column: string
    value: SqlValue
}

interface SqlExprGreaterThanOrEqual {
    kind: 'greater-than-or-equal'
    column: string
    value: SqlValue
}

interface SqlExprLessThan {
    kind: 'less-than'
    column: string
    value: SqlValue
}

interface SqlExprLessThanOrEqual {
    kind: 'less-than-or-equal'
    column: string
    value: SqlValue
}

interface SqlExprIn {
    kind: 'in'
    column: string
    values: SqlValue[]
}

interface SqlExprNot {
    kind: 'not'
    expression: SqlExpr
}

interface SqlExprAnd {
    kind: 'and'
    expressions: SqlExpr[]
}

interface SqlExprOr {
    kind: 'or'
    expressions: SqlExpr[]
}

interface SqlExprBetween {
    kind: 'between'
    column: string
    upper: SqlValue
    lower: SqlValue
}

export class SqliteGenerator {
    private value(item: SqlValue): string {
        function escape(s: string): string {
            return s.replaceAll("'", "''")
        }

        let text: string

        switch (item.kind) {
            case 'integer':
            case 'numeric':
            case 'real':
                text = `${item.value}`
                break
            case 'text':
                text = `'${escape(item.value)}'`
                break
            case 'ts-interpolation-numeric':
                text = '${' + item.name + '}'
                break
            case 'ts-interpolation-text':
                text = "'${" + item.name + "}'"
                break
            case 'sql-parameter':
                if (item.index)
                    text = `?${item.index}`
                else if (item.name)
                    text = `:${item.name}`
                else
                    text = '?'
                break
            case 'boolean':
                text = `${item.value}`
                break
            case 'date':
                switch (item.format) {
                    case 'epoch':
                        text = this.value({ kind: 'integer', value: item.value.getTime() })
                        break
                    case 'ISO-8601':
                        text = "TODO: ISO-8601"
                        break
                    case 'RFC-3339':
                        text = "TODO: RFC-3339"
                        break
                    default:
                        throw new Error(`invalid sql date format ${JSON.stringify(item)}`)
                }
                break
            default:
                throw new Error(`invalid sql value ${JSON.stringify(item)}`)
        }

        return text
    }

    private refd(item: SqlRference): string { return `${item.table}.${item.column}` }

    private refp(item: SqlRference): string { return `${item.table}(${item.column})` }

    expression(item: SqlExpr): string {
        let text = ''


        const handleList = (es: SqlExpr[]): string[] => {
            const results: string[] = []

            for (const e of es) {
                const or = this.expression(e)
                results.push(or)
            }

            return results
        }

        switch (item.kind) {
            case 'equal':
                text = `${item.column} = ${this.value(item.value)}`
                break
            case 'not-equal':
                text = `${item.column} != ${this.value(item.value)}`
                break
            case 'is':
                text = `${item.column} IS ${this.value(item.value)}`
                break
            case 'is-not':
                text = `${item.column} IS NOT ${this.value(item.value)}`
                break
            case 'greater-than':
                text = `${item.column} > ${this.value(item.value)}`
                break
            case 'greater-than-or-equal':
                text = `${item.column} >= ${this.value(item.value)}`
                break
            case 'less-than':
                text = `${item.column} > ${this.value(item.value)}`
                break
            case 'less-than-or-equal':
                text = `${item.column} <= ${this.value(item.value)}`
                break
            case 'between':
                text = `${item.column} BETWEEN ${this.value(item.lower)} AND ${this.value(item.upper)}`
                break
            case 'not':
                text = `NOT ${this.expression(item.expression)}`
                break
            case 'and':
                {
                    const ands = handleList(item.expressions)
                    text += ands.join(' AND ')
                }
                break
            case 'or':
                {
                    const ors = handleList(item.expressions)
                    text += ors.join(' OR ')
                }
                break
            case 'literal-value':
                text = this.value(item.value)
                break
            case 'in':
                {
                    const values: string[] = []

                    for (const value of item.values)
                        values.push(this.value(value))

                    text = `${item.column} IN (${values.join(',')})`
                }
                break
            case 'bind-parameter':
                // TODO
                break
            default:
                throw new Error(`invalid expression ${JSON.stringify(item)}`)
        }
        return text
    }

    updates(items: SqlUpdate[]): string {
        const updates: string[] = []

        if (items) {
            for (const item of items) {
                updates.push(this.update(item))
            }
        }

        return updates.join('\n')
    }

    update(item: SqlUpdate): string {
        let text = `UPDATE ${item.table} SET`

        if (item.values.length > 0) {
            text += '\n'

            const pairs: string[] = []

            for (const value of item.values) {
                const pText = `${value.column} = ${this.value(value.value)}`
                pairs.push(pText)
            }

            text += pairs.join(',')
        }

        if (item.where) {
            text += ` WHERE ${this.expression(item.where)}`
        }

        text += ';\n'

        return text
    }

    selects(items: SqlSelect[]): string {
        const selects: string[] = []

        if (items) {
            for (const item of items) {
                selects.push(this.select(item))
            }
        }

        return selects.join('\n')
    }

    select(item: SqlSelect): string {
        let text = `SELECT ${item.columns.join(', ')}`
        text += ` FROM ${item.table}`

        if (item.where)
            text += ` WHERE ${this.expression(item.where)}`

        if (item.joins && item.joins.length > 0) {
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

                        jtext += ` JOIN ${join.reference.table}`
                        jtext += ` ON ${this.refd(tableRef)} = ${this.refd(join.reference)}`
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



                text += jtext
            }
        }

        if (item.orderBy) {
            text += ` ORDER BY ${item.orderBy.column}`

            if (item.orderBy.term)
                text += ` ${item.orderBy.term}`
        }

        if (item.limit) {
            text += ` LIMIT ${this.expression(item.limit.expression)}`

            if (item.limit.offset)
                text += ` OFFSET ${this.expression(item.limit.offset)}`
        }

        text += ';\n'

        return text
    }

    deletes(items: SqlDelete[]): string {
        const deletes: string[] = []

        if (items) {
            for (const item of items) {
                deletes.push(this.delete(item))
            }
        }

        return deletes.join('\n')
    }


    delete(item: SqlDelete): string {
        let text = `DELETE FROM ${item.table}`

        if (item.where) {
            text += ` WHERE ${this.expression(item.where)}`
        }

        text += ';\n'

        return text
    }

    inserts(items: SqlInsert[]): string {
        const inserts: string[] = []

        if (items) {
            for (const item of items) {
                inserts.push(this.insert(item))
            }
        }

        return inserts.join('\n')
    }

    insert(item: SqlInsert): string {
        let text = `INSERT INTO ${item.table} (`

        if (item.values.length > 0) {
            text += '\n'

            const columns: string[] = []
            const values: string[] = []

            for (const value of item.values) {
                columns.push(value.column)
                values.push(this.value(value.value))
            }

            text += `(${columns.join(',')}) \n`
            text += `(${values.join(',')});\n`
        }

        return text
    }

    indexes(items: SqlIndex[]): string {
        const indexes: string[] = []

        if (items) {
            for (const item of items) {
                indexes.push(this.index(item))
            }
        }

        return indexes.join('\n')
    }

    index(item: SqlIndex): string {
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

        return text
    }

    tables(items: SqlTable[]): string {
        const tables: string[] = []

        if (items) {
            for (const item of items) {
                tables.push(this.table(item))
            }
        }

        return tables.join('\n')
    }

    table(item: SqlTable): string {
        if (!item.colunms)
            item.colunms = []

        if (!item.constraints)
            item.constraints = []

        if (!item.options)
            item.options = []

        let text = `CREATE TABLE ${item.name}`

        if (item.ifNotExists)
            text += ' IF NOT EXISTS'

        text += ' ('

        const columns: string[] = []
        const constraints: string[] = []

        for (const column of item.colunms) {
            const cText = `${column.column} ${column.type}`
            columns.push(cText)
        }

        if (columns.length > 0) {
            text += '\n'
            text += columns.join(',\n')
        }

        if (item.constraints.length > 0) {
            text += '\n'

            for (const constraint of item.constraints) {
                let cText = ``

                switch (constraint.kind) {
                    case 'primary-key':
                        cText += `PRIMARY KEY (${constraint.columns.join(',')})`
                        break
                    case 'unique-key':
                        cText += `UNIQUE (${constraint.columns.join(',')})`
                        break
                    case 'foreign-key':
                        cText += `FOREIGN KEY (${constraint.column}) REFERENCES ${this.refp(constraint.reference)}`

                        if (constraint.onDelete)
                            cText += ` ON DELETE ${constraint.onDelete}`

                        if (constraint.onUpdate)
                            cText += ` ON UPDATE ${constraint.onDelete}`

                        break
                    default:
                        throw new Error(`invalid constraint ${JSON.stringify(constraint)}`)
                }
                constraints.push(cText)
            }

            if (constraints.length > 0) {
                text += constraints.join(',\n')
            }
        }

        if (columns.length > 0 || constraints.length > 0)
            text += '\n'

        text += ');'

        return text
    }
}
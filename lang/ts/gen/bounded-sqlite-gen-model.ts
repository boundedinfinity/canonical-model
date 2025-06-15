import * as BI from './bounded-model.ts'
import { TsNamer } from './namer.ts'
import { StringBuffer, type StringBufferOptions } from "./utils.ts";
import * as SQL from './sqlite-gen-model.ts'

type BoundedSqlGeneratorOptions = StringBufferOptions

export class BoundedSqliteGenerator {
    options: BoundedSqlGeneratorOptions
    bounded: BI.BoundedGenerator
    namer = new TsNamer()
    sql: SQL.SqliteGenerator
    sb: StringBuffer

    constructor(bounded: BI.BoundedGenerator, sb: StringBuffer, options?: Partial<BoundedSqlGeneratorOptions>) {
        this.bounded = bounded
        this.options = {
            indent: 0,
            tabSize: 4,
            ...options
        }
        this.sb = sb
        this.sql = new SQL.SqliteGenerator(this.sb)
    }

    genTable(type: BI.BoundedObject) {
        const table = this.resolvePrimaryTable(type)
        this.sql.table(table)

        this.genIndexes(type)

        const childTables = this.resolveChildTables(type)

        for (const childTable of childTables)
            this.sql.table(childTable)
    }

    private resolveSqlType(type: BI.BoundedType): SQL.SqlColumnType | undefined {
        let ctype: SQL.SqlColumnType | undefined

        switch (type.kind) {
            case 'string':
                ctype = SQL.SqlColumnType.Text
                break
            case 'boolean':
            case 'integer':
                ctype = SQL.SqlColumnType.Integer
                break
            case 'number':
                ctype = SQL.SqlColumnType.Numeric
                break
            case 'array':
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        return ctype
    }

    resolvePrimaryTable(type: BI.BoundedObject): SQL.SqlTable {
        const name = this.resolveSqlTable(type)
        const columns: SQL.SqlColumn[] = []
        const constraints: SQL.SqlTableConstraint[] = []

        for (const prop of type.properties) {
            const pname = this.resolveSqlColumn(prop)
            const ptype = this.resolveSqlType(prop)

            if (!ptype)
                continue

            columns.push({ column: pname, type: ptype })

            switch (prop.kind) {
                case 'string':
                case 'integer':
                case 'number':
                    if (prop.primaryKey)
                        constraints.push({ kind: 'primary-key', columns: [pname] })

                    if (prop.unique)
                        constraints.push({ kind: 'unique-key', columns: [pname] })
                    break
                case 'ref':
                case 'array':
                    continue
                default:
                    throw new Error(`invalid type ${JSON.stringify(type)}`)
            }
        }

        return {
            name,
            columns,
            constraints,
        }
    }

    private resolveChildTables(type: BI.BoundedObject): SQL.SqlTable[] {
        const tables: SQL.SqlTable[] = []
        const parentName = this.resolveSqlTable(type)

        for (const prop of type.properties) {
            switch (prop.kind) {
                case 'string':
                case 'integer':
                case 'number':
                case 'boolean':
                case 'ref':
                    continue
                case 'array':
                    switch (prop.items.kind) {
                        case 'number':
                        case 'integer':
                        case 'string':
                            break
                        case 'ref':
                            continue
                        default:
                            throw new Error(`invalid type ${JSON.stringify(type)}`)
                    }
                    break
                default:
                    throw new Error(`invalid type ${JSON.stringify(type)}`)
            }

            const columnType = this.resolveSqlType(prop.items)
            if (!columnType)
                continue

            const columnName = this.resolveSqlTable(prop)
            const childName = `${parentName}__${columnName}`

            const colunms: SQL.SqlColumn[] = [
                {
                    column: `${parentName}__id`,
                    type: SQL.SqlColumnType.Text
                },
                {
                    column: columnName,
                    type: columnType
                },
                {
                    column: 'index',
                    type: SQL.SqlColumnType.Integer
                }
            ]

            const constraints: SQL.SqlTableConstraint[] = [
                {
                    kind: 'foreign-key',
                    column: `${parentName}__id`,
                    reference: {
                        column: 'id',
                        table: parentName
                    }
                }
            ]

            tables.push({
                name: childName,
                columns: colunms,
                constraints
            })
        }

        return tables
    }

    genSelect(type: BI.BoundedType) {
        switch (type.kind) {
            case 'object':
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        const table = this.resolveSqlTable(type)
        const columns: string[] = []
        const ors: SQL.SqlExpr[] = []

        for (const prop of type.properties) {
            const cname = this.resolveSqlColumn(prop)
            columns.push(cname)

            switch (prop.kind) {
                case 'integer':
                case 'number':
                case 'string':
                    if (prop.searchable) {
                        const expr: SQL.SqlExprEqual = {
                            kind: 'equal',
                            column: cname,
                            value: { kind: 'sql-parameter' }
                        }
                        ors.push(expr)
                    }
                    break
                case 'array':
                    break
                default:
                    throw new Error(`invalid type ${JSON.stringify(type)}`)
            }
        }


        this.sql.select({
            table,
            columns
        })
    }

    private genIndexes(type: BI.BoundedObject) {
        const tname = this.resolveSqlTable(type)

        for (const prop of type.properties) {
            this.sb.nl()
            const pname = this.resolveSqlColumn(prop)

            switch (prop.kind) {
                case 'string':
                case 'integer':
                case 'number':
                    if (prop.indexed || prop.primaryKey || prop.searchable) {
                        this.sql.index({
                            column: pname,
                            table: tname,
                            unique: prop.unique
                        })
                    }
                    break
                case 'ref':
                    continue
                case 'array':
                    break
                default:
                    throw new Error(`invalid type ${JSON.stringify(type)}`)
            }
        }
    }

    private genJoinTable(type: BI.BoundedObject) {
        const tname = this.resolveSqlTable(type)

        for (const prop of type.properties) {
            switch (prop.kind) {
                case 'string':
                case 'integer':
                case 'number':
                case 'boolean':
                case 'ref':
                    continue
                case 'array':
                    switch (prop.items.kind) {
                        case 'number':
                        case 'integer':
                        case 'string':
                            {
                                const cname = this.resolveSqlTable(prop)
                                const ctype = this.bounded.getType(prop)
                                const name = `${tname}__${cname}`
                                const id = `${tname}__id`
                                const amodel: BI.BoundedObject = {
                                    kind: 'object',
                                    name,
                                    properties: [
                                        { kind: 'string', name: id },
                                        ctype,
                                        { kind: 'number', name: 'index' }
                                    ]
                                }

                                this.genTable(amodel)
                            }
                            break
                        case 'ref':
                            {
                                const cname = this.resolveSqlTable(prop)
                                const ctype = this.bounded.getType(prop)
                                const name = `${tname}___${cname}`
                                const jmodel: BI.BoundedObject = {
                                    kind: 'object',
                                    properties: []
                                }
                            }
                            break
                        default:
                            throw new Error(`invalid type ${JSON.stringify(type)}`)
                    }
                    break
                default:
                    throw new Error(`invalid type ${JSON.stringify(type)}`)
            }
        }
    }

    private resolvePrimaryKey(type: BI.BoundedObject): BI.BoundedType {
        let found: BI.BoundedType | undefined

        for (const prop of type.properties) {
            const pname = this.resolveSqlColumn(prop)

            switch (prop.kind) {
                case 'string':
                case 'integer':
                case 'number':
                    break
                case 'ref':
                case 'array':
                    continue
                default:
                    throw new Error(`invalid type ${JSON.stringify(type)}`)
            }

            if (prop.primaryKey) {
                found = prop
                break
            }
        }

        if (!found)
            throw new Error(`invalid type ${JSON.stringify(type)}`)

        return found
    }

    private resolveSqlTable(type: BI.BoundedType): string {
        let name = this.bounded.resolveName(type)

        if (name === '')
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`)

        name = this.namer.sqlTable(name)
        return name
    }

    private resolveSqlColumn(type: BI.BoundedType): string {
        let name = this.bounded.resolveName(type)

        if (name === '')
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`)

        name = this.namer.sqlColumn(name)
        return name
    }
}
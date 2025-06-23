import * as BI from './bounded-model.ts'
import { BoundedGenerator } from './bounded-gen.ts'
import { TsNamer } from './namer.ts'
import { type StringBufferOptions, ce } from "./utils.ts";
import * as SQL from './sqlite-gen.ts'

type BoundedSqlGeneratorOptions = StringBufferOptions

export class BoundedSqliteGenerator {
    options: BoundedSqlGeneratorOptions
    bounded: BoundedGenerator
    namer = new TsNamer()
    sql: SQL.SqliteGenerator


    constructor(bounded: BoundedGenerator, options?: Partial<BoundedSqlGeneratorOptions>) {
        this.bounded = bounded
        this.options = {
            indent: 0,
            tabSize: 4,
            ...options
        }

        this.sql = new SQL.SqliteGenerator()
    }

    genTable(type: BI.BoundedObject) {
        const primaryTable = this.getPrimaryTable(type)
        this.sql.table(primaryTable)
        this.sql.index({ table: primaryTable })

        const childTables = this.getChildTables(type)
        for (const childTable of childTables) {
            this.sql.table(childTable)
            this.sql.index({ table: childTable })
        }
    }

    private getSqlType(type: BI.BoundedType): SQL.SqlColumnType | undefined {
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

    getPrimaryTable(type: BI.BoundedObject): SQL.SqlTable {
        const name = this.getSqlTableName(type)
        const table: SQL.SqlTable = {
            name,
            columns: [],
            constraints: [],
        }

        for (const prop of type.properties) {
            const pname = this.getSqlColumnName(prop)
            const ptype = this.getSqlType(prop)

            if (!ptype)
                continue

            const column: SQL.SqlColumn = {
                name: pname,
                type: ptype,
                optional: prop.optional,
                table
            }

            switch (prop.kind) {
                case 'string':
                case 'integer':
                case 'number':
                    column.indexed = prop.indexed
                    column.primaryKey = prop.primaryKey
                    column.unique = prop.unique
                    column.searchable = prop.searchable
                    break
                case 'ref':
                case 'array':
                    continue
                default:
                    throw new Error(`invalid type ${JSON.stringify(type)}`)
            }

            table.columns.push(column)
        }

        return table
    }

    private getChildTables(type: BI.BoundedObject): SQL.SqlTable[] {
        const tables: SQL.SqlTable[] = []
        const primaryTable = this.getPrimaryTable(type)
        const primaryKeys = this.sql.getPrimaryKeys(primaryTable)
        const primaryTablePrimaryKey = primaryKeys[0]!

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
                            if (prop.items.optional)
                                throw ce(`array item cannot be optional`, prop)
                            if (prop.items.unique)
                                throw ce(`array item cannot be unique`, prop)
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

            const columnType = this.getSqlType(prop.items)

            if (!columnType)
                continue

            const columnName = this.getSqlTableName(prop)
            const childName = `${primaryTable.name}__${columnName}`

            const childTable: SQL.SqlTable = {
                name: childName,
                columns: [],
                constraints: []
            }

            const childPrimaryKey: SQL.SqlColumn = {
                name: `${primaryTable.name}__${primaryTablePrimaryKey.name}`,
                type: SQL.SqlColumnType.Text,
                indexed: false,
                primaryKey: true,
                table: childTable
            }

            const columns: SQL.SqlColumn[] = [
                childPrimaryKey,
                {
                    name: columnName,
                    type: columnType,
                    indexed: prop.items.searchable,
                    // optional: prop.items.optional,
                    // unique: prop.items.unique,
                    searchable: prop.items.searchable,
                    table: childTable,
                },
                {
                    name: 'index',
                    type: SQL.SqlColumnType.Integer,
                    table: childTable,
                }
            ]

            childTable.columns.push(...columns)

            const constraints: SQL.SqlTableConstraint[] = [
                {
                    kind: 'foreign-key',
                    column: childPrimaryKey,
                    reference: primaryTablePrimaryKey,
                }
            ]

            childTable.constraints?.push(...constraints)

            tables.push(childTable)
        }

        return tables
    }

    genUpdate(type: BI.BoundedType) {
        switch (type.kind) {
            case 'object':
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        const primaryTable = this.getPrimaryTable(type)
        const childTables = this.getChildTables(type)

        for (const table of [primaryTable, ...childTables])
            this.sql.update({ table })
    }

    genSelect(type: BI.BoundedType) {
        switch (type.kind) {
            case 'object':
                break
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`)
        }

        const primaryTable = this.getPrimaryTable(type)
        const primaryTableKey = this.sql.getPrimaryKeys(primaryTable)![0]
        const childTables = this.getChildTables(type)
        const joins: SQL.SqlJoin[] = []
        const expressions: SQL.SqlExpr[] = []

        for (const column of primaryTable.columns) {
            if (column.searchable || column.primaryKey) {
                switch (column.type) {
                    case SQL.SqlColumnType.Text:
                        expressions.push({
                            kind: 'equal',
                            column: column,
                            value: { kind: 'sql-parameter' }
                        })
                        break
                    default:
                    // TODO
                }
            }
        }

        for (const childTable of childTables) {
            for (const column of childTable.columns) {
                if (column.searchable) {
                    switch (column.type) {
                        case SQL.SqlColumnType.Text:
                            expressions.push({
                                kind: 'equal',
                                column: column,
                                value: { kind: 'sql-parameter' }
                            })

                            joins.push({
                                kind: 'join-on',
                                column: primaryTableKey,
                                reference: column
                            })
                            break
                        default:
                        // TODO
                    }
                }
            }
        }

        const select: SQL.SqlSelect = {
            table: primaryTable,
            joins,
            where: { kind: 'or', expressions }
        }

        this.sql.select(select)
    }

    private genJoinTable(type: BI.BoundedObject) {
        const tname = this.getSqlTableName(type)

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
                                const cname = this.getSqlTableName(prop)
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
                                const cname = this.getSqlTableName(prop)
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

    // private getPrimaryKey(type: BI.BoundedObject): BI.BoundedType {
    //     let found: BI.BoundedType | undefined

    //     for (const prop of type.properties) {
    //         const pname = this.resolveSqlColumnName(prop)

    //         switch (prop.kind) {
    //             case 'string':
    //             case 'integer':
    //             case 'number':
    //                 break
    //             case 'ref':
    //             case 'array':
    //                 continue
    //             default:
    //                 throw new Error(`invalid type ${JSON.stringify(type)}`)
    //         }

    //         if (prop.primaryKey) {
    //             found = prop
    //             break
    //         }
    //     }

    //     if (!found)
    //         throw new Error(`invalid type ${JSON.stringify(type)}`)

    //     return found
    // }

    private getSqlTableName(type: BI.BoundedType): string {
        let name = this.bounded.resolveName(type)

        if (name === '')
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`)

        name = this.namer.sqlTable(name)
        return name
    }

    private getSqlColumnName(type: BI.BoundedType): string {
        let name = this.bounded.resolveName(type)

        if (name === '')
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`)

        name = this.namer.sqlColumn(name)
        return name
    }
}
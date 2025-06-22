import { assertEquals } from "@std/assert";
import * as SQL from './sqlite-gen-model.ts'
import { StringBuffer, stringUtils as su } from './utils.ts'

function save(file: string, text: string) {
    const encoder = new TextEncoder();
    Deno.writeFileSync(file, encoder.encode(text))
}

function createTable1(): SQL.SqlTable {
    return {
        name: 'name_prefix',
        columns: [
            { name: 'id', type: SQL.SqlColumnType.Text, primaryKey: true },
            { name: 'name', type: SQL.SqlColumnType.Text, indexed: true, unique: true },
            { name: 'description', type: SQL.SqlColumnType.Text, optional: true }
        ],
    }
}

function createTable2(): SQL.SqlTable {
    return {
        name: 'name_prefix__abbr',
        columns: [
            { name: 'name_prefix__id', type: SQL.SqlColumnType.Text, primaryKey: true },
            { name: 'abbr', type: SQL.SqlColumnType.Text, indexed: true },
            { name: 'index', type: SQL.SqlColumnType.Integer },
        ],
        constraints: []
    }
}

Deno.test('Generate create table', () => {
    const sb = new StringBuffer()
    const sql = new SQL.SqliteGenerator(sb)
    const table1 = createTable1()
    const table2 = createTable2()

    table2.constraints?.push({
        kind: 'foreign-key',
        column: sql.getPrimaryKeys(table2)[0],
        reference: {
            table: table1,
            column: sql.getPrimaryKeys(table1)[0]
        }
    })

    sql.tables([table1, table2])

    save('./gen-output/generate-table.gen.sql', sb.toString())

    assertEquals(su.test.normal(sb.toString()), su.test.normal(`
        Teest
    `))
})

Deno.test('Generate create index', () => {
    const table1 = createTable1()
    const table2 = createTable2()
    const sb = new StringBuffer()
    const sql = new SQL.SqliteGenerator(sb)

    sql.indexes([
        {
            table: table1,
        },
        {
            table: table2,
            columns: [sql.getColumnByName(table2, 'abbr')!],
        },
    ])

    save('./gen-output/generate-index.gen.sql', sb.toString())

    assertEquals(su.test.normal(sb.toString()), su.test.normal(`
        Test
    `))
})

Deno.test('Generate insert', () => {
    const sb = new StringBuffer()
    const sql = new SQL.SqliteGenerator(sb)
    const table1 = createTable1()
    const table2 = createTable2()
    sql.inserts([{ table: table1 }, { table: table2 }])

    save('./gen-output/generate-insert.gen.sql', sb.toString())

    assertEquals(su.test.normal(sb.toString()), su.test.normal(`
        Teest
    `))
})

Deno.test('Generate update', () => {
    const sb = new StringBuffer()
    const sql = new SQL.SqliteGenerator(sb)
    const table1 = createTable1()
    const table2 = createTable2()
    sql.updates([{ table: table1 }, { table: table2 }])

    save('./gen-output/generate-update.gen.sql', sb.toString())

    assertEquals(su.test.normal(sb.toString()), su.test.normal(`
        Teest
    `))
})

Deno.test('Generate delete', () => {
    const sb = new StringBuffer()
    const sql = new SQL.SqliteGenerator(sb)
    const table1 = createTable1()
    const table2 = createTable2()
    sql.deletes([{ table: table1 }, { table: table2 }])

    save('./gen-output/generate-delete.gen.sql', sb.toString())

    assertEquals(su.test.normal(sb.toString()), su.test.normal(`
        Teest
    `))
})

Deno.test('Generate select', () => {
    const sb = new StringBuffer()
    const sql = new SQL.SqliteGenerator(sb)
    const table1 = createTable1()
    const table2 = createTable2()
    sql.selects([
        // {
        //     table: table1,
        //     where: {
        //         kind: 'equal',
        //         column: sql.getColumnByName(table1, 'id')!,
        //         value: { kind: 'sql-parameter' }
        //     }
        // },
        {
            table: table2,
            columns: [sql.getColumnByName(table1, 'name')!],
            where: {
                kind: 'and',
                expressions: [
                    {
                        kind: 'equal',
                        column: sql.getColumnByName(table1, 'id')!,
                        value: { kind: 'sql-parameter', name: 'id' }
                    },
                    {
                        kind: 'in',
                        column: sql.getColumnByName(table1, 'name')!,
                        values: [
                            { kind: 'text', value: 'a' },
                            { kind: 'text', value: 'b' }
                        ]
                    },
                    {
                        kind: 'or', expressions: [
                            {
                                kind: 'between',
                                column: sql.getColumnByName(table1, 'id')!,
                                lower: { kind: 'text', value: 'a' },
                                upper: { kind: 'text', value: 'z' }
                            },
                            {
                                kind: 'greater-than',
                                column: sql.getColumnByName(table1, 'id')!,
                                value: { kind: 'numeric', value: 1 }
                            },
                            {
                                kind: 'greater-than',
                                column: sql.getColumnByName(table1, 'id')!,
                                value: { kind: 'ts-interpolation-text', name: 'id' }
                            }
                        ]
                    }
                ]
            },
            joins: [
                {
                    kind: 'join-on',
                    column: sql.getColumnByName(table1, 'id')!,
                    reference: {
                        table: table1,
                        column: sql.getColumnByName(table1, 'id')!,
                    }
                },
                {
                    kind: 'join-using',
                    reference: {
                        table: table1,
                        column: sql.getColumnByName(table1, 'id')!,
                    }
                }
            ],
            orderBy: {
                column: sql.getColumnByName(table1, 'id')!,
                term: SQL.SqlOrderingTerm.Asc
            },
            limit: {
                expression: { kind: 'literal-value', value: { kind: 'numeric', value: 10 } },
                offset: { kind: 'literal-value', value: { kind: 'numeric', value: 100 } }
            }
        },
    ])

    save('./gen-output/generate-select.gen.sql', sb.toString())

    assertEquals(su.test.normal(sb.toString()), su.test.normal(`
        Teest
    `))
})
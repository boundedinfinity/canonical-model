import { assertEquals } from "@std/assert";
import { SqliteGenerator, SqlColumnType, SqlOrderingTerm } from './sqlite-gen-model.ts'

function save(file: string, text: string) {
    const encoder = new TextEncoder();
    Deno.writeFileSync(file, encoder.encode(text))
}

function normal(text: string): string {
    return text
        .trim()
        .split('\n')
        .map(line => line.trim())
        .filter(line => line !== '')
        .join('\n')
        .trim()
}

Deno.test('Generate create table', () => {
    const actual = new SqliteGenerator().tables([
        {
            name: 'name_prefix',
            colunms: [
                { column: 'id', type: SqlColumnType.Text },
                { column: 'name', type: SqlColumnType.Text },
                { column: 'description', type: SqlColumnType.Text }
            ],
            constraints: [
                { kind: 'primary-key', columns: ['id'] },
                { kind: 'unique-key', columns: ['name'] }
            ],

        },
        {
            name: 'name_prefix__abbr',
            colunms: [
                { column: 'name_prefix__id', type: SqlColumnType.Text },
                { column: 'abbr', type: SqlColumnType.Text },
                { column: 'index', type: SqlColumnType.Integer },
            ],
            constraints: [
                {
                    kind: 'foreign-key',
                    column: 'id',
                    reference: {
                        column: 'name_prefix__id',
                        table: 'name_prefix__abbr',
                    }

                },
                { kind: 'unique-key', columns: ['index'] }
            ],
        }
    ])

    save('./gen-output/generate-table.gen.sql', actual)

    assertEquals(normal(actual), normal(`
        Teest
    `))
})

Deno.test('Generate create index', () => {
    const actual = new SqliteGenerator().indexes([
        {
            table: 'name_prefix',
            column: 'name'
        },
        {
            table: 'name_prefix__abbr',
            column: 'abbr',
            unique: true
        },
    ])

    save('./gen-output/generate-index.gen.sql', actual)

    assertEquals(normal(actual), normal(`
        Teest
    `))
})

Deno.test('Generate insert', () => {
    const actual = new SqliteGenerator().inserts([
        {
            table: 'name_prefix',
            values: [
                { column: 'id', value: { kind: 'sql-parameter', name: 'id' } },
                { column: 'name', value: { kind: 'sql-parameter', index: 0 } },
                { column: 'description', value: { kind: 'sql-parameter' } },
                { column: 'arg1', value: { kind: 'text', value: 'arg1' } },
                { column: 'arg2', value: { kind: 'numeric', value: 42 } },
            ]
        },
        {
            table: 'name_prefix__abbr',
            values: [
                { column: 'name_prefix__id', value: { kind: 'sql-parameter', name: 'id' } },
                { column: 'abbr', value: { kind: 'text', value: 'Mr.' } },
                { column: 'index', value: { kind: 'integer', value: 0 } },
            ]
        },

    ])

    save('./gen-output/generate-insert.gen.sql', actual)

    assertEquals(normal(actual), normal(`
        Teest
    `))
})

Deno.test('Generate update', () => {
    const actual = new SqliteGenerator().updates([
        {
            table: 'name_prefix',
            values: [
                { column: 'name', value: { kind: 'sql-parameter', index: 0 } },
                { column: 'description', value: { kind: 'sql-parameter' } },
            ],
            where: {
                kind: 'equal',
                column: 'id',
                value: { kind: 'sql-parameter' }
            }
        },

    ])

    save('./gen-output/generate-update.gen.sql', actual)

    assertEquals(normal(actual), normal(`
        Teest
    `))
})

Deno.test('Generate delete', () => {
    const actual = new SqliteGenerator().deletes([
        {
            table: 'name_prefix',
            where: {
                kind: 'equal',
                column: 'id',
                value: { kind: 'sql-parameter' }
            }
        },

    ])

    save('./gen-output/generate-delete.gen.sql', actual)

    assertEquals(normal(actual), normal(`
        Teest
    `))
})

Deno.test('Generate select', () => {
    const actual = new SqliteGenerator().selects([
        {
            table: 'name_prefix',
            columns: [
                'id'
            ],
            where: {
                kind: 'equal',
                column: 'id',
                value: { kind: 'sql-parameter' }
            }
        },
        {
            table: 'name_prefix',
            columns: [
                'id'
            ],
            where: {
                kind: 'and',
                expressions: [
                    { kind: 'equal', column: 'id', value: { kind: 'sql-parameter', name: 'id' } },
                    {
                        kind: 'in', column: 'name', values: [
                            { kind: 'text', value: 'a' },
                            { kind: 'text', value: 'b' }
                        ]
                    },
                    {
                        kind: 'or', expressions: [
                            {
                                kind: 'between',
                                column: 'id',
                                lower: { kind: 'text', value: 'a' },
                                upper: { kind: 'text', value: 'z' }
                            },
                            {
                                kind: 'greater-than',
                                column: 'id',
                                value: { kind: 'numeric', value: 1 }
                            },
                            {
                                kind: 'greater-than',
                                column: 'id',
                                value: { kind: 'ts-interpolation-text', name: 'id' }
                            }
                        ]
                    }
                ]
            },
            joins: [
                {
                    kind: 'join-on',
                    column: 'id',
                    reference: { table: 'name_prefix__abbr', column: 'name_prefix__id' }
                },
                {
                    kind: 'join-using',
                    reference: { table: 'table', column: 'column' }
                }
            ],
            orderBy: { column: 'id', term: SqlOrderingTerm.Asc },
            limit: {
                expression: { kind: 'literal-value', value: { kind: 'numeric', value: 10 } },
                offset: { kind: 'literal-value', value: { kind: 'numeric', value: 100 } }
            }
        },
    ])

    save('./gen-output/generate-select.gen.sql', actual)

    assertEquals(normal(actual), normal(`
        Teest
    `))
})
import * as BI from "./bounded-model";
import { BoundedManager } from "./bounded-gen";
import { TsNamer } from "./namer";
import { ce, StringBuffer, type StringBufferOptions } from "./utils";
import * as SQL from "./sqlite-gen";

type BoundedSqlGeneratorOptions = StringBufferOptions;

export class BoundedSqliteGenerator {
    options: BoundedSqlGeneratorOptions;
    bounded: BoundedManager;
    namer = new TsNamer();
    sql: SQL.SqliteGenerator;

    constructor(
        bounded: BoundedManager,
        options?: Partial<BoundedSqlGeneratorOptions>,
    ) {
        this.bounded = bounded;
        this.options = {
            indent: 0,
            tabSize: 4,
            ...options,
        };

        this.sql = new SQL.SqliteGenerator();
    }

    private ensureObject(type: BI.BoundedObject) {
        switch (type.kind) {
            case "object":
                break;
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`);
        }
    }

    emitCreate(sb: StringBuffer, type: BI.BoundedObject) {
        const primaryTable = this.getPrimaryTable(type);
        this.sql.table(sb, primaryTable);
        this.sql.index(sb, { table: primaryTable });

        const childTables = this.getChildTables(type);
        for (const childTable of childTables) {
            this.sql.table(sb, childTable);
            this.sql.index(sb, { table: childTable });
        }
    }

    private getSqlType(type: BI.BoundedKind): SQL.SqlColumnType | undefined {
        let ctype: SQL.SqlColumnType | undefined;

        switch (type.kind) {
            case "string":
                ctype = SQL.SqlColumnType.Text;
                break;
            case "boolean":
            case "integer":
                ctype = SQL.SqlColumnType.Integer;
                break;
            case "number":
                ctype = SQL.SqlColumnType.Numeric;
                break;
            case "array":
                break;
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`);
        }

        return ctype;
    }

    getPrimaryTable(type: BI.BoundedObject): SQL.SqlTable {
        this.ensureObject(type);

        const name = this.getSqlTableName(type);
        const table: SQL.SqlTable = {
            name,
            columns: [],
            constraints: [],
        };

        for (const prop of type.properties) {
            const pname = this.getSqlColumnName(prop);
            const ptype = this.getSqlType(prop);

            if (!ptype) {
                continue;
            }

            const column: SQL.SqlColumn = {
                name: pname,
                type: ptype,
                optional: prop.optional,
                table,
            };

            switch (prop.kind) {
                case "string":
                case "integer":
                case "number":
                    column.indexed = prop.indexed;
                    column.primaryKey = prop.primaryKey;
                    column.unique = prop.unique;
                    column.searchable = prop.searchable;
                    break;
                case "ref":
                case "array":
                    continue;
                default:
                    throw new Error(`invalid type ${JSON.stringify(type)}`);
            }

            table.columns.push(column);
        }

        return table;
    }

    private getChildTables(type: BI.BoundedObject): SQL.SqlTable[] {
        this.ensureObject(type);

        const tables: SQL.SqlTable[] = [];
        const primaryTable = this.getPrimaryTable(type);
        const primaryKeys = this.sql.getPrimaryKeys(primaryTable);
        const primaryTablePrimaryKey = primaryKeys[0]!;

        for (const prop of type.properties) {
            switch (prop.kind) {
                case "string":
                case "integer":
                case "number":
                case "boolean":
                case "ref":
                    continue;
                case "array":
                    switch (prop.items.kind) {
                        case "number":
                        case "integer":
                        case "string":
                            if (prop.items.optional) {
                                throw ce(`array item cannot be optional`, prop);
                            }
                            if (prop.items.unique) {
                                throw ce(`array item cannot be unique`, prop);
                            }
                            break;
                        case "ref":
                            continue;
                        default:
                            throw new Error(`invalid type ${JSON.stringify(type)}`);
                    }
                    break;
                default:
                    throw new Error(`invalid type ${JSON.stringify(type)}`);
            }

            const columnType = this.getSqlType(prop.items);

            if (!columnType) {
                continue;
            }

            const columnName = this.getSqlTableName(prop);
            const childName = `${primaryTable.name}__${columnName}`;

            const childTable: SQL.SqlTable = {
                name: childName,
                columns: [],
                constraints: [],
            };

            const childPrimaryKey: SQL.SqlColumn = {
                name: `${primaryTable.name}__${primaryTablePrimaryKey.name}`,
                type: SQL.SqlColumnType.Text,
                indexed: false,
                primaryKey: true,
                table: childTable,
            };

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
                    name: "index",
                    type: SQL.SqlColumnType.Integer,
                    table: childTable,
                },
            ];

            childTable.columns.push(...columns);

            const constraints: SQL.SqlTableConstraint[] = [
                {
                    kind: "foreign-key",
                    column: childPrimaryKey,
                    reference: primaryTablePrimaryKey,
                },
            ];

            childTable.constraints?.push(...constraints);

            tables.push(childTable);
        }

        return tables;
    }

    emitUpdate(sb: StringBuffer, type: BI.BoundedObject): string {
        this.ensureObject(type);

        const primaryTable = this.getPrimaryTable(type);
        const childTables = this.getChildTables(type);

        for (const table of [primaryTable, ...childTables]) {
            this.sql.update(sb, { table });
        }

        return sb.toString();
    }

    emitInsert(sb: StringBuffer, type: BI.BoundedObject): string {
        this.ensureObject(type);

        const primaryTable = this.getPrimaryTable(type);
        this.sql.insert(sb, { table: primaryTable });

        const childTables = this.getChildTables(type);
        for (const childTable of childTables) {
            this.sql.insert(sb, { table: childTable });
        }

        return sb.toString();
    }

    emitSelect(sb: StringBuffer, type: BI.BoundedObject): string {
        this.ensureObject(type);

        const primaryTable = this.getPrimaryTable(type);
        const primaryTableKey = this.sql.getPrimaryKeys(primaryTable)![0];
        const childTables = this.getChildTables(type);
        const joins: SQL.SqlJoin[] = [];
        const expressions: SQL.SqlExpr[] = [];

        for (const column of primaryTable.columns) {
            if (column.searchable || column.primaryKey) {
                switch (column.type) {
                    case SQL.SqlColumnType.Text:
                        expressions.push({
                            kind: "equal",
                            column: column,
                            value: { kind: "sql-parameter" },
                        });
                        break;
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
                                kind: "equal",
                                column: column,
                                value: { kind: "sql-parameter" },
                            });

                            joins.push({
                                kind: "join-on",
                                column: primaryTableKey,
                                reference: column,
                            });
                            break;
                        default:
                        // TODO
                    }
                }
            }
        }

        const select: SQL.SqlSelect = {
            table: primaryTable,
            joins,
            where: { kind: "or", expressions },
        };

        this.sql.select(sb, select);
        return sb.toString();
    }

    private genJoinTable(sb: StringBuffer, type: BI.BoundedObject) {
        const tname = this.getSqlTableName(type);

        for (const prop of type.properties) {
            switch (prop.kind) {
                case "string":
                case "integer":
                case "number":
                case "boolean":
                case "ref":
                    continue;
                case "array":
                    switch (prop.items.kind) {
                        case "number":
                        case "integer":
                        case "string":
                            {
                                const cname = this.getSqlTableName(prop);
                                const ctype = this.bounded.lookup(prop);
                                const name = `${tname}__${cname}`;
                                const id = `${tname}__id`;
                                const amodel: BI.BoundedObject = {
                                    kind: "object",
                                    name,
                                    properties: [
                                        { kind: "string", name: id },
                                        ctype,
                                        { kind: "number", name: "index" },
                                    ],
                                };

                                this.emitCreate(sb, amodel);
                            }
                            break;
                        case "ref":
                            {
                                const cname = this.getSqlTableName(prop);
                                const ctype = this.bounded.lookup(prop);
                                const name = `${tname}___${cname}`;
                                const jmodel: BI.BoundedObject = {
                                    kind: "object",
                                    properties: [],
                                };
                            }
                            break;
                        default:
                            throw new Error(`invalid type ${JSON.stringify(type)}`);
                    }
                    break;
                default:
                    throw new Error(`invalid type ${JSON.stringify(type)}`);
            }
        }
    }

    private getSqlTableName(type: BI.BoundedKind): string {
        let name = this.bounded.itemName(type);

        if (name === "") {
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`);
        }

        name = this.namer.sqlTable(name);
        return name;
    }

    private getSqlColumnName(type: BI.BoundedKind): string {
        let name = this.bounded.itemName(type);

        if (name === "") {
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`);
        }

        name = this.namer.sqlColumn(name);
        return name;
    }
}

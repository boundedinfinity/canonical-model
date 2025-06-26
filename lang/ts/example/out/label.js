"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.LabelSqlliteRepository = exports.Label = void 0;
class Label {
    constructor(args) {
        this.id = args.id;
        this.name = args.name;
        this.description = args.description;
    }
    toObject() {
        return {
            id: this.id,
            name: this.name,
            description: this.description ?? null,
        };
    }
    validate() {
        const errors = [];
        if (this.id === undefined || this.id === null) {
            errors.push(new Error('id is required'));
        }
        else {
            if (this.id.length < 36) {
                errors.push(new Error('id is less than 36 characters'));
            }
            if (this.id.length > 36) {
                errors.push(new Error('id is greater than 36 characters'));
            }
        }
        if (this.name === undefined || this.name === null) {
            errors.push(new Error('name is required'));
        }
        else {
            if (this.name.length < 2) {
                errors.push(new Error('name is less than 2 characters'));
            }
            if (this.name.length > 50) {
                errors.push(new Error('name is greater than 50 characters'));
            }
        }
        if (this.description) {
            if (this.description.length < 2) {
                errors.push(new Error('description is less than 2 characters'));
            }
            if (this.description.length > 50) {
                errors.push(new Error('description is greater than 500 characters'));
            }
        }
        return errors;
    }
}
exports.Label = Label;
// https://bun.sh/docs/api/sqlite
// https://docs.turso.tech/libsql#encryption-at-rest
const client_1 = require("@libsql/client");
class LabelSqlliteRepository {
    constructor(config) {
        this.config = {
            file: 'file:label.db'
        };
        Object.assign(this.config, config);
        this.db = (0, client_1.createClient)({ url: this.config.file });
    }
    async ensureCreate() {
        const result = await this.db.execute(`
            SELECT * FROM sqlite_master 
            WHERE type = 'table' AND name = 'label';
        `);
        if (result.rows.length == 0) {
            await this.db.execute(`CREATE TABLE label (
                    id TEXT,
                    name TEXT NOT NULL,
                    description TEXT,
                    PRIMARY KEY (id),
                );`);
        }
    }
    async insert(label) {
        await this.db.execute({
            sql: `
                INSERT INTO label (
                    id, name, description
                ) VALUES (
                    :id, :name, :description
                );
            `,
            args: label.toObject()
        });
    }
    async select(options) {
        const found = [];
        let sql = `
            SELECT id, name, description
            FROM label
        `;
        const wheres = [];
        const args = {};
        if (options?.id) {
            wheres.push(`id = :id`);
        }
        if (options?.name?.equals) {
            wheres.push(`name = :name`);
            args['name'] = options.name.equals;
        }
        if (options?.description?.equals) {
            wheres.push(`description = :description`);
            args['description'] = options.description.equals;
        }
        if (options?.name?.contains) {
            wheres.push(`name LIKE :name_pattern`);
            args['name_pattern'] = `%${options.name.contains}%`;
        }
        if (options?.description?.contains) {
            wheres.push(`description LIKE :description_pattern`);
            args['description_pattern'] = `%${options.description.contains}%`;
        }
        if (wheres.length > 0) {
            sql += `
                WHERE ${wheres.join(' OR ')}
            `;
        }
        sql += `;`;
        try {
            const result = await this.db.execute({ sql: sql, args: args });
            for (const row of result.rows) {
                console.log(row);
                found.push(new Label({
                    id: row['id'].toLocaleString(),
                    name: row['name'].toLocaleString(),
                    description: row['description']?.toLocaleString()
                }));
            }
        }
        catch (e) {
            console.log(e);
        }
        return found;
    }
}
exports.LabelSqlliteRepository = LabelSqlliteRepository;
// type LabelJsonOptions = {
//     kind: 'full' | 'id',
//     indent: number,
//     typed: boolean
// }
// export class LabelJson {
//     options: LabelJsonOptions = {
//         kind: 'id',
//         indent: 4,
//         typed: true,
//     }
//     constructor(options?: Partial<LabelJsonOptions>) {
//         Object.assign(this.options, options)
//     }
//     serialize(input: Label, options?: Partial<LabelJsonOptions>): string {
//         const roptions: LabelJsonOptions = { ...this.options, ...options }
//         const obj: object = {}
//         if (roptions.typed)
//             Object.assign(obj, { kind: 'label' })
//         switch (roptions.kind) {
//             case 'id':
//                 Object.assign(obj, { id: input.id })
//                 break
//             case 'full':
//                 Object.assign(obj, {
//                     id: input.id,
//                     name: input.name,
//                     description: input.description
//                 })
//                 break
//         }
//         return JSON.stringify(obj, null, roptions.indent)
//     }
//     deserialize(json: string): Label {
//         const obj = JSON.parse(json)
//         return new Label(obj)
//     }
// }
//# sourceMappingURL=label.js.map
import { createClient } from "@libsql/client";
import { Label } from './label'
import { Name } from './name'
import { NamePrefix } from './name-prefix'
import { NameSuffix } from './name-suffix'

import {
  ColumnType,
  Generated,
  Insertable,
  JSONColumnType,
  Selectable,
  Updateable,
} from 'kysely'

interface Database {
    person: Label_Table
    namePrefix: NamePrefix_Table
    nameSuffix: NameSuffix_Table
    name: Name_Table
}

interface Label_Table {
    id: string
    name: string
    description: string | null
}

interface NamePrefix_Table {
    id: string
    name: string
    description?: string | null
}

interface NamePrefix__Abbr__Table {
    id: string
    index: number
}

interface NameSuffix_Table {
    id: string
    name: string
    description?: string | null
}

interface Name_Table {
    id: string
}

// type SqlInput = Array<{ sql: string, args?: { [name: string]: string | number | null } }>
// type ManagedTypes = Label | NamePrefix | NameSuffix | Name

// /////////////////////////////////////////////////////////////////////// 
// // Label
// ///////////////////////////////////////////////////////////////////////

// type LabelSelectArgs = {
//     id?: string,
//     name?: {
//         equals?: string,
//         contains?: string
//     },
//     description?: {
//         equals?: string,
//         contains?: string
//     },
//     page?: {
//         limit?: number,
//         offset?: number
//     }
// }

// type LabelReferecnced = {
//     names?: Name[]
//     namePrefixes?: NamePrefix[]
//     nameSuffixes?: NameSuffix[]
// }

// /////////////////////////////////////////////////////////////////////// 
// // Name
// ///////////////////////////////////////////////////////////////////////

// type NameSelectArgs = {
//     id?: string,
//     prefix?: NamePrefixSelectArgs,
//     first?: {
//         equals?: string,
//         contains?: string
//     },
//     middle?: {
//         equals?: string,
//         contains?: string
//     },
//     last?: {
//         equals?: string,
//         contains?: string
//     },
//     suffix?: NameSuffixSelectArgs,
//     label?: LabelSelectArgs,
//     page?: {
//         limit?: number,
//         offset?: number
//     }
// }

// /////////////////////////////////////////////////////////////////////// 
// // Name Prefix
// ///////////////////////////////////////////////////////////////////////

// type NamePrefixSelectArgs = {
//     id?: string,
//     name?: {
//         equals?: string,
//         contains?: string
//     },
//     description?: {
//         equals?: string,
//         contains?: string
//     },
//     abbr?: {
//         equals?: string,
//         contains?: string
//     },
//     label?: LabelSelectArgs,
//     page?: {
//         limit?: number,
//         offset?: number
//     }
// }

// /////////////////////////////////////////////////////////////////////// 
// // Name Suffix
// ///////////////////////////////////////////////////////////////////////

// type NameSuffixSelectArgs = {
//     id?: string,
//     name?: {
//         equals?: string,
//         contains?: string
//     },
//     description?: {
//         equals?: string,
//         contains?: string
//     },
//     abbr?: {
//         equals?: string,
//         contains?: string
//     },
//     label?: LabelSelectArgs,
//     page?: {
//         limit?: number,
//         offset?: number
//     }
// }

// /////////////////////////////////////////////////////////////////////// 
// // Manager
// ///////////////////////////////////////////////////////////////////////

// export class Manager {
//     db: ReturnType<typeof createClient>
//     config: {} = {}
//     repo: Repository

//     constructor(db: ReturnType<typeof createClient>, config?: Partial<typeof this.config>) {
//         this.db = db
//         Object.assign(this.config, config)
//         this.repo = new Repository(config)
//     }

//     async create() {
//         const batch: SqlInput = []

//         batch.push(...this.repo.create.label())
//         batch.push(...this.repo.create.name())
//         batch.push(...this.repo.create.namePrefix())
//         batch.push(...this.repo.create.nameSuffix())

//         for (const item of batch) {
//             try {
//                 await this.db.execute(item)
//             } catch (e) {
//                 console.log(e)
//                 console.log(item.sql)
//             }
//         }
//     }

//     async save(type: ManagedTypes) {
//         try {
//             switch (type.constructor) {
//                 case Label:
//                     await this._save.label(type as Label)
//                     break
//                 case Name:
//                     this._save.name(type as Name)
//                     break
//                 case NamePrefix:
//                     this._save.namePrefix(type as NamePrefix)
//                     break
//                 case NameSuffix:
//                     this._save.nameSuffix(type as NamePrefix)
//                     break
//             }
//         } catch (e) {
//             console.log(e)
//         }
//     }

//     async labelReferenced(labels: Label[]): Promise<LabelReferecnced> {
//         const found: LabelReferecnced = {}
//         const batch = this.repo.select.labelReferecnced(labels)

//         const [names, namePrefixes, nameSuffixes] = await Promise.all([
//             this.db.batch(batch.name),
//             this.db.batch(batch.namePrefix),
//             this.db.batch(batch.nameSuffix),
//         ])

//         for(const result of names) {
            
//         }

//         return found
//     }

//     private _save = {
//         label: async (obj: Label) => {
//             const select = this.repo.select.label({ id: obj.id })[0]
//             const found = await this.db.execute(select)
//             let batch: SqlInput = []

//             switch (found.rows.length) {
//                 case 0:
//                     batch.push(...this.repo.insert.label(obj))
//                     break
//                 case 1:
//                     batch.push(...this.repo.update.label(obj))
//                     break
//                 default:
//                 // TODO
//             }

//             for (const item of batch) {
//                 await this.db.execute(item)
//             }
//         },

//         name: async (obj: Name) => {
//             const select = await this.repo.select.name({ id: obj.id })[0]
//             const found = await this.db.execute(select)
//             let batch: SqlInput = []

//             switch (found.rows.length) {
//                 case 0:
//                     batch.push(...this.repo.insert.name(obj))
//                     break
//                 case 1:
//                     break
//                 default:
//                 // TODO
//             }

//             batch.push(...this.repo.delete.nameNamePrefix(obj))
//             if (obj.prefix) {
//                 batch.push(...this.repo.insert.nameNamePrefix(obj))
//             }

//             batch.push(...this.repo.delete.nameFirst(obj))
//             if (obj.first) {
//                 batch.push(...this.repo.insert.nameFirst(obj))
//             }

//             batch.push(...this.repo.delete.nameMiddle(obj))
//             if (obj.middle) {
//                 batch.push(...this.repo.insert.nameMiddle(obj))
//             }

//             batch.push(...this.repo.delete.nameLast(obj))
//             if (obj.middle) {
//                 batch.push(...this.repo.insert.nameLast(obj))
//             }

//             batch.push(...this.repo.delete.nameNameSuffix(obj))
//             if (obj.suffix) {
//                 batch.push(...this.repo.insert.nameNameSuffix(obj))
//             }

//             batch.push(...this.repo.delete.nameLabel(obj))
//             if (obj.labels) {
//                 batch.push(...this.repo.insert.nameLabel(obj))
//             }

//             await this.db.batch(batch, 'write')
//         },

//         namePrefix: async (obj: NamePrefix) => {
//             const select = await this.repo.select.namePrefix({ id: obj.id })[0]
//             const found = await this.db.execute(select)
//             let batch: SqlInput = []

//             switch (found.rows.length) {
//                 case 0:
//                     batch.push(...this.repo.insert.namePrefix(obj))
//                     break
//                 case 1:
//                     batch.push(...this.repo.update.namePrefix(obj))
//                     break
//                 default:
//                 // TODO
//             }

//             batch.push(...this.repo.delete.namePrefixAbbr(obj))
//             if (obj.abbr)
//                 batch.push(...this.repo.insert.namePrefixAbbr(obj))

//             await this.db.batch(batch, 'write')
//         },

//         nameSuffix: async (obj: NameSuffix) => {
//             const select = await this.repo.select.nameSuffix({ id: obj.id })[0]
//             const found = await this.db.execute(select)
//             let batch: SqlInput = []

//             switch (found.rows.length) {
//                 case 0:
//                     batch.push(...this.repo.insert.nameSuffix(obj))
//                     break
//                 case 1:
//                     batch.push(...this.repo.update.nameSuffix(obj))
//                     break
//                 default:
//                 // TODO
//             }

//             batch.push(...this.repo.delete.nameSuffixAbbr(obj))

//             if (obj.abbr)
//                 batch.push(...this.repo.insert.nameSuffixAbbr(obj))

//             await this.db.batch(batch, 'write')
//         }
//     }
// }

// /////////////////////////////////////////////////////////////////////// 
// // Repository
// ///////////////////////////////////////////////////////////////////////

// class Repository {
//     config: {} = {}

//     constructor(config?: Partial<typeof this.config>) {
//         Object.assign(this.config, config)
//     }

//     create = {
//         label: (): SqlInput => {
//             const label = `
//                 CREATE TABLE IF NOT EXISTS label (
//                         "id" TEXT,
//                         "name" TEXT NOT NULL,
//                         "description" TEXT,
//                         PRIMARY KEY ("id")
//                 );
//             `

//             return [
//                 {
//                     sql: `
//                 CREATE TABLE IF NOT EXISTS label (
//                         "id" TEXT,
//                         "name" TEXT NOT NULL,
//                         "description" TEXT,
//                         PRIMARY KEY ("id")
//                 );
//             `},
//                 { sql: `CREATE INDEX IF NOT EXISTS "idx_label__name" ON label ("name");` },
//                 { sql: `CREATE INDEX IF NOT EXISTS idx_label__description ON label ("description");` }
//             ]
//         },

//         namePrefix: (): SqlInput => {
//             const namePrefix = `
//                 CREATE TABLE IF NOT EXISTS name_prefix (
//                         "id" TEXT,
//                         "name" TEXT NOT NULL,
//                         "description" TEXT,
//                         PRIMARY KEY ("id")
//                 );
//                 CREATE INDEX IF NOT EXISTS idx_name_prefix__name ON name_prefix ("name");
//                 CREATE INDEX IF NOT EXISTS idx_name_prefix__description ON name_prefix ("description");
//             `

//             const namePrefixAbbr = `
//                 CREATE TABLE IF NOT EXISTS name_prefix__abbr (
//                     "name_prefix__id" TEXT,
//                     "index" INTEGER,
//                     "abbr" TEXT,
//                     PRIMARY KEY ("name_prefix__id", "index"),
//                     FOREIGN KEY ("name_prefix__id") REFERENCES name_prefix("id")
//                 );
//                 CREATE INDEX IF NOT EXISTS idx_name_prefix__abbr__abbr ON name_prefix__abbr ("abbr");
//             `

//             const namePrefixLabel = `
//                 CREATE TABLE IF NOT EXISTS name_prefix__label (
//                     "name_prefix__id" TEXT,
//                     "label__id" TEXT,
//                     "index" INTEGER,
//                     PRIMARY KEY ("name_prefix__id", "label__id", "index"),
//                     FOREIGN KEY ("name_prefix__id") REFERENCES name_prefix("id"),
//                     FOREIGN KEY ("label__id") REFERENCES label("id")
//                 );
//             `

//             return [{ sql: namePrefix }, { sql: namePrefixAbbr }, { sql: namePrefixLabel }]
//         },

//         nameSuffix: (): SqlInput => {
//             const nameSuffix = `
//                 CREATE TABLE IF NOT EXISTS name_suffix (
//                         "id" TEXT,
//                         "name" TEXT NOT NULL,
//                         "description" TEXT,
//                         PRIMARY KEY ("id")
//                 );
//                 CREATE INDEX IF NOT EXISTS "idx_name_suffix__name" ON name_suffix ("name");
//                 CREATE INDEX IF NOT EXISTS "idx_name_suffix__description" ON name_suffix ("description");
//             `

//             const nameSuffixAbbr = `
//                 CREATE TABLE IF NOT EXISTS name_suffix__abbr (
//                     "name_suffix__id" TEXT,
//                     "index" INTEGER,
//                     "abbr" TEXT,
//                     PRIMARY KEY ("name_suffix__id", "index"),
//                     FOREIGN KEY ("name_suffix__id") REFERENCES name_suffix("id")
//                 );
//                 CREATE INDEX IF NOT EXISTS idx_name_suffix__abbr__abbr ON name_suffix__abbr (abbr);
//             `

//             const nameSuffixLabel = `
//                 CREATE TABLE IF NOT EXISTS name_suffix__label (
//                     "name_suffix__id" TEXT,
//                     "label__id" TEXT,
//                     "index" INTEGER,
//                     PRIMARY KEY ("name_suffix__id", "label__id", "index"),
//                     FOREIGN KEY ("name_suffix__id") REFERENCES name_suffix("id"),
//                     FOREIGN KEY ("label__id") REFERENCES label("id")
//                 );
//             `

//             return [{ sql: nameSuffix }, { sql: nameSuffixAbbr }, { sql: nameSuffixLabel }]
//         },

//         name: (): SqlInput => {
//             const name = `
//                 CREATE TABLE IF NOT EXISTS name (
//                     "id" TEXT,
//                     PRIMARY KEY ("id")
//                 );
//             `

//             const namePrefix = `
//                 CREATE TABLE IF NOT EXISTS name__name_prefix (
//                     "name__id" TEXT,
//                     "name_prefix__id" TEXT,
//                     "index" INTEGER,
//                     PRIMARY KEY ("name__id", "name_prefix__id", "index"),
//                     FOREIGN KEY ("name__id") REFERENCES name("id"),
//                     FOREIGN KEY ("name_prefix__id") REFERENCES name_prefix("id")
//                 );
//             `

//             const nameFirst = `
//                 CREATE TABLE IF NOT EXISTS name__first (
//                     "name__id" TEXT,
//                     "index" INTEGER,
//                     "first" TEXT,
//                     PRIMARY KEY ("name__id", "index"),
//                     FOREIGN KEY ("name__id") REFERENCES name("id")
//                 );                
//             `

//             const nameFirstIdx = `
//                 CREATE INDEX IF NOT EXISTS idx_name__first__first ON name__first (first);
//             `

//             const nameMiddle = `
//                 CREATE TABLE IF NOT EXISTS name__middle (
//                     "name__id" TEXT,
//                     "index" INTEGER,
//                     "middle" TEXT,
//                     PRIMARY KEY ("name__id", "index"),
//                     FOREIGN KEY ("name__id") REFERENCES name("id")
//                 );
//             `

//             const nameMiddleIdx = `
//                 CREATE INDEX IF NOT EXISTS idx_name__middle__middle ON name__middle ("middle");
//             `

//             const nameLast = `
//                 CREATE TABLE IF NOT EXISTS name__last (
//                     "name__id" TEXT,
//                     "index" INTEGER,
//                     "last" TEXT,
//                     PRIMARY KEY ("name__id", "index"),
//                     FOREIGN KEY ("name__id") REFERENCES name("id")
//                 );
//             `

//             const nameLastIdx = `
//                 CREATE INDEX IF NOT EXISTS idx_name__last__last ON name__last (last);
//             `

//             const nameSuffix = `
//                 CREATE TABLE IF NOT EXISTS name__name_suffix (
//                     "name__id" TEXT,
//                     "name_suffix__id" TEXT,
//                     "index" INTEGER,
//                     PRIMARY KEY ("name__id", "name_suffix__id", "index"),
//                     FOREIGN KEY ("name__id") REFERENCES name("id"),
//                     FOREIGN KEY ("name_suffix__id") REFERENCES name_suffix("id")
//                 );
//             `

//             const nameLabel = `
//                 CREATE TABLE IF NOT EXISTS name__label (
//                     "name__id" TEXT,
//                     "label__id" TEXT,
//                     "index" INTEGER,
//                     PRIMARY KEY ("name__id", "label__id", "index"),
//                     FOREIGN KEY ("name__id") REFERENCES name("id"),
//                     FOREIGN KEY ("label__id") REFERENCES label("id")
//                 );
//             `

//             return [
//                 { sql: name },
//                 { sql: nameFirst }, { sql: nameFirstIdx },
//                 { sql: nameMiddle }, { sql: nameMiddleIdx },
//                 { sql: nameLast }, { sql: nameLastIdx },
//                 { sql: namePrefix }, { sql: nameSuffix },
//                 { sql: nameLabel }
//             ]
//         }
//     }

//     insert = {
//         label: (obj: Label): SqlInput => {
//             const sql = `
//                 INSERT INTO label (
//                     "id", "name", "description"
//                 ) VALUES (
//                     :id, :name, :description
//                 );
//             `
//             const args = {
//                 id: obj.id,
//                 name: obj.name,
//                 description: obj.description ?? null
//             }

//             return [{ sql, args }]
//         },

//         name: (obj: Name): SqlInput => {
//             return [{
//                 sql: `INSERT INTO name ("id") VALUES (:id);`,
//                 args: {
//                     id: obj.id,
//                 }
//             }]
//         },

//         nameNamePrefix: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             if (obj.prefix) {
//                 for (let i = 0; i < obj.prefix.length; i++) {
//                     batch.push({
//                         sql: `
//                             INSERT INTO name__name_prefix (
//                                 "name__id", "name_prefix__id", "index"
//                             ) VALUES (
//                                 :name__id, :name_prefix__id, :index
//                             )
//                         `,
//                         args: {
//                             name__id: obj.id,
//                             name_prefix__id: obj.prefix[i].id,
//                             index: i
//                         }
//                     })
//                 }
//             }

//             return batch
//         },

//         nameFirst: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             const sql = `
//                 INSERT INTO name__first (
//                     "name__id", "index", "first"
//                 ) VALUES (
//                     :name_id, :index, :first
//                 )
//             `

//             if (obj.first) {
//                 for (let i = 0; i < obj.first.length; i++) {
//                     const args = {
//                         name_first__id: obj.id,
//                         index: i,
//                         first: obj.first[i],
//                     }

//                     batch.push({ sql, args })
//                 }
//             }

//             return batch
//         },

//         nameMiddle: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             const sql = `
//                 INSERT INTO name__middle (
//                     "name__id", "index", "middle"
//                 ) VALUES (
//                     :name_id, :index, :middle
//                 )
//             `

//             if (obj.middle) {
//                 for (let i = 0; i < obj.middle.length; i++) {
//                     const args = {
//                         name_middle__id: obj.id,
//                         index: i,
//                         middle: obj.middle[i],
//                     }

//                     batch.push({ sql, args })
//                 }
//             }

//             return batch
//         },

//         nameLast: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             const sql = `
//                 INSERT INTO name__last (
//                     "name__id", "index", "last"
//                 ) VALUES (
//                     :name_id, :index, :last
//                 )
//             `

//             if (obj.last) {
//                 for (let i = 0; i < obj.last.length; i++) {
//                     const args = {
//                         name_last__id: obj.id,
//                         index: i,
//                         last: obj.last[i],
//                     }

//                     batch.push({ sql, args })
//                 }
//             }

//             return batch
//         },

//         nameNameSuffix: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             if (obj.suffix) {
//                 for (let i = 0; i < obj.suffix.length; i++) {
//                     batch.push({
//                         sql: `
//                             INSERT INTO name__name_suffix (
//                                 "name__id", name_suffix__id, "index"
//                             ) VALUES (
//                                 :name__id, :name_suffix__id, :index
//                             )
//                         `,
//                         args: {
//                             name__id: obj.id,
//                             name_suffix__id: obj.suffix[i].id,
//                             index: i
//                         }
//                     })
//                 }
//             }

//             return batch
//         },

//         nameLabel: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             if (obj.labels) {
//                 for (let i = 0; i < obj.labels.length; i++) {
//                     batch.push({
//                         sql: `
//                             INSERT INTO name__label (
//                                 "name__id", "label__id", "index"
//                             ) VALUES (
//                                 :name__id, :label__id, :index
//                             )
//                         `,
//                         args: {
//                             name__id: obj.id,
//                             label__id: obj.labels[i].id,
//                             index: i
//                         }
//                     })
//                 }
//             }

//             return batch
//         },

//         namePrefix: (namePrefix: NamePrefix): SqlInput => {
//             return [{
//                 sql: `INSERT INTO name_prefix ("id", "name", "description") VALUES (:id, :name, :description);`,
//                 args: {
//                     id: namePrefix.id,
//                     name: namePrefix.name,
//                     description: namePrefix.description || null
//                 }
//             }]
//         },

//         namePrefixAbbr: (namePrefix: NamePrefix): SqlInput => {
//             const batch: SqlInput = []
//             const namePrefixAbbrArgs = {
//                 name_prefix__id: namePrefix.id
//             }

//             if (namePrefix.abbr) {
//                 for (let i = 0; i < namePrefix.abbr.length; i++) {
//                     const namePrefixAbbrArgs = {
//                         name_prefix__id: namePrefix.id,
//                         index: i,
//                         abbr: namePrefix.abbr[i]
//                     }

//                     batch.push({
//                         sql: `
//                                 INSERT INTO name_prefix__abbr (
//                                     "name_prefix__id", "index", abbr
//                                 ) VALUES (
//                                     :name_prefix__id, :index, :abbr
//                                 )
//                             `,
//                         args: namePrefixAbbrArgs
//                     })
//                 }
//             }

//             return batch
//         },

//         namePrefixLabel: (obj: NamePrefix): SqlInput => {
//             const batch: SqlInput = []

//             if (obj.labels) {
//                 for (let i = 0; i < obj.labels.length; i++) {
//                     batch.push({
//                         sql: `
//                             INSERT INTO name_prefix__label (
//                                 "name_prefix__id", "label__id", "index"
//                             ) VALUES (
//                                 :name_prefix__id, :label__id, :index
//                             )
//                         `,
//                         args: {
//                             name_prefix__id: obj.id,
//                             label__id: obj.labels[i].id,
//                             index: i
//                         }
//                     })
//                 }
//             }

//             return batch
//         },

//         nameSuffix: (nameSuffix: NamePrefix): SqlInput => {
//             const nameSuffixArgs = {
//                 id: nameSuffix.id,
//                 name: nameSuffix.name,
//                 description: nameSuffix.description || null
//             }

//             return [{
//                 sql: `
//                             INSERT INTO name_suffix (
//                                 "id", "name", "description"
//                             ) VALUES (
//                                 :id, :name, :description
//                             );
//                         `,
//                 args: nameSuffixArgs
//             }]
//         },

//         nameSuffixAbbr: (nameSuffix: NameSuffix): SqlInput => {
//             const nameSuffixAbbrArgs = {
//                 name_suffix__id: nameSuffix.id
//             }

//             const batch: SqlInput = [{
//                 sql: `
//                     DELETE FROM name_suffix__abbr
//                     WHERE "name_suffix__id" = :name_suffix__id
//                 `,
//                 args: nameSuffixAbbrArgs
//             }]

//             if (nameSuffix.abbr) {
//                 for (let i = 0; i < nameSuffix.abbr.length; i++) {
//                     const nameSuffixAbbrArgs = {
//                         name_suffix__id: nameSuffix.id,
//                         index: i,
//                         abbr: nameSuffix.abbr[i]
//                     }

//                     batch.push({
//                         sql: `
//                                 INSERT INTO name_suffix__abbr (
//                                     "name_suffix__id", "index", "abbr"
//                                 ) VALUES (
//                                     :name_suffix__id, :index, :abbr
//                                 )
//                             `,
//                         args: nameSuffixAbbrArgs
//                     })
//                 }
//             }

//             return batch
//         },

//         nameSuffixLabel: (obj: NameSuffix): SqlInput => {
//             const batch: SqlInput = []

//             if (obj.labels) {
//                 for (let i = 0; i < obj.labels.length; i++) {
//                     batch.push({
//                         sql: `
//                             INSERT INTO name_suffix__label (
//                                 "name_suffix__id", "label__id", "index"
//                             ) VALUES (
//                                 :name_suffix__id, :label__id, :index
//                             )
//                         `,
//                         args: {
//                             name_suffix__id: obj.id,
//                             label__id: obj.labels[i].id,
//                             index: i
//                         }
//                     })
//                 }
//             }

//             return batch
//         },
//     }

//     update = {
//         label: (obj: Label): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `
//                     UPDATE label SET
//                         "name" = :name,
//                         "description" = :description
//                     WHERE
//                         "id" = :id
//                     ;
//                 `,
//                 args: {
//                     id: obj.id,
//                     name: obj.name,
//                     description: obj.description ?? null
//                 }
//             })

//             return batch
//         },

//         namePrefix: (namePrefix: NamePrefix): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `
//                     UPDATE name_prefix SET
//                         "name" = :name, 
//                         "description" = :description
//                     WHERE "id" = :id;                 
//                 `,
//                 args: {
//                     id: namePrefix.id,
//                     name: namePrefix.name,
//                     description: namePrefix.description || null
//                 }
//             })

//             return batch
//         },

//         nameSuffix: (nameSuffix: NamePrefix): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `
//                     UPDATE name_suffix SET
//                         "name" = :name, 
//                         "description" = :description
//                     WHERE "id" = :id;                 
//                 `,
//                 args: {
//                     id: nameSuffix.id,
//                     name: nameSuffix.name,
//                     description: nameSuffix.description || null
//                 }
//             })

//             return batch
//         }
//     }

//     select = {
//         label: (options?: LabelSelectArgs): SqlInput => {
//             let sql = `
//                 SELECT "id", "name", "description"
//                 FROM label
//             `
//             const wheres: string[] = []
//             const args: { [name: string]: string | number } = {}

//             if (options?.id) {
//                 wheres.push(`"id" = :id`)
//                 args['id'] = options.id
//             }

//             if (options?.name?.equals) {
//                 wheres.push(`"name" = :name`)
//                 args['name'] = options.name.equals
//             }

//             if (options?.description?.equals) {
//                 wheres.push(`"description" = :description`)
//                 args['description'] = options.description.equals
//             }

//             if (options?.name?.contains) {
//                 wheres.push(`"name" LIKE :name_pattern`)
//                 args['name_pattern'] = `%${options.name.contains}%`
//             }

//             if (options?.description?.contains) {
//                 wheres.push(`"description" LIKE :description_pattern`)
//                 args['description_pattern'] = `%${options.description.contains}%`
//             }

//             if (wheres.length > 0) {
//                 sql += `
//                     WHERE ${wheres.join(' OR ')}
//                 `
//             }

//             if (options?.page) {
//                 args['limit'] = options.page.limit ?? 10
//                 args['offset'] = options.page.offset ?? 0

//                 sql += `
//                     LIMIT :limit OFFSET :offset
//                 `
//             }

//             sql += `;`

//             return [{ sql, args }]
//         },

//         labelReferecnced(labels: Label[]): {
//             name: SqlInput,
//             namePrefix: SqlInput,
//             nameSuffix: SqlInput
//         } {
//             const found: {
//                 name: SqlInput,
//                 namePrefix: SqlInput,
//                 nameSuffix: SqlInput
//             } = {
//                 name: [],
//                 namePrefix: [],
//                 nameSuffix: []
//             }

//             for (const label of labels) {
//                 found.name.push(...this.name({ label: { id: label.id } }))
//                 found.namePrefix.push(...this.namePrefix({ label: { id: label.id } }))
//                 found.nameSuffix.push(...this.nameSuffix({ label: { id: label.id } }))
//             }

//             return found
//         },

//         name: (options?: NameSelectArgs): SqlInput => {
//             const batch: Array<SqlInput> = []

//             let sql = `
//                 SELECT name."id"
//                 FROM name
//             `
//             let wheres: string[] = []
//             let joins: string[] = []
//             const args: { [name: string]: string | number } = {}

//             if (options?.id) {
//                 wheres.push(`name."id" = :id`)
//                 args['id'] = options.id
//             }

//             if (options?.first?.equals) {
//                 wheres.push(`name__first."name" = :first`)
//                 joins.push(`INNER JOIN name__first ON name."id" = name__first."name__id"`)
//                 args['first'] = options.first.equals
//             }

//             if (options?.middle?.equals) {
//                 wheres.push(`name__middle."name" = :middle`)
//                 joins.push(`INNER JOIN name__middle ON name."id" = name__middle."name__id"`)
//                 args['middle'] = options.middle.equals
//             }

//             if (options?.last?.equals) {
//                 wheres.push(`name__last."name" = :last`)
//                 joins.push(`INNER JOIN name__last ON name."id" = name_last."name__id"`)
//                 args['last'] = options.last.equals
//             }

//             if (options?.label?.name?.equals) {
//                 wheres.push(`label."name" = :label_name`)
//                 joins.push(`INNER JOIN name__label ON name__label."name__id" = name."id"`)
//                 joins.push(`INNER JOIN name__label ON name__label."label__id" = label."id"`)
//                 args['label_name'] = options.label.name.equals
//             }

//             if (options?.first?.contains) {
//                 wheres.push(`name__first."name" LIKE :first_pattern`)
//                 joins.push(`INNER JOIN name__first ON name."id" = name__first."name__id"`)
//                 args['first_pattern'] = `%${options.first.contains}%`
//             }

//             if (options?.middle?.contains) {
//                 wheres.push(`name__middle."name" LIKE :middle_pattern`)
//                 joins.push(`INNER JOIN name__middle ON name."id" = name__middle."name__id"`)
//                 args['middle_pattern'] = `%${options.middle.contains}%`
//             }

//             if (options?.last?.contains) {
//                 wheres.push(`name__last."name" LIKE :last_pattern`)
//                 joins.push(`INNER JOIN name__last ON name."id" = name__last."name__id"`)
//                 args['last_pattern'] = `%${options.last.contains}%`
//             }

//             if (options?.label?.name?.contains) {
//                 wheres.push(`label."name" LIKE :label_pattern`)
//                 joins.push(`INNER JOIN name__label ON name__label."name__id" = name.id`)
//                 joins.push(`INNER JOIN name__label ON name__label."label__id" = label.id`)
//                 args['label_pattern'] = options.label.name.contains
//             }

//             wheres = [...new Set(wheres)]
//             joins = [...new Set(joins)]

//             if (wheres.length > 0) {
//                 sql += `
//                     WHERE ${wheres.join(' OR\n')}
//                 `
//             }

//             if (joins.length > 0) {
//                 sql += joins.join('\n')
//             }

//             if (options?.page) {
//                 args['limit'] = options.page.limit ?? 10
//                 args['offset'] = options.page.offset ?? 0

//                 sql += `
//                     LIMIT :limit OFFSET :offset
//                 `
//             }

//             sql += `;`

//             return [{ sql, args }]
//         },

//         namePrefix_Label: (option?: { id: string }): SqlInput => {
//             const batch: SqlInput = []
//             return batch
//         },

//         namePrefix: (options?: NamePrefixSelectArgs): SqlInput => {
//             const batch: SqlInput = []

//             let sql = `
//                 SELECT name_prefix."id", name_prefix."name", name_prefix.description
//                 FROM name_prefix
//             `
//             let wheres: string[] = []
//             let joins: string[] = []
//             const args: { [name: string]: string | number } = {}

//             if (options?.id) {
//                 wheres.push(`name_prefix."id" = :id`)
//                 args['id'] = options.id
//             }

//             if (options?.name?.equals) {
//                 wheres.push(`name_prefix."name" = :name`)
//                 args['name'] = options.name.equals
//             }

//             if (options?.description?.equals) {
//                 wheres.push(`name_prefix."description" = :description`)
//                 args['description'] = options.description.equals
//             }

//             if (options?.abbr?.equals) {
//                 wheres.push(`name_prefix__abbr."abbr" = :abbr`)
//                 joins.push(`INNER JOIN name_prefix__abbr ON name_prefix."id" = name_prefix__abbr."name_prefix__id"`)
//             }


//             if (options?.name?.contains) {
//                 wheres.push(`name_prefix."name" LIKE :name_pattern`)
//                 args['name_pattern'] = `%${options.name.contains}%`
//             }

//             if (options?.description?.contains) {
//                 wheres.push(`name_prefix.description LIKE :description_pattern`)
//                 args['description_pattern'] = `%${options.description.contains}%`
//             }

//             if (options?.abbr?.contains) {
//                 wheres.push(`name_prefix__abbr.abbr LIKE :abbr_pattern`)
//                 joins.push(`INNER JOIN name_prefix__abbr ON name_prefix."id" = name_prefix__abbr."name_prefix__id"`)
//             }

//             wheres = [...new Set(wheres)]
//             joins = [...new Set(joins)]

//             if (wheres.length > 0) {
//                 sql += `
//                     WHERE ${wheres.join(' OR\n')}
//                 `
//             }

//             if (joins.length > 0) {
//                 sql += joins.join('\n')
//             }

//             if (options?.page) {
//                 args['limit'] = options.page.limit ?? 10
//                 args['offset'] = options.page.offset ?? 0

//                 sql += `
//                     LIMIT :limit OFFSET :offset
//                 `
//             }

//             sql += `;`

//             return [{ sql, args }]
//         },

//         nameSuffix: (options?: NameSuffixSelectArgs): SqlInput => {
//             const batch: SqlInput = []

//             let sql = `
//                 SELECT name_suffix."id", name_suffix."name", name_suffix."description"
//                 FROM name_suffix
//             `
//             let wheres: string[] = []
//             let joins: string[] = []
//             const args: { [name: string]: string | number } = {}

//             if (options?.id) {
//                 wheres.push(`name_suffix."id" = :id`)
//                 args['id'] = options.id
//             }

//             if (options?.name?.equals) {
//                 wheres.push(`name_suffix."name" = :name`)
//                 args['name'] = options.name.equals
//             }

//             if (options?.description?.equals) {
//                 wheres.push(`name_suffix."description" = :description`)
//                 args['description'] = options.description.equals
//             }

//             if (options?.abbr?.equals) {
//                 wheres.push(`name_suffix__abbr."abbr" = :abbr`)
//                 joins.push(`INNER JOIN name_suffix__abbr ON name_suffix."id" = name_suffix__abbr."name_suffix__id"`)
//             }


//             if (options?.name?.contains) {
//                 wheres.push(`name_suffix."name" LIKE :name_pattern`)
//                 args['name_pattern'] = `%${options.name.contains}%`
//             }

//             if (options?.description?.contains) {
//                 wheres.push(`name_suffix."description" LIKE :description_pattern`)
//                 args['description_pattern'] = `%${options.description.contains}%`
//             }

//             if (options?.abbr?.contains) {
//                 wheres.push(`name_suffix__abbr.abbr LIKE :abbr_pattern`)
//                 joins.push(`INNER JOIN name_suffix__abbr ON name_suffix."id" = name_suffix__abbr."name_suffix__id"`)
//             }

//             wheres = [...new Set(wheres)]
//             joins = [...new Set(joins)]

//             if (wheres.length > 0) {
//                 sql += `
//                     WHERE ${wheres.join(' OR\n')}
//                 `
//             }

//             if (joins.length > 0) {
//                 sql += joins.join('\n')
//             }

//             if (options?.page) {
//                 args['limit'] = options.page.limit ?? 10
//                 args['offset'] = options.page.offset ?? 0

//                 sql += `
//                     LIMIT :limit OFFSET :offset
//                 `
//             }

//             sql += `;`

//             return [{ sql, args }]
//         }
//     }

//     delete = {
//         label: (obj: Label): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `DELETE FROM label WHERE "id" = :id`,
//                 args: { id: obj.id }
//             })

//             return batch
//         },

//         name: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `DELETE FROM name WHERE "id" = :id`,
//                 args: { id: obj.id }
//             })

//             return batch
//         },

//         nameNamePrefix: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `DELETE FROM name__name_prefix WHERE "name__id" = :id`,
//                 args: { id: obj.id }
//             })

//             return batch
//         },

//         nameFirst: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `DELETE FROM name__first WHERE "id" = :id`,
//                 args: { id: obj.id }
//             })

//             return batch
//         },

//         nameMiddle: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `DELETE FROM name__middle WHERE "id" = :id`,
//                 args: { id: obj.id }
//             })

//             return batch
//         },

//         nameLast: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `DELETE FROM name__last WHERE "id" = :id`,
//                 args: { id: obj.id }
//             })

//             return batch
//         },

//         nameNameSuffix: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `DELETE FROM name__name_suffix WHERE "name__id" = :id`,
//                 args: { id: obj.id }
//             })

//             return batch
//         },

//         nameLabel: (obj: Name): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `DELETE FROM name__label WHERE "name__id" = :id`,
//                 args: { id: obj.id }
//             })

//             return batch
//         },

//         namePrefix: (obj: NamePrefix): SqlInput => {
//             return [{
//                 sql: `
//                     DELETE FROM name_prefix
//                     WHERE "id" = :id
//                 `,
//                 args: { id: obj.id }
//             }]
//         },

//         namePrefixAbbr: (obj: NamePrefix): SqlInput => {
//             return [{
//                 sql: `
//                     DELETE FROM name_prefix__abbr
//                     WHERE "name_prefix__id" = :id
//                 `,
//                 args: { id: obj.id }
//             }]
//         },

//         namePrefixLabel: (obj: NamePrefix): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `DELETE FROM name_prefix__label WHERE "name__id" = :id`,
//                 args: { id: obj.id }
//             })

//             return batch
//         },

//         nameSuffix: (obj: NameSuffix): SqlInput => {
//             return [{
//                 sql: `
//                     DELETE FROM name_suffix
//                     WHERE "id" = :id
//                 `,
//                 args: { id: obj.id }
//             }]
//         },

//         nameSuffixAbbr: (obj: NameSuffix): SqlInput => {
//             return [{
//                 sql: `
//                     DELETE FROM name_suffix__abbr
//                     WHERE "name_suffix__id" = :id
//                 `,
//                 args: { id: obj.id }
//             }]
//         },

//         nameSuffixLabel: (obj: NameSuffix): SqlInput => {
//             const batch: SqlInput = []

//             batch.push({
//                 sql: `DELETE FROM name_suffix__label WHERE "name__id" = :id`,
//                 args: { id: obj.id }
//             })

//             return batch
//         },
//     }
// }

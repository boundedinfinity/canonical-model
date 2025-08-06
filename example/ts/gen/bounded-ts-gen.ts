import * as BI from "./bounded-model";
import { BoundedManager } from "./bounded-gen";
import { BoundedSqliteGenerator } from "./bounded-sqlite-gen";
import * as TS from "./ts-gen";
import { TsNamer } from "./namer";
import * as npath from "path";
import { StringBuffer } from "./utils";

type EmitValidationOptions = {
    name?: string;
};

export class BoundedTsGenerator {
    bounded: BoundedManager;
    namer = new TsNamer();
    tsGen = new TS.TsGenerator();
    sqlGen: BoundedSqliteGenerator;

    constructor(bounded: BoundedManager) {
        this.bounded = bounded;
        this.sqlGen = new BoundedSqliteGenerator(bounded);
    }

    private invalidTypeErr(type: BI.BoundedKind): Error {
        return new Error(`invalid token: ${JSON.stringify(type)}`);
    }

    emit(type: BI.BoundedKind): string {
        const sb = new StringBuffer();

        switch (type.kind) {
            case "object":
                for (const imp of this.emitImports(type)) {
                    this.tsGen.emit(sb, imp);
                }

                this.tsGen.emit(sb, {
                    kind: "class-decl",
                    name: this.emitTsClassName(type),
                    modifiers: this.tsGen.Builder.modifiers(["export"]),
                    properties: type.properties.map((prop) => this.emitProperty(prop)),
                    methods: [
                        this.emitConstructor(type),
                        this.emitValidator(type),
                        this.emitSqlCreate(type),
                        this.emitSqlSelect(type),
                        this.emitSqlInsert(type),
                        this.emitSqlUpdate(type),
                    ],
                });

                break;
            default:
                throw this.invalidTypeErr(type);
        }
        return sb.toString();
    }

    private emitProperty(type: BI.BoundedKind): TS.TsProperty {
        let result: TS.TsProperty;

        switch (type.kind) {
            case "string":
            case "number":
            case "boolean":
                result = {
                    kind: "property",
                    name: this.emitTsVarName(type),
                    type: this.tsGen.Builder.type(type.kind),
                    optional: type.optional,
                };
                break;
            case "integer":
                result = {
                    kind: "property",
                    name: this.emitTsVarName(type),
                    type: this.tsGen.Builder.type("number"),
                    optional: type.optional,
                };
                break;
            case "array":
                result = this.emitProperty(type.items);
                result.name = this.emitTsVarName(type);
                result.array = true;
                break;
            case "ref":
                {
                    const ref = this.bounded.lookup(type);
                    result = {
                        kind: "property",
                        name: this.emitTsVarName(type),
                        type: {
                            kind: "type",
                            type: {
                                kind: "raw",
                                value: this.emitTsClassName(ref),
                            },
                        },
                    };
                }
                break;
            default:
                throw this.invalidTypeErr(type);
        }

        return result;
    }

    private emitImports(type: BI.BoundedObject): TS.TsImport[] {
        const imports: TS.TsImport[] = [];

        const process = (prop: BI.BoundedKind) => {
            switch (prop.kind) {
                case "array":
                    process(prop.items);
                    break;
                case "ref":
                    {
                        const ref = this.bounded.lookup(prop);
                        const imp: TS.TsImport = {
                            kind: "import",
                            from: this.emitFilename(ref, "."),
                            symbols: [{
                                kind: "import-symbol",
                                symbol: this.emitTsClassName(ref),
                            }],
                        };

                        imports.push(imp);
                    }
                    break;
            }
        };

        for (const prop of type.properties) {
            process(prop);
        }

        return imports;
    }

    private emitConstructor(type: BI.BoundedObject): TS.TsMethodDecl {
        const sb = new StringBuffer();

        for (const prop of type.properties) {
            const name = this.emitTsVarName(prop);
            sb.l(`this.${name} = args.${name}\n`);
        }

        if (sb.lines.length > 0) {
            sb.nl();
        }

        const ctor: TS.TsMethodDecl = {
            kind: "method-decl",
            name: "constructor",
            args: {
                kind: "arguments",
                sep: ", ",
                items: [
                    {
                        kind: "property",
                        name: "args",
                        type: {
                            kind: "type",
                            type: {
                                kind: "object-type",
                                sep: ",\n",
                                properties: type.properties.map((prop) =>
                                    this.emitProperty(prop)
                                ),
                            },
                        },
                    },
                ],
            },
            body: sb.lines,
        };

        return ctor;
    }

    private emitValidator(type: BI.BoundedObject): TS.TsMethodDecl {
        const sb = new StringBuffer();
        this.emitValidation(sb, type);

        const ctor: TS.TsMethodDecl = {
            kind: "method-decl",
            name: "validate",
            return: "Error[]",
            body: sb.lines,
        };

        return ctor;
    }

    private emitValidation(
        sb: StringBuffer,
        type: BI.BoundedKind,
        options?: Partial<EmitValidationOptions>,
    ) {
        const name = options?.name || `this.${this.emitTsVarName(type)}`;
        const opt = type.optional ? `${name} && ` : "";

        const emitErr = (text: string): string => {
            text = text.replaceAll(`'`, `\\'`);
            const result = `   errors.push(new Error('${name} ${text}'))`;
            return result;
        };

        const required = () => {
            if (!type.optional) {
                sb.l(`if(!${name}) {`);
                sb.i(() => {
                    sb.l(emitErr("is required"));
                });
                sb.l(`}`);
                sb.nl();
            }
        };

        const emitValues = (items: TS.TsPrimitiveLiterals[]): string => {
            const values = this.tsGen.Builder.list(items);
            const sb2 = new StringBuffer();

            for (let i = 0; i < values.length; i++) {
                const value = values[i];
                this.tsGen.emit(sb2, value);

                if (i < values.length - 1) {
                    sb2.a(", ");
                }
            }

            return sb2.toString();
        };

        switch (type.kind) {
            case "object":
                sb.l(`const errors: Error[] = []`);
                sb.nl();

                for (const prop of type.properties) {
                    this.emitValidation(sb, prop);
                }

                sb.nl();
                sb.l(`return errors`);
                break;
            case "string":
                required();

                if (type.min) {
                    sb.l(`if(${opt}${name}.length < ${type.min}) {`);
                    sb.i(() => {
                        sb.l(emitErr(`length is less than ${type.min}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }

                if (type.max) {
                    sb.l(`if(${opt}${name}.length > ${type.max}) {`);
                    sb.i(() => {
                        sb.l(emitErr(`length is greater than ${type.max}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }

                if (type.endsWith) {
                    sb.l(`if(${opt}${name}.endsWith('${type.endsWith})') {`);
                    sb.i(() => {
                        sb.l(emitErr(`does not end with ${type.endsWith}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }

                if (type.startsWith) {
                    sb.l(`if(${opt}${name}.startsWith('${type.startsWith})') {`);
                    sb.i(() => {
                        sb.l(emitErr(`does not start with ${type.startsWith}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }

                if (type.anyOf) {
                    const values = emitValues(type.anyOf);
                    sb.l(`if(${opt}![${values}].includes(${name})) {`);
                    sb.i(() => {
                        sb.l(emitErr(`is not any of ${values}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }

                if (type.noneOf) {
                    const values = emitValues(type.noneOf);
                    sb.l(`if(${opt}[${values}].includes(${name})) {`);
                    sb.i(() => {
                        sb.l(emitErr(`is any of ${values}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }
                break;
            case "number":
            case "integer":
                required();

                if (type.min) {
                    sb.l(`if(${opt}${name} < ${type.min}) {`);
                    sb.i(() => {
                        sb.l(emitErr(`is less than ${type.min}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }

                if (type.max) {
                    sb.l(`if(${opt}${name} > ${type.max}) {`);
                    sb.i(() => {
                        sb.l(emitErr(`is greater than ${type.max}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }

                if (type.anyOf) {
                    sb.l(`if(${opt}![${emitValues(type.anyOf)}].includes(${name})) {`);
                    sb.i(() => {
                        sb.l(emitErr(`is not any of ${emitValues(type.anyOf!)}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }

                if (type.noneOf) {
                    sb.l(`if(${opt}[${emitValues(type.noneOf)}].includes(${name})) {`);
                    sb.i(() => {
                        sb.l(emitErr(`is any of ${emitValues(type.noneOf!)}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }
                break;
            case "array":
                required();

                if (type.min) {
                    sb.l(`if(${opt}${name}.length < ${type.min}) {`);
                    sb.i(() => {
                        sb.l(emitErr(`length is less than ${type.min}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }

                if (type.max) {
                    sb.l(`if(${opt}${name}.length > ${type.max}) {`);
                    sb.i(() => {
                        sb.l(emitErr(`length is greater than ${type.max}`));
                    });
                    sb.l(`}`);
                    sb.nl();
                }

                {
                    sb.l(`for(const item of ${name}) {`);
                    sb.i(() => {
                        this.emitValidation(sb, type.items, { name: "item" });
                    });
                    sb.l(`}`);
                    sb.nl();
                }
                break;
            case "ref":
                {
                    required();

                    const fn = () => {
                        sb.l(`errors.push(...${name}.validate())`);
                    };

                    if (type.optional) {
                        sb.l(`if(${opt}) {`);
                        fn();
                        sb.l(`}`);
                    } else {
                        fn();
                    }

                    sb.nl();
                }
                break;
            default:
                throw this.invalidTypeErr(type);
        }
    }

    private emitSql(
        name: string,
        fn: (sb: StringBuffer) => void,
    ): TS.TsMethodDecl {
        const sb = new StringBuffer();
        sb.l("return `");
        sb.nl();
        sb.i(() => {
            fn(sb);
        }, 4);
        sb.l("`");

        const decl: TS.TsMethodDecl = {
            kind: "method-decl",
            name,
            return: "string",
            body: sb.lines,
        };

        return decl;
    }

    private emitSqlCreate(type: BI.BoundedObject): TS.TsMethodDecl {
        return this.emitSql(
            "sqlCreate",
            (sb: StringBuffer) => this.sqlGen.emitCreate(sb, type),
        );
    }

    private emitSqlInsert(type: BI.BoundedObject): TS.TsMethodDecl {
        return this.emitSql(
            "sqlInsert",
            (sb: StringBuffer) => this.sqlGen.emitInsert(sb, type),
        );
    }

    private emitSqlSelect(type: BI.BoundedObject): TS.TsMethodDecl {
        return this.emitSql(
            "sqlSelect",
            (sb: StringBuffer) => this.sqlGen.emitSelect(sb, type),
        );
    }

    private emitSqlUpdate(type: BI.BoundedObject): TS.TsMethodDecl {
        return this.emitSql(
            "sqlUpdate",
            (sb: StringBuffer) => this.sqlGen.emitUpdate(sb, type),
        );
    }

    emitFilename(type: BI.BoundedKind, ...path: string[]): string {
        let name = this.bounded.typeName(type);
        name = this.namer.tsClass(name);
        name = `${name}.gen.ts`;
        name = npath.join(...path, name);

        if (path.length > 0 && path[0].startsWith(".") && !name.startsWith(".")) {
            name = "./" + name;
        }

        return name;
    }

    private emitTsClassName(type: BI.BoundedKind): string {
        let name = this.bounded.itemName(type);

        if (name === "") {
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`);
        }

        name = this.namer.tsClass(name);
        return name;
    }

    private emitTsVarName(type: BI.BoundedKind): string {
        let name = this.bounded.itemName(type);

        if (name === "") {
            throw new Error(`invalid type: missing name ${JSON.stringify(type)}`);
        }

        name = this.namer.tsVar(name);
        return name;
    }
}

export class BoundedString {
    config: {
        bounded: BoundedManager;
        namer: TsNamer;
        tsGen: TS.TsGenerator;
    };

    constructor(config: typeof this.config) {
        this.config = config;
    }

    property(sb: StringBuffer, kind: BI.BoundedString) {
        let varName = this.config.bounded.itemName(kind);
        varName = this.config.namer.tsVar(varName);
        let typeName = this.config.bounded.typeName(kind);

        if (this.config.bounded.isPrimitive(kind)) {
            typeName = this.config.namer.tsClass(typeName);
        }

        sb.a(`${varName}`);

        if (kind.optional) {
            sb.a(`?`);
        }

        sb.a(`: ${typeName}`);
    }
}

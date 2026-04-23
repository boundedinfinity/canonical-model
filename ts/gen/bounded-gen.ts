import { BoundedKind } from "./bounded-model";
import { BoundedTsGenerator } from "./bounded-ts-gen";

export class BoundedManager {
    private typeMap: { [name: string]: BoundedKind } = {};

    private invalidTypeErr(type: BoundedKind): Error {
        return new Error(`invalid token: ${JSON.stringify(type)}`);
    }

    register(token: BoundedKind) {
        switch (token.kind) {
            case "object":
                if (token.name && !(token.name in this.typeMap)) {
                    this.typeMap[token.name] = token;
                }
                break;
            default:
                throw this.invalidTypeErr(token);
        }
    }

    validate() {
        let errors: Error[] = [];

        for (const token of Object.values(this.typeMap)) {
            errors.push(...this._validate(token))
        }

        errors = [...new Set(errors)].filter(error => error !== undefined && error !== null)

        if (errors.length > 0) {
            throw errors;
        }
    }

    private _validate(token: BoundedKind): Error[] {
        const errors: Error[] = []

        switch (token.kind) {
            case 'object':
                for (const prop of token.properties) {
                    errors.push(...this._validate(prop))
                }
                break
            case 'array':
                errors.push(...this._validate(token.items))
                break
            case 'ref':
                if (!(token.ref in this.typeMap)) {
                    errors.push(this.invalidTypeErr(token))
                }
                break
            case "boolean":
            case "integer":
            case "number":
            case "string":
                // Nothing to do
                break;
            default:
                throw this.invalidTypeErr(token);
        }

        return errors;
    }

    lookup(token: BoundedKind): BoundedKind {
        let found: BoundedKind;

        switch (token.kind) {
            case "ref":
                found = this.typeMap[token.ref];

                if (!found)
                    this.invalidTypeErr(token)
                break;
            case "array":
                found = token.items;
                break;
            default:
                found = token;
        }

        return found;
    }

    isPrimitive(type: BoundedKind): boolean {
        const resolved = this.lookup(type);
        let primitive = true;

        switch (resolved.kind) {
            case "object":
                primitive = false;
                break;
        }

        return primitive;
    }

    isRef(type: BoundedKind): boolean {
        let found = false;

        switch (type.kind) {
            case "ref":
                found = true;
                break;
            case "array":
                found = this.isRef(type.items);
                break;
        }

        return found;
    }

    isArray(type: BoundedKind): boolean {
        let found = false;

        switch (type.kind) {
            case "ref":
                {
                    const rtype = this.lookup(type);
                    found = this.isArray(rtype);
                }
                break;
            case "array":
                found = true;
                break;
        }

        return found;
    }

    typeName(type: BoundedKind): string {
        let name: string | undefined;

        switch (type.kind) {
            case "object":
                name = type.name;
                break;
            case "array":
                name = this.typeName(type.items);
                break;
            case "ref":
                {
                    const rtype = this.lookup(type);
                    name = this.typeName(rtype);
                }
                break;
            case "boolean":
            case "number":
            case "integer":
            case "string":
                name = type.kind;
                break;
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`);
        }

        if (!name) {
            throw new Error(`cannot resolve type for ${JSON.stringify(type)}`);
        }

        return name;
    }

    itemName(type: BoundedKind): string {
        let name: string | undefined;

        switch (type.kind) {
            case "object":
                name = type.name;
                break;
            case "array":
                name = type.name ?? this.itemName(type.items);
                break;
            case "ref":
                name = type.name ?? this.typeName(type);
                break;
            case "boolean":
            case "number":
            case "integer":
            case "string":
                name = type.name;
                break;
            default:
                throw new Error(`invalid type ${JSON.stringify(type)}`);
        }

        if (!name && type.parent) {
            name = type.parent.name;
        }

        if (!name || name === "") {
            throw new Error(`invalid type ${JSON.stringify(type)}`);
        }

        return name;
    }
}


type Files = { [path: string]: string };

export class BoundedGenerator {
    private tsGen: BoundedTsGenerator;
    private manager = new BoundedManager();

    constructor() {
        this.tsGen = new BoundedTsGenerator(this.manager);
    }

    emit(...types: BoundedKind[]): Files {
        for (const type of types) {
            this.manager.register(type);
        }

        this.manager.validate()

        const files: Files = {};

        for (const type of types) {
            const content = this.tsGen.emit(type);
            const path = this.tsGen.emitFilename(type, "./gen-output/ts-gen");
            files[path] = content;
        }

        return files;
    }
}

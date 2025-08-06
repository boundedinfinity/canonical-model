import { NamePrefix } from "./name-prefix.schema";
import { NameSuffix } from "./name-suffix.schema";

export class Name {
    prefixes: NamePrefix[] = [];
    givenNames: string[] = [];
    familyNames: string[] = [];
    suffixes: NameSuffix[] = [];

    toString(): string {
        let names: string[] = [];

        if (this.prefixes && this.prefixes.length > 0) {
            names.push(...this.prefixes.map(p => p.toString()))
        }

        if (this.givenNames && this.givenNames.length > 0) {
            names.push(...this.givenNames)
        }

        if (this.familyNames && this.familyNames.length > 0) {
            names.push(...this.familyNames)
        }

        if (this.suffixes && this.suffixes.length > 0) {
            names.push(...this.suffixes.map(p => p.toString()))
        }

        return names.join(" ");
    }
}

import { Name } from "./name.schema";
import { NamePseudonym } from "./name-pseudonym.schema";

export class Person {
    names: Name[] = [];
    pseudonyms: NamePseudonym[] = [];

    name(): string {
        const names: string[] = [];

        if (this.names && this.names.length > 0) {
            names.push(this.names[0].toString());
        }

        if (this.pseudonyms && this.pseudonyms) {
            let pseudonym = this.pseudonyms[0].toString();

            if (names.length > 0) {
                pseudonym = `(${pseudonym})`;
            }

            names.push(pseudonym);
        }

        return names.join(" ");
    }
}

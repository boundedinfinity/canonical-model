import { Description, DescriptionTextType } from "./description.schema";

export class NamePrefix {
    id: string;
    title: string;
    description?: Description;
    abbreviations?: string[];

    constructor(
        id: string,
        title: string,
        description?: Description,
        abbreviations?: string[]
    ) {
        this.id = id;
        this.title = title;
        this.description = description;
        this.abbreviations = abbreviations;
    }

    toString() : string {
        if(this.abbreviations && this.abbreviations.length > 0) {
            return this.abbreviations[0]
        }

        return this.title
    }
}

export const List: NamePrefix[] = [
    new NamePrefix(
        "https://www.boundedinfinity.com/canonical/people/prefix/mister",
        "Mister",
        new Description(
            `A title used before a surname or full name to address or refer to a man without a higher, honorific or professional title.`,
            DescriptionTextType.Plain
        ),
        ["Mr."]
    ),
    new NamePrefix(
        "https://www.boundedinfinity.com/canonical/people/prefix/miss",
        "Mistress",
        new Description(
            `A title used before a surname or full name to address or refer to a unmarried or young woman without a higher, honorific or professional title.`,
            DescriptionTextType.Plain
        ),
        ["Miss."]
    ),
];

export enum NamePrefixKeys {
    Mr = "Mr",
    Miss = "Miss",
}

export const Map: { [key: string]: NamePrefix } =
    List.reduce((map: { [key: string]: NamePrefix }, obj) => {
        map[obj.id] = obj;
        return map;
    }, {});

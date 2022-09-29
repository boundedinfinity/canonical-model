import { Description, DescriptionTextType } from "./description.schema";

export class NamePseudonymType {
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

    toString(): string {
        if (this.abbreviations && this.abbreviations.length > 0) {
            return this.abbreviations[0];
        }

        return this.title;
    }
}

export const List: NamePseudonymType[] = [
    new NamePseudonymType(
        "https://www.boundedinfinity.com/schema/people/pseudonym-type/nickname",
        "Nickname",
        new Description(
            `A type of pseudonym which is used as a substitue for the proper name of a familiar person.`,
            DescriptionTextType.Plain
        )
    ),
];

export const Map: { [key: string]: NamePseudonymType } = List.reduce(
    (map: { [key: string]: NamePseudonymType }, obj) => {
        map[obj.id] = obj;
        return map;
    },
    {}
);

export class NamePseudonym {
    id: string;
    title: string;
    type?: NamePseudonymType;

    constructor(
        id: string,
        title: string,
        type?: NamePseudonymType,
    ) {
        this.id = id;
        this.title = title;
        this.type = type;
    }

    toString(): string {
        return this.title;
    }
}

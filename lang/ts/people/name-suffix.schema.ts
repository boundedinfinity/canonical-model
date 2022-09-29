import { Description, DescriptionTextType } from "./description.schema";

export class NameSuffix {
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

export const List: NameSuffix[] = [
    new NameSuffix(
        "https://www.boundedinfinity.com/schema/people/suffix/ba",
        "Bachelor of Arts",
        new Description(
            `A degree from a university or other higher learning instituion awarded for an undergraduate program in the arts or other disciplines.`,
            DescriptionTextType.Plain
        ),
        ["BA", "AB"]
    ),
    new NameSuffix(
        "https://www.boundedinfinity.com/schema/people/suffix/bs",
        "Bachelor of Sciences",
        new Description(
            `A degree from a university or other higher learning instituion awarded for an undergraduate program in the sciences.`,
            DescriptionTextType.Plain
        ),
        ["BS", "B.Sc.", "SB", "ScB."]
    ),
];

export enum NameSuffixKeys {
    Mr = "Mr",
    Miss = "Miss",
}

export const Map: { [key: string]: NameSuffix } =
    List.reduce((map: { [key: string]: NameSuffix }, obj) => {
        map[obj.id] = obj;
        return map;
    }, {});

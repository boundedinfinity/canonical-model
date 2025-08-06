import { describe, expect, it } from "vitest";
import { Reviver, Jsonizer } from "@badcafe/jsonizer";
// https://badcafe.github.io/jsonizer

@Reviver<Name>({ 
    '.':  Jsonizer.Self.apply(Name),
})
class Name {
    constructor(
        public prefixes: string[],
        public names: string[],
        public suffixes: string[]
    ) {}
}

@Reviver<Person>({ 
    '.': Jsonizer.Self.apply(Person),
    bdate: Date,
})
class Person {
    constructor(public names: Name[], public bdate: Date) {}
}



describe("Serialization", function () {
    it("serializies an class", function () {
        const input = new Person([new Name(['Mr'], ['John', 'Smith'], ['Jr.'])], new Date('2022-09-26T17:42:22.014Z'))
        const expected = '{"names":[{"prefixes":["Mr"],"names":["John","Smith"],"suffixes":["Jr."]}],"bdate":"2022-09-26T17:42:22.014Z"}'
        const actual = JSON.stringify(input)
        expect(actual).toEqual(expected)
    });

    it("deserializies an class", function () {
        const input = '{"names":[{"prefixes":["Mr"],"names":["John","Smith"],"suffixes":["Jr."]}],"bdate":"2022-09-26T17:42:22.014Z"}'
        const expected = new Person([new Name(['Mr'], ['John', 'Smith'], ['Jr.'])], new Date('2022-09-26T17:42:22.014Z'))
        const actual = JSON.parse(input, Reviver.get(Person))

        expect(actual.names.length).toEqual(expected.names.length)
        expect(actual.bdate).toBeInstanceOf(Date)
        expect(actual.bdate).toEqual(expected.bdate)
    });
});

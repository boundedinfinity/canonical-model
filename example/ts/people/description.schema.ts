export enum DescriptionTextType {
    Plain = "Plain",
    Markdown = "Markdown",
    HTML = "HTML",
}

export class Description {
    public readonly text: string;
    public readonly type?: DescriptionTextType;

    constructor(text: string, type?: DescriptionTextType) {
        this.text = text;
        this.type = type || DescriptionTextType.Plain;
    }
}

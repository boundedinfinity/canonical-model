export class Name {
    id: string
    prefix?: NamePrefix[]
    first?: string[]
    middle?: string[]
    last?: string[]
    suffix?: NameSuffix[]
    
    constructor(
        options: NameInterface,
    ){
        this.id = options.id
        this.prefix = options.prefix
        this.first = options.first
        this.middle = options.middle
        this.last = options.last
        this.suffix = options.suffix
    }
}
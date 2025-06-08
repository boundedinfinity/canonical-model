interface NamePrefixInterface {
    id: string
    name: string
    abbr?: string[]
    description?: string
}

export class NamePrefix {
    id: string
    name: string
    abbr?: string[]
    description?: string
    
    constructor(options: NamePrefixInterface, ){
        this.id = options.id
        this.name = options.name
        this.abbr = options.abbr
        this.description = options.description
    }
}

enum NamePrefixConvertType {
    Full,
    Embedded,
}

interface NamePrefixObjectOptions {
    marshalType: NamePrefixConvertType
}

export class NamePrefixObject {
    options: NamePrefixObjectOptions
    
    constructor(options: Partial<NamePrefixObjectOptions>, ){
        this.options = options
    }
    
    marshal(item: NamePrefix, ): object {
        const obj = {
            id: item.id,
            name: item.name,
            abbr: item.abbr,
            description: item.description,
        }
        return obj
    }
    
    unmarshal(item: NamePrefixInterface, ): NamePrefix {
        return new NamePrefix(item, )
    }
}

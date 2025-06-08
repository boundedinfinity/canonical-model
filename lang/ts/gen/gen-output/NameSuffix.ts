interface NameSuffixInterface {
    id: string
    name: string
    abbr?: string[]
    description?: string
}

export class NameSuffix {
    id: string
    name: string
    abbr?: string[]
    description?: string
    
    constructor(options: NameSuffixInterface, ){
        this.id = options.id
        this.name = options.name
        this.abbr = options.abbr
        this.description = options.description
    }
}

enum NameSuffixConvertType {
    Full,
    Embedded,
}

interface NameSuffixObjectOptions {
    marshalType: NameSuffixConvertType
}

export class NameSuffixObject {
    options: NameSuffixObjectOptions
    
    constructor(options: Partial<NameSuffixObjectOptions>, ){
        this.options = options
    }
    
    marshal(item: NameSuffix, ): object {
        const obj = {
            id: item.id,
            name: item.name,
            abbr: item.abbr,
            description: item.description,
        }
        return obj
    }
    
    unmarshal(item: NameSuffixInterface, ): NameSuffix {
        return new NameSuffix(item, )
    }
}

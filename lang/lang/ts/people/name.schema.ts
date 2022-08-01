class Name {
    prefixes: string[] = []
    names: string[] = []
    suffixes: string[] = []

    fullName(): string {
        let name: string = ""

        if(this.prefixes) {
            name += this.prefixes.join(" ")
            name += " "
        }

        if(this.names) {
            name += this.names.join(" ")
            name += " "
        }

        if(this.suffixes) {
            name += this.suffixes.join(" ")
        }
        
        return name
    }

    validate() {
        
    }
}

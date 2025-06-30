class BoundedObjectTemplate {
    REQUIRED_STRING: string

    constructor() {
        this.REQUIRED_STRING = ''
    }

    validate() {
        const errors: Error[] = []

        if (!this.REQUIRED_STRING)
            errors.push(new Error(`REQUIRED_STRING is required`))

        if (this.REQUIRED_STRING.length < 0)
            errors.push(new Error(`REQUIRED_STRING is less than 0 characters`))

        if (this.REQUIRED_STRING.length > 0)
            errors.push(new Error(`REQUIRED_STRING is greater than 0 characters`))

        return errors
    }
}

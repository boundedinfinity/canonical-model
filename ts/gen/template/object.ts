class BOUNDED_OBJECT {
    BOUNDED_STRING: string

    constructor(
        args: {
            BOUNDED_STRING: string
        }
    ) {
        this.BOUNDED_STRING = args.BOUNDED_STRING
    }

    validate(): Error[] {
        const errors: Error[] = []

        return errors
    }
}
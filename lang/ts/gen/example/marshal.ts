export interface Unmarshaller<I, O> {
    unmarshal(raw: I): O
}

export interface Marshaller<I, O> {
    marshal(item: I): O
}
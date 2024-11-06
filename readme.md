## Specification

### Types

-   `type`:

    The type of this type.

    This field can contain one of the built-in types:

    -   `string`
    -   `integer`
    -   `float`
    -   `array`
    -   `object`
    -   `enum`
    -   `uuid`
    -   `union`

    or the full URL of another type's `id` field.

    If the type's value is the `id` field of another type, this turns
    this type into a referenced type.

    If this is a referenced type, then referenced type's fields will
    be inserted into this type. If this type defines any fields, the
    fields in this type will overried the fields from the given `id`
    referenced type.

-   `id`:

    The identifier of this type.

    This is an optional field.

    If this field is present it indicates that this is a custom type.
    Code generation tooling should create a custom class, struct, or
    relavent type for the language.

    The last portion of the `id` field will be used in a language
    appropriate way for the targetied language, if no other name
    overrides the name by defining the `name` property.

-   `name`:

    The name of this type.

    If this field is defined if may do one of the following:

    -   If defined on a type with a defined `id` field, this type will
        become the name of the custom type.
    -   Otherwise this will become a field name inside the `properties`
        field of an `object` type.

    If this field isn't defined a name is created using the method described
    in the `id` section.

-   `description`:

    A common field for all types.

    If defined this fields should be a detailed description of this
    type.

-   `min`:

    Optional minimum of this type.

    If defined on an `integer` or `float` type this will be the inclusve
    minimum value which that type will contain.

    If defined on a `string` type, this will be the minimum length of
    the string.

    This will be an error on any other type.

-   `max`:

    Optional maximum of this type.

    If defined on an `integer` or `float` type this will be the inclusve
    maximum value which that type will contain.

    If defined on a `string` type, this will be the maximum length of
    the string.

    This will be an error on any other type.

-   `regex`:

    Optional regular expression pattern.

    The given `string` value must conform to the pattern defined in the
    `regex` field.

    This will be an error on any other type.

-   `query`:

    Marks this field as queryable.

### Operations

-   `id`:

    The identifier of this operation.

    This is a required field.

    The last portion of the `id` field will be used in a language
    appropriate way for the targetied language, if no other name
    overrides the name by defining the `name` property.

-   `name`:

    The name of this operation.

    If this field isn't defined a name is created using the method described
    in the `id` section.

-   `inputs`:

    List of input types for this operation.

    These follow the same rules defined in the **Types** section.

-   `outputs`:

    List of output types for this operation.

    These follow the same rules defined in the **Types** section.

### Data

-   `type`:

    The URL of the an `id` from the **Types** section.

-   `items`:

    List of items that must conform to the shape of the type referenced
    by the `id` pointed to in the `type` field.

## Reference

-   https://www.learnjsonschema.com/2020-12/core/
-   https://json-schema.org/understanding-json-schema/reference/string
-   https://zod.dev/?id=table-of-contents
-   https://entgo.io/
-   https://docs.asciidoctor.org/asciidoc/latest/
-   https://typeschema.org/
-   https://jsontypedef.github.io/json-typedef-js/index.html
-   https://protobuf.dev/
-   https://codalogic.github.io/jcr/
-   https://raml.org/
-   https://apiblueprint.org/
-   https://schema.org/
-   https://graphql.org/
-   https://azimutt.app/blog/aml-a-language-to-define-your-database-schema#relation-definition
-   https://avro.apache.org/

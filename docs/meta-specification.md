## Common specifications

-   `id`:

    The type ID.

    required if this is a type definition.  

-   `name`: 

    The name of the type.
    
    required: **false**

    If the name is not given. The name of this type if required
    will be generated from the type ID.

-   `type`: 

    The type ID of the the type.

    required: **true**

    The type ID can be one of the following built-in types basic types:
    - `integer`
    - `float`
    - `boolean`
    - `string`
    - `time`
    - `date`
    - `date-time`
    - `duration`
    - `uuid`

    or one of the compound types:
    - `object`
    - `array`
    - `enum`
    - `range`
    - `ref`

    or a URL encoded reference to another type's `type` parameter.

-   `version`: 

    Version of this item.

    required: **false**
-   `upgrade`:  

    A reference to another specification as the upgraded, or
    the next version, of this type

    required: **false**
-   `downgrade`:  

    A reference to another specification as the downgraded, or
    previous version, of this type

    required: **false**
-   `required`: 

    A boolean which determines if this item is required or not

    required: **false**
-   `default`: 

    If provided, the type specific value which will be assigned
    if not value is given.

    required: **false**

-   `example[]`: 

    If provided, list of the type specific examples for a given value.

    required: **false**


## Integer speficiation

-   `ranges[]` 
    
    The list of range constraints for the number

    This contains a list of the following sub-constraints:

    -   `min`: 
        
        The inclusive minimum constraint for the number

        This item is mutually exclusive with the `exclusive-min` constraint
    -   `max`: 
        
        The inclusive maximum for the number

        This item is mutually exclusive with the `exclusive-max` constraint
    -   `exclusive-min`: The exclusive minimum for the number

        This item is mutually exclusive with the `min` constraint
    -   `exclusive-max`: The exclusive maximum for the number

        This item is mutually exclusive with the `max` constraint
-   `multiple-of`: Constrains the number to be a multiple of this value.
-   `negative`: Constrains the number to be less than `0`

    This constraints is mutually exclusive with the `positive` constraint.
-   `positive`: Constrains the number to be greater than `0`

    This constraints is mutually exclusive with the `negative` constraint.
-   `excludes`: 

    List of integers which are not value.

    NOTE: These values are combined with values for from other contraints.
-   `includes`: 

    List of integrars which are considered valid.

    NOTE: These values are combined with values for from other contraints.

## Bool specification

There are no additional items for this type

## Float speficiation

The `float` constraints are the same as for the `integer` type with the
following additions:
-   `tolerance`:

    Is an equality constraint which calculates the difference
    between another constraint and if the difference is below the `tolerance`
    the equality is assumed to be true.
-   `precision`:

    Comparse a number down to a certain number of places below
    the period.

## String specification
-   `min`: 

    The inclusive minimum length of the string
-   `max`: 

    The inclusive maximum lenght of the string
-   `regex`: 

    A regular expression to which the string must conform.
-   `abnf`: 

    An ABNF expression to which the string must conform
-   `includes`: 

    List of sub-string items which must appear in the string
-   `excludes`: 

    List of sub-string items which must not appear in the string

## Enum specification
-   `items`:

    List if sub-items in these enumeration

    -   `name`:
    
        The name of the item
    -   `values[]`:

        The list of values that represent this item.  
        
        The first item in the list should be the primary item.
    -   `description`:

        The description of this item.

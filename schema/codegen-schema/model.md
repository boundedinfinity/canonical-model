# types
  
A list of types contained within the system.

Types can be one of the following:

- `bool`
- `int`
- `float`
- `string`
- `enum`
- `array`
- `object`

## The `int` type

The integer type. The integer type has the following properties:

- `min`
  
  The inclusive minimum constraint value of the integer.
- `max`
  
  The inclusive maximum constraint value of the integer.
- `multiple-of`
  
  The integer constraint must be a multiple of the values.  Meaning 
  that if the result of `value` modulus of `multiple-of` evaluates 
  to `0`.
- `required`
  
  Determines if the integer is `required`.  If this value is set to
  `true` the value must be present, and must conform to any of the
  constraint values.
- `default`
  
  The default value of the integer is the integer isn't defined.  If
  the the default is defined it must conform to the contraints.
- `description`
  A description of this field.

  In the generated code this will typcially be integrated into the 
  target languages's comments.

## The `float` type

The float type. The float type has the same configurations as the `int`
type.


## The `bool` type

The boolean type. The boolean type has the following properties:

- `required`
  
  Determines if the boolean is `required`.  If this value is set to
  `true` the value must be present.
- `default`
  
  The default value of the boolean is the boolean isn't defined.
- `description`
  A description of this field.

  In the generated code this will typcially be integrated into the 
  target languages's comments.
types:
    -   type: array
        items:
            type: object
            properties:
                -   name: id
                    type: string
                -   name: name
                    type: string
                -   name: description
                    type: string
                -   name: properties
                    type: array
                    items:
                        ref: type
operations:
    -   type: array
        items:
            type: object
            properties:
                -   name: id
                    type: string
                -   name: name
                    type: string
                -   name: inputs
                    type: object
                    properties:
                        -   ref: type
                -   name: outputs
                    type: object
                    properties:
                        -   ref: type
data:
    -   type: array
        items:
            type: object
            properties:
                -   type: type
                    
            
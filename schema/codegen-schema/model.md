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



## The `bool` type

The boolean type. The boolean type has the following properties:

- `name`

    The name of this value.  If this isn't present the type name of the
    will be used instead.

- `required`
  
  Determines if the boolean is `required`.  If this value is set to
  `true` the value must be present.

- `default`
  
  The default value of the boolean is the boolean isn't defined.

- `description`

  A description of this field.

  
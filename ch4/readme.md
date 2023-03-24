

## Reflection

- Why was the reflection included in Go?
- When should I use reflection?

*Ans*
1. Reflection allows you to dynamically learn the type of an arbitrary object among with the information about its structure.
2. Reflection allows you to handle and work with data types that do not exist at the time at which you write your code but might exist in the future.


`reflect.Value`

`reflect.Type`

`reflect.ValueOf()` returns reflect.Value

`reflect.TypeOf()` returns reflect.Type

`reflect.NameField()` lists the number of fields in structure.

`reflect.Kind`

### disadvantages of reflection
1. hard to read 
2. slower than particular data type
3. refection error cannot be caught at build time
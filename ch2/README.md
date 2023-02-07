# Basic Go Data Types

## Numeric  data type
- int: int is -2^31 ~ 2^31
- unit: 0 ~ 2^32

## Non-numeric data type
A GO string is just a collection of types and can be accessed as a whole or as a array.
- A single byte can store any ASCII.
- Multiple bytes are usually Unicode.

### 
- strings (unit8)
- characters
- runes (support Unicode, is a int32, but can't compare with int32, GO considers rune & int32 is 2 types)
- dates
- 

## Zero Value 


## Slice 

- array's size cannot be changed
- when pass a array to a function, GO creates a copy of that array, therefore any changes we made will be lost
- slice value is a header that contains a **pointer to an underlying array**
- the length of underlying array is slice's capacity
- side effect: passing the slice header is faster than pass a copy slice to function
- if you don't want to initialize a slice, then using **make()** is better & faster.
- capacity can't be bigger than origin, unless using append to allocate new one
- []byte is string
- the length of []byte might not be the same as the length of the string (unicode)
- no default function for deleting an element from a slice 




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

## Pointers
* Go has support Pointers but not for pointer arithmetic. (cause bugs & errors)
* Why we use pointers in Go, the main benefit is that passing a variable to a function as a pointer does not discard any changes you make.
* if a slice won't change it size, we don't need passing pointer (slice is a pointer to underlying array)

### Pros:
1. share pointer between functions/goroutines, we should be extra careful with race condition issues.
2. use pointer to tell difference between zero value and value that is not set(nil value) this is useful with structure.
3. support Linked List & Tree, without pointer, it's hard to implement and may be too slow

## Random Number/String

- You do not need to define a seed when using the `crypto/rand` package



# Go

## Imports

* import packages with:

```go
import (
    "fmt"
    "os"
)
```

## Variables

* If any declared is not used it'll result in a compilation error.

* Cannot redefine variables unless it is done when declaring a new variable along the reassignment, also not possible to change the type of a previously declared variable.
  
```go
//Allowed
foo := 123
foo, bar := 42, "baz"

//Not allowed 
foo := "bar"
foo, baz := 10, 123
```

## Functions

* Functions can return multiple values, i.e. a tuple of a int and a string. (Note that if you return a tuple the actual return value in the declaration must be enclosed by paranthesis while the return statement has no such requirement)
* Like in Haskell we can ignore values with `_`, might be useful when eg. returning a tuple for which we only care about one if the return values.
* If the parameters of a function are the same you can omit the type for the first parameter

Example:

```go
function (a,b int) int {
    //do stuff
}
```

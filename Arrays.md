### Arrays
- fixed lenght series of elements of chosen datatype
- elements accessed using subscript notation []
- indices start at 0
- elements initialized to 0
- eg:

```
var x [5]int
x[0] = 2
fmt.Println(x[1])
```

### Array Literal
- an array predefined with values
- eg:

```
var x [5]int = [5]{1,2,3,4,5}
```

- length of array literal should be equal to length of array
- use ... opeartor to infer size of array literal from the number of initializers
- eg:

```
x := [...]int{1,2,3,4}
```

### Iterating through array
- use a for loop with range keyword
- eg:

```
x := [3]int {1,2,3}
for i,v range x{
fmt.Printf("ind %d,val %d",i,v)
}
```

- range returns 2 values, index and element at that index

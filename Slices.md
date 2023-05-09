### Slices
- a window on an underlying array
- has variable size, upto the size of the array
- pointer indicates the start of the slice
- length indicates the number of elements in the slice
- capacity indicates the maximum number of elements in the slice
- capacity ranges from start of the slice to the end of the array
- eg:

```
arr := [...]string{"a","b","c","d","e","f","g"}
s1 := arr[1:3] // [beg. index : end index (not included)] -> {b,c}
s2 := arr[2:5] // ->{c,d,e}
```
### Length and Capacity
- len() returns the length of the slice
- cap() returns the capacity of the slice
- eg:

```
a1 := [3]string("a","b","c")
s1 := a1[0:1]
fmt.Printf(len(s1), cap(s1)) // prints 1,3
```

### Accessing Slices
- writing to a slice changes the underlying array
- overlapping slice refers to the same array elements

### Slice Literals
- can be used to initialize a slice
- creates the underlying array and creates a slice to reference the entire array
- slice points to the start of the array, length is same as capacity
- eg:

```
s1 := []int{1,2,3}
```


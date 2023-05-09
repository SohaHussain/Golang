### Make
- create a slice and array using make()
- 2 arguments version : specify type and length/capacity
```
s1 = make([]int,10)
```
- 3 arguments version : datatype, length and capacity
```
s2 = make([]int,10,15)
```

### Append
- size of the slice can be increased using append()
- adds element to the end of the slice
- inserts into the underlying array
- increases the size of the array if necessary
```
sl = make([]int,0,3)
```
- length of sl is 0
```
sl = append(sl,100)
```

- a pointer is an address to data in memory
- & -> returns address of a variable
- * -> returns data at an address (dereferencing)
- eg:

```var x int = 1
var y int
var ip *int // ip is pointer of type int
ip = &x // ip now points to x
y = *ip //y is now 1
```

- new() is an alternate way to create a variable 
- new() function creates a variable and returns a pointer to the variable
- variable is initialized to 0 by default

```ptr := new(int)
*ptr = 3
```

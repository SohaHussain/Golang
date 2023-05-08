- places in code where a variable can be accessed
- eg:

```var x = 4
func f(){
fmt.Printf("%d",x) // will print 4
}
func g(){
fmt.Printf("%d",x) // will print 4
}
```

```func f(){
var x = 4
fmt.Printf("%d",x) // will print 4
}
func g(){
fmt.Printf("%d",x) // will throw an error
}
```

- in Go variable scoping is done using blocks

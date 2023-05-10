### Variable argument number
- functions can take variable number of arguments
- use ellipsis ... to specify
- treated as a slice inside function

```
func getmax(val ... int) int{
maxv := -1
for _, v := range val{
if v > maxv{
maxv = v
}
}
return maxv
}
```

### Variadic slice argument
- can pass a slice to a variadic function
- need the ... suffix

```
func main(){
fmt.Println(getmax(1,3,6,4))
vslice := []int{1,3,6,4}
fmt.Println(getmax(vslice...))
}
```

### Deferred function calls
- call can be deferred until the surrounding function completes
- typically used for cleanup activities

```
func main(){
defer fmt.Println("bye")
fmt.Println("hello")
}
```

- arguments of a deferred call are evaluated immediately

```
func main(){
i := 1
defer fmt.Println(i+1) // prints 2
i++
fmt.Println("hello")
}
```

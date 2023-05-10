### Passing array arguments
- array arguments are copied
- arrays can be big, so this can be a problem

```
func foo(x [3]int) int{
return x[0]
}
func main(){
a := [3]int{1,2,3}
fmt.Print(foo(a))
}
```
- it is possible to pass array pointers

```
func foo( x*[3]int) int{
(*x)[0] = (*x)[0]+1
}
func main(){
a := [3]int{1,2,3}
foo(&a)
fmt.Print(a)
}
```
### Pass slices instead
- slices contain a pointer to the array
- passing a slice copies the pointer

```
func foo(sl int) int{
sl[0] = sl[0]+1
}
func main(){
a := []int{1,2,3} // declared without size, so its a slice
foo(a)
fmt.Print(a)
}
```

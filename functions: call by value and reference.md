### Call by value
- passed arguments are copied to parameters
- modifying parameters has no effect outside the function

```
func foo(y int){
y = y+1
}
func main(){
x := 2
foo(x)
fmt.Println(x) // prints 2
}
```
### Tradeoffs of call by value
- advantage : data encapsulation -> function variables only changed inside the function
- disadvantage : copying time -> large objects  may take a long time to copy

### Call by reference
- programmer can pass a pointer as an argument
- called function has direct access to caller variable in memory

```
func foo(y *int){
*y = *y+1
}
func main(){
x := 2
foo(&x)
fmt.Print(x)
}
```
### Tradeoffs of call by reference
- advantage : copying time -> don't need to copy arguments
- disadvantage : data encapsulation -> function variables may be changed in called functions

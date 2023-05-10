### Function parameters
- parameters are listed in parenthesis after function name
- arguments are supplied in the call

```
func foo(x int, y int){
fmt.Print(x*y)
}
func main(){
foo(2,3)
}
```
### Parameter options
- if no parameters, use empty parenthesis
- list arguments of same type

```
func foo(x,y int){
}
```
### Return values
- type of return value is declared after parameters
- function call is used on right hand side of assignment

```
func foo(x int) int{
return x+1
}
y := foo(1)
```
### Multiple return values
- multiple value types must be listed in declaration

```
func foo2(x int) (int,int){
return x, x+1
}
a, b := foo2(3)
```

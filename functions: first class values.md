### Functions are first class
- functions can be treated like other types
- variables can be declared with a function type
- can be created dynamically
- can be passed as arguments and returned as values
- can be stored in data structures

### Variables as functions
- declare a variable as a function

```
var funcvar func(int) int
func incfn(x int) int{
return x+1
}
func main(){
funcvar = incfn
fmt.Println(funcvar(1))
}
```
### Functions as arguments
- function can be passed to another function as argument

```
func applyit(afunct func(int) int, val int) int{
return afunct(val)
}
```

```
func applyit(afunct func(int) int, val int) int{
return afunct(val)
}
func inct(val int) int {return val+1}
func decc(val int) int {return val-1}
func main(){
fmt.Println(applyit(inct,2))
fmt.Println(applyit(decc,2))
}
```

### Anonymous functions
- don't need to name a function
- used when functions are passed as arguments

```
func applyit(afunct func(int) int , val int) int{
return afunct(val)
}
func main(){
v := applyit(func(x int) int {return x+1}, 2)
fmt.Println(v)
}
```

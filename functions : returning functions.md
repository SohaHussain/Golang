### Functions as return values
- functions can return functions
- might create a function with controllable parameters
- eg : distance to origin function
-> takes a point (x,y,coordinates)
and returns distance to origin
- what if you want to change origin
- option 1: pass origin as argument
- option 2: define function with new origin

```
func makedistorigin(ox,oy float64) func (float64, float64) float64{
fn := func(x,y float64) float64{
return math.Sqrt(math.Pow(x-ox,2) + math.Pow(y-oy,2))
}
return fn
}
```

- origin location is passed as argument
- origin is built into returned function

```
func main(){
dist1 := makedistorigin(0,0)
dist2 := makedistorigin(2,2)
fmt.Println(dist1(2,2))
fmt.println(dist2(2,2))
}
```

### Environment of a function
- set of all names that are valid inside a function
- names defined locally in the function
- lexical scoping
- environment includes names defined in the block where function is defined

```
var x int
func foo(y int){
z:= 1
...
}
```

- x,y,z are included in the environment

### Closure
- function + its environment
- when functions are passed / returned, their environment comes with it

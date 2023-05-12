### Ways to use an interface
- need a function which take multiple types of parameters

```
type shape2d interface{
area() float64
peri() float64
}

type triangle{....}
func( t triangle ) area() float64 {....}
func( t triangle) peri() float64 {.....}
type rectangle{.....}
func(r rectangle) area() float64{........}
func(r rectangle) peri() float64{.....}
```

- rectangle and triangle satisfy shape2d interface

```
func fitinyard(s shape2d) bool{
if(s.area()>100 && s.peri()>100){
return true
}
return false
}
```

- parameter is any type that satisfies shape2d interface

### Empty interfaces
- empty interface specifies no methods
- all types satifies empty interface
- use it to have a function accept any type as parameter

```
func printme(val interface{}){
fmt.Println(val)
}
```

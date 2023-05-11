### No 'Class' keyword
- most OO languages have class keyword
- data fields and methods are defined inside a class block

### Associating methods with data
- method has a recienver type that it is associated with
- use dot notation to call a method

```
type myint int
func (mi myint) Double () int{  // Double is the func name
return int(mi*2)
}
func main(){
v := myint(3) // v is the object
fmt.Println(v.Double())
}
```

### Implicit Method argument

```
func (mi myint) Double () int{
return int(mi*2)
}
func main(){
v := myint(3)
```

- object v is an implicit argument (hidden argument) to the method
- object v is passed by value

### Structs with methods
- structs and methods together allow arbitary data and functions to be composed

```
type point struct{
x float64
y float64
}

func ( p point) disttoorig() {
t := math.Pow(p.x,2) + math.Pow(p.y,2)
return math.Sqrt(t)
}
func main(){
p1 := point(3,4)
fmt.Println(p1.disttoorig())
}
```

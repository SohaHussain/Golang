### Interfaces
- concept in Go used to accomplish polymorphism
- set of method signatures .i.e. name, parameters, return types
- implementation is not defined
- used to express conceptual similarity between types

### Satisfying an interface
- type satisfies an interface if type defines all methods specified in the interface -> same method signature
- additinal methods are OK
- similar to inheritance with overriding

### Defining an interface type

```
type shape2d interface{
area() float64
perimeter() float64
}
type Triangle {.....}
func(t Triangle ) area() float64 {.....}
func(t Triangle) perimeter() float64 {.....}
```

- here Triangle type satisfies the shape2d interface

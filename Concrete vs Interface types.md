### Concrete Types
- specify the exact representation of data and methods
- complete method implementation is included

### Interface types
- specifies some method signatures
- implementations are abstract

### Interface values
- can be treated like other values
- interface values have two components
- 1. dynamic type ->concrete type which it is assigned to
- 2. dynamic value -> value of the dynamic type
- interface value is actually a pair (dynamic type, dynamic value)

### Defining an interface type

```
type speaker interface{speak()}
type dog struct{name string}
func (d dog)speak(){
fmt.Println(d.name)
}
func main(){
var s1 speaker
var d1 dog{"brian"}
s1 = d1
s1.speak()
}
```

- dynamic type is dog and dynamic value is d1

### Interface with nil dynamic value
- an interface can have a nil dynamic value

``` 
var s1 speaker
var d1 *dog
s1=d1
```

- d1 has no concrete value yet
- s1 has a dynamic type but no dynamic value
- can still call speak() method of s1
- doesnt need a dynamic value to call
- need to check inside the method

```
func (d *dog) speak(){
if d== nil{
fmt.Println("noise")
} else{
fmt.Println(d.name)
}
}
var s1 speaker
var d1 *dog
s1=d1
s1.speak()
}
```



### Structs
- aggregate data type
- groups together other objects of arbitary type
- eg:
```
type struct person{
name string
addr string
phone string
}
var p1 person
```
- each property is a field
- p1 contains values for all fields

### Accessing Struct Fields
- use dot notation
```
p1.name = "joe"
x = p1.addr
```

### Initializing Structs
- can use new()
- initializes fields to 0
```
p1 := new(person)
```
- can initialize using a struct literal
```
p1 := person(name: "joe",
addr : "a st.", phone : "123")
```

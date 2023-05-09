### Generating JSON representation from an object
``` type struct person{
name string
addr string
phone string
}
```

- Marshal() returns json representation as []byte
```
p1 := person(name:"joe",
addr:"a st.", phone :"123")
barr, err := json.Marshal(p1)
```
- Unmarshal() converts a json []byte into a Go object
```
var p2 person
err := json.Unmarshal(barr,&p2)
```
- pointer to go object is passed to Unmarshal()

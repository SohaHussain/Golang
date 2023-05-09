### Maps
- implementation of hash table
- use make() to create a map
```
var idmap map[string]int //[value type] key type
idmap = make(map[string]int)
```
- may define a map literal
```
idmap := map[string]int{
"joe":100
}
```

### Accessing Maps
- referncing a vlaue with [key]
- returns 0 if key is not present
```
fmt.Println(idmap["joe"])
```
- adding a key-value pair
```
idmap["jane"] = 456
```
- deleting a key-value pair
```
delete(idmap,"joe")
```

### Map functions
- two value assignment test for existence of key
```
id,p := idmap["joe"]
```
- id is value and p is a boolean indicating the presence of key
- len() returns how many key-value pairs are inside the map
```
fmt.Println(len(idmap))
```

### Iterating through maps
- use a for loop with range keyword
- two value assignment with range
```
for key, val := range idmap{
fmt.Println(key,val)
}
```

### Limitations of methods
- receiver is passed implicitly as an argument to the method
- method cannot modify the data inside the receiver

### Point receivers

```
func (p *point) offsetx(v float64){
p.x = p.x+v
}
```

- receiver can be a pointer to a type
- call by reference , pointer is passed to the method
- dereferencing is automatic with . operator 
- we do not need to write
```
*p.x = *p.x +v
```

- no need to reference when calling the method

```
func main(){
p := point(3,4)
p.offswtx(5)
fmt.Println(p.x)
}
```

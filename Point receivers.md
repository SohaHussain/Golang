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

### Type assertions for Disambiguation
- type assertions can be used to determine and extract the underlying concrete type

```
func Drawshape( sshape2d ) bool{
rect, ok := s.(Rectangle)
if ok{
DrawRect(rect)
}
tri , ok := s.(Triangle)
if ok{
DrawRect(tri)
}
}
```

- type assertion extracts rectangle from shape2d -> concrete type in parenthesis
- if interface contains concrete type -> rect==concrete type , ok==true
- if interface does not contain concrete type -> rect==zero, ok==false

### Using switch

```
func drawshape(s shape2d) bool{
switch := sh := s.(type){
case rectangle:
drawrect(sh)
case triangle:
drawrect(sh)
}}
```

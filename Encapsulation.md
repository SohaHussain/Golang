### Controlling Access
- can define public functions to allow access to hidden data

```
package data
var x int = 1
func printx(){
fmt.println(x)
}

package main
import "data"
func main(){
data.printx()
}
```

### Controlling access to structs
- hide fields of struct by starting the field name with a lowercase letter
- define public methods which access hidden data

```
package data
type point struct{
x float64
y float64
}
func (p *point) initme(xn,yn float64){
p.x = xn
p.y = yn
}
```

- need initme() to assign hidden data fields

```
package main
import "data"
func main(){
var p data.point
p.initme(3,4)
}
```



### For loops

```
for i:=0; i<10; i++{
fmt.Printf("hi")
}
```

### Switch case

```
switch x{
case 1:
fmt.Printf("1")
case 2:
fmt.Printf("2")
case 3:
fmt.Printf("3")
default:
fmt.Printf("No Case")
}
```

### Tagless Switch case

```
switch{
case x>1:
fmt.printf("case 1")
case x<1:
fmt.Printf("case 2")
default:fmt.Printf("no case")
}
```

### Scan
- reads user input
- takes a pointer as an argument
- typed data is written to pointer
- returns number of scanned items
- eg:

```
var applenum int
fmt.Printf("number of apples?")
num, err :=
fmt.Scan(&applenum)
fmt.Printf(applenum)
```

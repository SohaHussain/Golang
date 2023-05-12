### Creating a Goroutine
- one goroutine is created automatically to execute the main()
- other goroutines are created using the go keyword

```
a=1
foo()
a=2
```

- main goroutine blocks on call to foo()

```
a=1
go foo()
a=2
```

- new goroutine created for foo()
- main goroutine does not block

### Exiting Goroutine
- a goroutine exits when its code is complete
- when the main goroutine is complete, all other goroutines exit
- a goroutine may not complete its execution because main completes early

### Early Exit

```
func main(){
go fmt.Printf("go routine")
fmt.Printf("main routine")
}
```

- if only main routine is printed, it means that main finished before the new go routine started

### Delayed exit

```
func main(){
go fmt.Printf("go routine")
time.Sleep(100 *time.Millisecond)
fmt.Printf("main routine")
}
```

- add a delay to the main routine to give the new routine a chance to complete
- this prints go routinemain routine

### Timing with Go routine
- adding delay to wait for a goroutine is bad
- timing assumptions may be wrong
- timing is non deterministic


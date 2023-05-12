### Channels
- transfer data between goroutines
- channels are typed
- use make() to create a channel 

```
c:=make(chan int)
```

- send an receive data using <- operator
- send data on a channel

```
c <- 3
```

- receive data from a channel

```
x := <- c
```

- example

```
func prod(v1 int, v2 int, c chan int){
c <- v1*v2
}
func main(){
c := make(chan int)
go prod(1,2,c)
go prod(3,4,c)
a := <- c
b := <- c
fmt.Println(a*b)
}
```

### Unbuffered Channel
- unbuffered channels cannot hold data in transit
- default is unbuffered
- sending blocks until data is received
- receiving blocks until data is sent

### Blocking and Synchronization
- channel communication is synchronous
- blocking is same as waiting for communication
- receinving and ignoring the result is same as Wait()

### Channel capacity
- channels can contain a limited number of objects , default size is 0 (unbuffered)
- capacity is the number of objects it acn hold in transit

```
c := make(chan int , 3)
```



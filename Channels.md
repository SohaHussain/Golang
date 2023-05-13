### Iterating through a channel
- common to iteratively read from a channel

```
for i := range c {
fmt.Println(i)
}
```

- continues to read from channel c
- on each iteration a new value is received
- i is assigned to the read value
- iterates until sender call close(c)

### Receiving from multiple goroutines
- multiple channels may be used to receive from multiple sources
- data from both channels may be needed
- read sequentially

```
a := <- c1
b := <- c2
fmt.Println(a*b)
```

### Select Statement
- may have a choice on which data to use incase of multiple channels
- using select statement to wait on the first data from a set of channels

```
select{
case a = <- c1:
fmt.Println(a)
case b = <- c2:
fmt.Println(b)
}
```

- may select either send or receive operations

```
select{
case a = <- inchan:
fmt.Println("received")
case outchan <- b:
fmt.Println("send")
}
```

### Select with an abort channel
- use select with a separate abort channel
- may want to receive data until abort signal is received
```
for{
select{
case a <- c:
fmt.Println(a)
case <- abort :
return
}
}
```

### Synchronization
- using global events whose execution is viewed by all threads simultaneously
- global event is viewed by all tasks at the same time
- synchronization is used to restrict bad interleavings

### Sync WaitGroup
- Sync package contains functions to synchronize between goroutines
- sync.WaitGroup forces a goroutine to wait for other goroutines
- contains an internal counter
- increment counter for each goroutine to wait for
- decrement counter when each goroutine completes
- waiting goroutine cannot continue until counter is 0

### using WaitGroup
- Add() -> increments the counter
- Done() -> decrements the counter
- Wait() -> blocks until counter == 0

```
func foo(wg *sync.WaitGroup){
fmt.Printf("new routine")
wg.Done()
}
func main(){
var wg sync.WaitGroup
wg.Add(1)
go foo(&wg)
wg.Wait()
fmt.Printf("main routine")
}
```

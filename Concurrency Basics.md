### Processes
- an instance of a running program
- things unique to a process:
- 1. memory
- virtual address space
- code, stack, heap, shared libraries
- 2. registers
- program counters, data regs, stack ptr

### Operating systems
- allows many processes to execute concurrently
- processes are switched quickly
- user has the impression of parallelism
- operating system must give processes fair access to resources

### Threads vs process
- threads share some context
- many threads can exist together
- os schedules threads rather than processes

### Goroutines
- like a thread in Go
- many goroutines execute within a single os thread

### Interleavings
- order of execution of tasks is known
- order of execution between concurrent tasks is not known
- interleaving of instructions between tasks is unknown

### OS package file access
- os.Open() -> opens a file. returns a file descriptor (file)
- os.Close() -> closes a file
- os.Read() -> reads from a file into []byte, fills the []byte, controls the amount of read
- os.Write() -> writes a []byte into file

### OS file reading
```
f, err := os.Open("file.txt")
barr := make([]byte,10)
nb, err := f.Read(barr)
f.close()
```
- reads and fills barr
- Read returns number of bytes read
- maybe less than []byte length

### OS file create/write
```
f, err := os.Create("file.txt")
barr := []byte{1,2,3}
nb, err := f.Write(barr)
nb, err := f.WriteString("Hi")
```
- WriteString -> writes a string
- Write -> writes a []byte

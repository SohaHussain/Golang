### File Operations
- open -> get handle for access
- read -> read bytes into []byte
- write -> write []byte into file
- close -> release handle
- seek -> move read/write head

### ioutil file read
```
dat, err := ioutil.ReadFile("file.txt")
```
- dat is []byte filled with contents of entire file
- explicit open/close are not required
- large files can cause problems

### ioutil file write
```
dat = "Hello world!"
err := ioutil.WriteFile("file.txt",dat,0777)
```
- writes []byte to file
- creates a file
- unix style permission bytes

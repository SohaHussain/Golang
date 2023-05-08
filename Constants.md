- expressions whose value is known at compile time
- eg:

```
const x = 1.3
const(
y = 4
z = "Hi"
)
```

- type is inferred from right hand side

### iota
- function used to generate constants
- it generates a set of related but distinct constant
- often represents a property which has several distinct possible values like days of week
- constants must be different but their value is not important
- eg:

```
type grades int
const (
a grades = iota
b
c
d
e
f
)
```

- here each constant is assigned to a unique integer

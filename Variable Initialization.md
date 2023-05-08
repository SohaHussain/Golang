### Type Declarations
- defining an alias or alternate name for a type
- eg: 
-     type Celcius float64
-     var temp Celcius

### Initializing Variables
- initialize in the declaration
- eg:
-     var x int = 100
-     var x = 100
- initialize after declaration
- eg:
-     var x int
-     x = 100
- uninitialized variables have 0 value
- eg:
-     var x int // x = 0
-     var x string // x = ""

### short variable declarations
- can perform declaration and initialization together using := operator
- eg:
-     x := 100
- variable is declared as the type of expression on the left hand side
- can only do short variable declarations inside a function

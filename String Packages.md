### String search functions
- Compare(a,b) -> returns an integer comparing 2 strings lexicographically , 0 if a==b, -1 if a<b, +1 if a>b
- Contains(s,substr) -> returns true if substring is inside s
- HasPrefix(s,prefix) -> returns true if string begins with prefix 
- Index(s,substr) -> returns the index of first occurence of substr in s

strings are unmutable but modified strings are returned

- Replace(s,old,new,n) -> returns a copy of string s with first n instances of old replaced by new
- ToLower(s)
- ToUpper(s)
- TrimSpace(s) -> returns a new string with all leading and trailing whitespaces removed

### Conversion of strings to and from basic data types
- Atoi(s) -> converts string s to int
- Itoa(s) -> converts int(base 10) to string
- FormatFloat(f,fmt,prec,bitSzie) -> converts floating point number to a string
- ParseFloat(s, bitSize) -> converts string to a floating point number

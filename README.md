# strunpack

Go library to unpack a string into a struct using named regexps.

## Example

If you have...

```go
type Typ struct {
	Name string
	Age  int
}
s := "Mary 30"
```

...do this:

```go
var res Typ
if err := Unpack(s, regexp.MustCompile(`(?P<Name>\w+) (?P<Age>\d+)`), &res); {
    panic(err)
}
```

...instead of:

```go
re := regexp.MustCompile(`(?P<Name>\w+) (?P<Age>\d+)`)
m := re.FindStringSubmatch(s)
if len(m) != 3 {
    panic("wrong number of results")
}
name := m[1]
age, err := strconv.Atoi(m[2])
if err != nil {
    panic(err)
}
res := Type{name, age}
```
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
res, err := strunpack.FromString[Typ](`(?P<Name>\w+) (?P<Age>\d+)`).Unpack(s)
if err != nil {
    panic(err)
}
```

--or--

```go
var res Typ
if err := strunpack.Unpack(s, regexp.MustCompile(`(?P<Name>\w+) (?P<Age>\d+)`), &res); {
    panic(err)
}
```

...instead of:

```go
re := regexp.MustCompile(`(\w+) (\d+)`)
m := re.FindStringSubmatch(s)
if len(m) != 3 {
    panic("wrong number of results")
}
name := m[1]
age, err := strconv.Atoi(m[2])
if err != nil {
    panic(err)
}
res := Typ{name, age}
```

### Lazy

You can match the fields in order, too. e.g.

```sh
var res Typ
res, err := strunpack.FromString[Typ]("(\w+) (\d+)").Unpack(s)
if err != nil {
    panic(err)
}
```
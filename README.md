A JSON parser made by implementing the content of this [blog](https://notes.eatonphil.com/writing-a-simple-json-parser.html), made in python, in golang. For learning purposes.

### Input
```go
func main() {
	golang_map := from_string(`{"foo": 1, "bar":[1, "teste", 3, null, true, false, "true", "false", "null"]}`)
	fmt.Println(golang_map)
	fmt.Println(golang_map["bar"])
}

```

### Output
```bash
map[bar:[1 teste 3 <nil> true false true false null] foo:1]
[1 teste 3 <nil> true false true false null]
```

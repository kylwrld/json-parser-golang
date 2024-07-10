package main

import "fmt"

func from_string(str string) map[string]interface{} {
	tokens := Lex(str)
	obj, _ := parse(tokens)

	if obj, ok := obj.(map[string]interface{}); ok {
		return obj
	}
	return nil
}

func main() {
	golang_map := from_string(`{"foo": 1, "bar":[1, "teste", 3, null, true, false, "true", "false", "null"]}`)
	fmt.Println(golang_map)
	fmt.Println(golang_map["bar"])
}

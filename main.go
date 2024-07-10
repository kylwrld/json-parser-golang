package main

import "fmt"

func from_string(str string) map[string]interface{} {
	tokens := Lex(str)
	fmt.Println(tokens)
	obj, _ := parse(tokens)

	if obj, ok := obj.(map[string]interface{}); ok {
		fmt.Println("AA")
		return obj
	}
	return nil
}

func main() {
	m := from_string(`{"foo": 1}`)
	fmt.Println(m)
	fmt.Println(m["foo"])
}

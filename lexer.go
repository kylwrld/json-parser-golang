package main

import (
	"fmt"
	"strconv"
)

var (
	JSON_WHITESPACE = []string{" ", "\t", "\b", "\n", "\r"}
	JSON_SYNTAX     = []string{JSON_COMMA, JSON_COLON, JSON_LEFTBRACKET, JSON_RIGHTBRACKET, JSON_LEFTBRACE, JSON_RIGHTBRACE}

	FALSE_LEN = len("false")
	TRUE_LEN  = len("true")
	NULL_LEN  = len("null")
)

func ContainsByteStringList(bt byte, iter []string) bool {
	for _, value := range iter {
		if string(bt) == value {
			return true
		}
	}
	return false
}

func ContainsByteString(bt byte, iter string) bool {
	for _, value := range iter {
		if rune(bt) == value {
			return true
		}
	}
	return false
}

func ContainsString(bt string, iter []string) bool {
	for _, value := range iter {
		if bt == value {
			return true
		}
	}
	return false
}

func Lex(str string) []interface{} {
	var tokens []interface{}

	for len(str) != 0 {
		var json_string string
		var ok bool
		json_string, str, ok = lex_string(str)
		if ok {
			tokens = append(tokens, json_string)
			continue
		}

		var json_number interface{}
		json_number, str, ok = lex_number(str)
		if ok {
			tokens = append(tokens, json_number)
			continue
		}

		var json_bool bool
		json_bool, str, ok = lex_bool(str)
		if ok {
			tokens = append(tokens, json_bool)
			continue
		}

		str, ok = lex_null(str)
		if ok {
			tokens = append(tokens, nil)
			continue
		}

		char := str[0]
		if ContainsByteStringList(char, JSON_WHITESPACE) {
			str = str[1:]
		} else if ContainsByteStringList(char, JSON_SYNTAX) {
			tokens = append(tokens, string(char))
			str = str[1:]
		} else {
			err := fmt.Sprintf("UNEXPECTED CHARACTER: %c", char)
			panic(err)
		}
	}

	return tokens

}

func lex_string(str string) (string, string, bool) {
	json_string := ""

	if string(str[0]) == JSON_QUOTE {
		str = str[1:]
	} else {
		return "", str, false
	}

	for _, char := range str {
		if string(char) == JSON_QUOTE {
			return json_string, str[len(json_string)+1:], true
		} else {
			json_string += string(char)
		}
	}

	panic("Expected end of string quote")
}

func lex_number(str string) (interface{}, string, bool) {
	json_number := ""

	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "-", "e", "."}

	for _, char := range str {
		if ContainsString(string(char), numbers) {
			json_number += string(char)
		} else { break }
	}

	if len(json_number) == 0 {
		return 0, str, false
	}

	str = str[len(json_number):]
	if ContainsByteString('.', json_number) {
		float, err := strconv.ParseFloat(json_number, 32)
		if err != nil {
			panic(err)
		}
		return float, str, true
	}

	int, err := strconv.ParseInt(json_number, 10, 64)
	if err != nil {
		panic(err)
	}
	return int, str, true
}

func lex_bool(str string) (bool, string, bool) {
	string_len := len(str)

	if string_len >= TRUE_LEN && str[:TRUE_LEN] == "true" {
		return true, str[TRUE_LEN:], true
	} else if string_len >= FALSE_LEN && str[:FALSE_LEN] == "false" {
		return false, str[FALSE_LEN:], true
	}
	return false, str, false
}

func lex_null(str string) (string, bool){
	string_len := len(str)
	if  string_len >= NULL_LEN && str[:NULL_LEN] == "null" {
		return str[NULL_LEN:], true
	}

	return str, false
}

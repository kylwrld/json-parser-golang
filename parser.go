package main

func parse(tokens []interface{}) (interface{}, []interface{}) {
	token := tokens[0]

	switch token.(type) {
	case string:
		if token == JSON_LEFTBRACKET {
			return parse_array(tokens[1:])
		} else if token == JSON_LEFTBRACE {
			return parse_object(tokens[1:])
		} else {
			return token, tokens[1:]
		}
	default:
		return token, tokens[1:]
	}

}

func parse_array(tokens []interface{}) ([]interface{}, []interface{}) {
	var json_array []interface{}

	token := tokens[0]

	switch token.(type) {
	case string:
		if token == JSON_RIGHTBRACKET {
			return json_array, tokens[1:]
		}
	}

	for {
		var json interface{}
		json, tokens = parse(tokens)
		json_array = append(json_array, json)

		token = tokens[0]
		switch token.(type) {
		case string:
			if token == JSON_RIGHTBRACKET {
				return json_array, tokens[1:]
			} else if token != JSON_COMMA {
				panic("Expected comma after object in array")
			} else {
				tokens = tokens[1:]
			}
		default:
			tokens = tokens[1:]
		}
	}
}

func parse_object(tokens []interface{}) (map[string]interface{}, []interface{}) {
	var json_object = make(map[string]interface{})

	token := tokens[0]
	switch token.(type) {
	case string:
		if token == JSON_RIGHTBRACE {
			return json_object, tokens[1:]
		}
	}

	for {
		json_key := tokens[0]
		switch json_key.(type) {
		case string:
			tokens = tokens[1:]
		default:
			panic("Expected string key")
		}

		switch tokens[0].(type) {
		case string:
			if tokens[0] != JSON_COLON {
				panic("Expected colon after key in object")
			}
		}

		var json_value interface{}
		json_value, tokens = parse(tokens[1:])

		if key, ok := json_key.(string); ok {
			json_object[key] = json_value
		}

		token = tokens[0]
		switch token.(type) {
		case string:
			if token == JSON_RIGHTBRACE {
				return json_object, tokens[1:]
			} else if token != JSON_COMMA {
				panic("Expected comma after pair in object")
			}
		}
		tokens = tokens[1:]
	}

	// panic("Expected end-of-object brace")
}
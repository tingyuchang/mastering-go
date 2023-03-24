package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var JSONrecord = `{
	"Flag": true,
	"Array": ["a","b","c"],
	"Entity": {
	  "a1": "b1",
	  "a2": "b2",
	  "Value": -456,
	  "Null": null
	},
	"Message": "Hello Go!"
  }`

func typeSwitch(m map[string]interface{}) {
	for k, v := range m {
		switch c := v.(type) {
		case string:
			fmt.Println("string", k, c)
		case float64:
			fmt.Println("float", k, c)
		case bool:
			fmt.Println("bool", k, c)
		case map[string]interface{}:
			fmt.Println("map[string]interface{}", k, c)
			typeSwitch(v.(map[string]interface{}))
		default:
			fmt.Printf("id %v: %T!\n", k, c)
		}
	}
}
func exploreMap(m map[string]interface{}) {
	for k, v := range m {
		embMap, ok := v.(map[string]interface{})

		if ok {
			fmt.Printf("{\"%v\": \n", k)
			exploreMap(embMap)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v: %v\n", k, v)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("use default JSON record")
	} else {
		JSONrecord = os.Args[1]
	}

	JSONMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(JSONrecord), &JSONMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	exploreMap(JSONMap)
	typeSwitch(JSONMap)
}

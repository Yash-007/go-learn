package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"cost"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJSON() {
	mycourses := []course{
		{"reactjs", 299, "abc.com", "1234", []string{"yash", "agrawal"}},
		{"nodejs", 399, "abc.com", "1234", nil},
		{"expressjs", 199, "abc.com", "1234", nil},
		{"mongodb", 999, "abc.com", "1234", nil},
	}

	finalJSON, err := json.MarshalIndent(mycourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJSON)
}

func DecodeJSON() {
	jsonFromWeb := []byte(`
	       {
                "name": "reactjs",
                "cost": 299,
                "platform": "abc.com",
                "tags": [
                        "yash",
                        "agrawal"
                ]
           }
	 `)

	// var courseJson course
	var courseJson map[string]interface{}
	if json.Valid(jsonFromWeb) {
		json.Unmarshal(jsonFromWeb, &courseJson)
		fmt.Printf("%#v\n\n", courseJson)

		for key, val := range courseJson {
			fmt.Printf("key is %v and value is %v and type is %T\n", key, val, val)
		}
	} else {
		panic("The data in not the valid json")
	}
}

func main() {
	// EncodeJSON()
	DecodeJSON()
}

package main

import (
	"fmt"
	"encoding/json"
)

// Take a string as input and generate a JSON object converted to bytes[] - for input to http.NewRequest
func toJSONBytes(key string, message string) (bytes []byte) {
	payload := map[string]interface{}{key: message}
	byts, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return byts
}

func main(){
	var message = "Hello Slack!"
	bytes := toJSONBytes("text", message)
	fmt.Println(bytes)
	return 
}
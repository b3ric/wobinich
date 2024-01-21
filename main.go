package main

import (
	"encoding/json"
	"fmt"
	c "internal/consensus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	if os.Args[1] == "" {
		fmt.Println("Usage: go run main.go -<proto>")
		os.Exit()
	}

	proto := os.Args[1]

	// protv,
	c.MyIp()

	url := "https://ipinfo.io/" + ipAddress.String() + "/json"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for key, value := range data {
		fmt.Println(key + ": " + strings.Replace(value.(string), "%!d(string=", "", 1))
	}

}

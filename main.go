package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"internal/consensus"
)

func main() {

	if os.Args[1] == "" {
		fmt.Println("Usage: go run main.go <proto>")
		os.Exit(1)
	}

	proto := parseProto()

	ipAddress, err := consensus.MyIp(proto)

	url := "https://ipinfo.io/" + ipAddress.String() + "/json"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

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

func parseProto() int {
	proto, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("argument must be either 4 or 6")
		fmt.Println("defaulting to ipv4")
		proto = 4
	}

	if proto != 4 && proto != 6 {
		fmt.Println("argument must be either 4 or 6")
		panic(err)
	}

	return proto
}

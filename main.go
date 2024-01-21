package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	externalip "github.com/glendc/go-external-ip"
)

func main() {

	consensus := externalip.DefaultConsensus(nil, nil)
	consensus.UseIPProtocol(4)
	ipAddress, err := consensus.ExternalIP()

	if err == nil {
		fmt.Println("Your external IP: " + ipAddress.String())
	}

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

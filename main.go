package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
)

const REST_HOST = "http://ip-api.com/json/"

func main() {
	var host = ""
	if len(os.Args) >= 2 {
		host = os.Args[1]
	} else {
		host = ""
	}
	ip, err := net.ResolveIPAddr("ip4", host)
	if err == nil {

		url := REST_HOST + ip.String()
		res, err := http.Get(url)
		defer res.Body.Close()

		if err == nil {
			wrap := make(map[string]interface{})
			decoder := json.NewDecoder(res.Body)
			decoder.Decode(&wrap)

			if err == nil {
				fmt.Printf("IP: %s\n", wrap["query"])
				fmt.Printf("Country: %s\n", wrap["country"])
				fmt.Printf("Region: %s\n", wrap["regionName"])
				fmt.Printf("City: %s\n", wrap["city"])
				fmt.Printf("ISP: %s\n", wrap["isp"])
			} else {
				panic(err)
			}

		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}

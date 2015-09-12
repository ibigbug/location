package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

const VERSION = "0.0.1"
const REST_HOST = "http://ip-api.com/json/"

func main() {
	host := flag.String("host", "", "A hostname to query.")
	version := flag.Bool("V", false, "Show version")
	flag.Parse()

	if *version {
		fmt.Printf("Version: %s\n", VERSION)
		return
	}

	if *host == "" {
		flag.PrintDefaults()
		return
	}

	url := REST_HOST + *host
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
}

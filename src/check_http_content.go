package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	address	= flag.String("address", "", "Address to check for content")
	content	= flag.String("content", "", "Content to look for in response")
	timeout = flag.Int("timeout", 2, "Response timeout in seconds")
)

func main() {
	flag.Parse()

	if len(*address) == 0 || len (*content) == 0 {
		fmt.Println("Missing required --address or --content parameter")
		os.Exit(4)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout: time.Duration(*timeout) * time.Second,
	}
	resp, err := client.Get(*address)
	if err != nil {
		fmt.Printf("Error making http request. %s\n", err.Error())
		os.Exit(4)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if strings.Contains(strings.ToUpper(string(body)), strings.ToUpper(*content)) {
		fmt.Printf("%s found in response from %s\n", *content, *address)
		os.Exit(0)
	} else {
		fmt.Printf("%s not found in response from %s\n", *content, *address)
		os.Exit(2)
	}
}
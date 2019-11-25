/*
	write a function that gets a url and returns the value of the content-type of the response HTTP header
	It should return an error if it can't perform a GET request

*/

package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("COntent-Type")
	if ctype == "" {
		return "", fmt.Errorf("can't find Content-Type header")
	}

	return resp.Header.Get("COntent-Length"), nil
}

func main() {
	value, err := contentType("https://golang.org")

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(value)
	}
}

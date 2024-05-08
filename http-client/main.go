package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c1 := http.Client{Timeout: time.Duration(2) * time.Second}
	resp, err := c1.Get("https://go.dev/")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	fmt.Printf("Body : %s", body)
}

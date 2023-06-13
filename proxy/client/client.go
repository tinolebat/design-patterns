package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}

	response, err := c.Get("http://localhost:3333")
	if err != nil {
		fmt.Println("Fail to get client: ", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to Read Response body: ", err)
	}
	fmt.Println("Body :\n", string(body))

	// unmarshed_body := json.Unmarshal(body)

}

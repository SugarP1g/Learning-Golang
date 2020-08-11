package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Get方法只有一个参数：url
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	headers := resp.Header
	for k, v := range headers {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}

	fmt.Printf("resp status %s,statusCode %d\n", resp.Status, resp.StatusCode)

	fmt.Printf("resp Proto %s\n", resp.Proto)
}

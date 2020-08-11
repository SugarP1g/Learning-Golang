package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://www.baidu.com", strings.NewReader("mobile=xxxxxxxxx&isRemberPwd=1"))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
}

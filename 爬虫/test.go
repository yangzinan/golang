package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://biz.kaslyju.net")
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

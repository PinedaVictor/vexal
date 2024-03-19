package test

import (
	"log"
	"vx/tools"
)

var client = tools.HTTPClient

func TestHTTP() bool {
	req, err := client.Get("http://localhost:3000/api")
	if err != nil {
		log.Println("error:", err)
	}
	resp, respError := client.Do(req.Request)
	if respError != nil {
		log.Println("Error with server respoonse:", respError)
	}
	log.Println(resp)
	return true
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Response struct {
	Message   string `json:"message"`
	DateTime  string `json:"date_time"`
	IPAddress string `json:"ip_address"`
}

func main() {
	fmt.Println("Content-Type: application/json")
	fmt.Println("")

	ip := os.Getenv("REMOTE_ADDR")
	now := time.Now()

	response := Response{
		Message:   "Hello World",
		DateTime:  now.Format(time.RFC1123),
		IPAddress: ip,
	}

	jsonResponse, _ := json.Marshal(response)

	fmt.Println(string(jsonResponse))
}

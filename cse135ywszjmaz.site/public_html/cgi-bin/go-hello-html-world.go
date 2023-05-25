package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Content-Type: text/html")
	fmt.Println("")

	ip := os.Getenv("REMOTE_ADDR")
	now := time.Now()

	fmt.Printf(`<!DOCTYPE html>
<html>
<head>
	<title>Hello HTML World</title>
</head>
<body>
	<h1>Hello World</h1>
	<p>Current Date/Time: %s</p>
	<p>Your IP address: %s</p>
</body>
</html>
`, now.Format(time.RFC1123), ip)
}

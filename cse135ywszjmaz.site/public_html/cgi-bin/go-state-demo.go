package main

import (
	"fmt"
)

func main() {
	// Headers
	fmt.Println("Content-Type: text/html; charset=utf-8")
	fmt.Println("")

	// HTML content
	fmt.Println(`<!doctype html>
<html>
<head>
  <title>CGI Form</title>
</head>
<body>
  <h1 align="center">Session Test</h1>
  <hr>
  <label for="cgi-lang">CGI using Go</label>
  <form action="/cgi-bin/go-session-1.cgi" method="GET" id="form">
    <label>What is your name? <input type="text" name="username" autocomplete="off"></label>
    <br>
    <input type="submit" value="Test Sessioning">
  </form>
  <a href="/" style="display:inline-block;margin-top:20px;">Home</a> 
</body>
</html>`)
}

package main

import (
	"fmt"
)

func main() {
	// Headers
	fmt.Println("Cache-Control: no-cache")
	fmt.Println("Set-Cookie: destroyed")
	fmt.Println("Content-type: text/html\n")

	// Body - HTML
	fmt.Println("<html>")
	fmt.Println("<head><title>Go Destroy Session</title></head>")
	fmt.Println("<body>")
	fmt.Println("<h1>Session Destroyed</h1>")
	fmt.Println("<br />")
	fmt.Println("<a href=\"/cgi-bin/go-session-1.cgi\">Session Page 1</a>")
	fmt.Println("<br />")
	fmt.Println("<a href=\"/cgi-bin/go-session-2.cgi\">Session Page 2</a>")
	fmt.Println("<br />")
	fmt.Println("<a href=\"/cgi-bin/go-state-demo.cgi\">Go CGI Form</a>")
	fmt.Println("<br /><br />")
	fmt.Println("</body>")
	fmt.Println("</html>")
}

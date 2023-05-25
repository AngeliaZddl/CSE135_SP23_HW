package main

import (
	"fmt"
	"os"
)

func main() {
	// Headers
	fmt.Println("Cache-Control: no-cache")
	fmt.Println("Content-type: text/html\n")

	// Body - HTML
	fmt.Println("<html>")
	fmt.Println("<head><title>Go Sessions</title></head>")
	fmt.Println("<body>")
	fmt.Println("<h1>Go Sessions Page 2</h1>")
	fmt.Println("<table>")

	// Check for Cookie
	if cookie := os.Getenv("HTTP_COOKIE"); len(cookie) > 0 && cookie != "destroyed" {
		fmt.Printf("<tr><td>Cookie:</td><td>%s</td></tr>\n", cookie)
	} else {
		fmt.Println("<tr><td>Cookie:</td><td>None</td></tr>\n")
	}

	fmt.Println("</table>")

	// Links for other pages
	fmt.Println("<br />")
	fmt.Println("<a href=\"/cgi-bin/go-session-1.cgi\">Session Page 1</a>")
	fmt.Println("<br />")
	fmt.Println("<a href=\"/cgi-bin/go-state-demo.cgi\">Go CGI Form</a>")
	fmt.Println("<br /><br />")

	// Destroy Cookie button
	fmt.Println("<form action=\"/cgi-bin/go-destroy-session.cgi\" method=\"get\">")
	fmt.Println("<button type=\"submit\">Destroy Session</button>")
	fmt.Println("</form>")

	fmt.Println("</body>")
	fmt.Println("</html>")
}

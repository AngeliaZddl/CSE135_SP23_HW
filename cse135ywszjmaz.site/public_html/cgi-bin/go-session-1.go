package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Headers
	fmt.Println("Cache-Control: no-cache")

	// Get Name from Environment
	var name string
	if len(os.Getenv("QUERY_STRING")) > 0 {
		query := os.Getenv("QUERY_STRING")
		if strings.HasPrefix(query, "username=") {
			name = query[9:]
			fmt.Printf("Content-type: text/html\n")
			fmt.Printf("Set-Cookie: %s\n\n", name)
		}
	} else {
		fmt.Println("Content-type: text/html\n")
	}

	// Body - HTML
	fmt.Println("<html>")
	fmt.Println("<head><title>Go Sessions</title></head>")
	fmt.Println("<body>")
	fmt.Println("<h1>Go Sessions Page 1</h1>")
	fmt.Println("<table>")

	// First, check for new Cookie, then check for old Cookie
	if len(name) > 0 {
		fmt.Printf("<tr><td>Cookie:</td><td>%s</td></tr>\n", name)
	} else if cookie := os.Getenv("HTTP_COOKIE"); len(cookie) > 0 && cookie != "destroyed" {
		fmt.Printf("<tr><td>Cookie:</td><td>%s</td></tr>\n", cookie)
	} else {
		fmt.Println("<tr><td>Cookie:</td><td>None</td></tr>\n")
	}

	fmt.Println("</table>")

	// Links for other pages
	fmt.Println("<br />")
	fmt.Println("<a href=\"/cgi-bin/go-session-2.cgi\">Session Page 2</a>")
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

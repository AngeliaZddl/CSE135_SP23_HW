package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Content-type: text/html\n\n")
	fmt.Println("<html>")
	fmt.Println("<head>")
	fmt.Println("<style>")
	fmt.Println("body { font-family: Arial, sans-serif; }")
	fmt.Println(".center { text-align: center; }")
	fmt.Println(".header { font-size: 24px; font-weight: bold; }")
	fmt.Println(".content { font-size: 12px; }")
	fmt.Println("</style>")
	fmt.Println("</head>")
	fmt.Println("<body>")
	fmt.Println("<div class='center header'>GET Request Echo</div>")
	fmt.Println("<hr>")
	fmt.Println("<div class='center content'>Query String:</div>")
	fmt.Println("<pre class='content'>")

	req, err := http.NewRequest("GET", os.Getenv("REQUEST_URI"), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.ParseForm()

	fmt.Println(req.Form.Encode())

	fmt.Println("</pre>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}

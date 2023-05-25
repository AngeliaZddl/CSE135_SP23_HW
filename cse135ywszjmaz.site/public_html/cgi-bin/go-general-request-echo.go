package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println("<div class='center header'>General Request Echo</div>")
	fmt.Println("<hr>")
	fmt.Println("<div class='center content'>Method: ", os.Getenv("REQUEST_METHOD"), "</div>")
	fmt.Println("<div class='center content'>Message Body:</div>")
	fmt.Println("<pre class='content'>")

	req, err := http.NewRequest(os.Getenv("REQUEST_METHOD"), os.Getenv("REQUEST_URI"), ioutil.NopCloser(os.Stdin))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", os.Getenv("CONTENT_TYPE"))
	req.Header.Add("Content-Length", os.Getenv("CONTENT_LENGTH"))

	req.ParseForm()

	if req.Method == "POST" || req.Method == "PUT" {
		fmt.Println(req.PostForm.Encode())
	} else {
		fmt.Println(req.Form.Encode())
	}

	fmt.Println("</pre>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}

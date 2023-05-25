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
	fmt.Println(".header { font-size: 36px; font-weight: bold; }")
	fmt.Println(".content { font-size: 20px; }")
	fmt.Println("</style>")
	fmt.Println("</head>")
	fmt.Println("<body>")
	fmt.Println("<div class='center header'>POST Request Echo</div>")
	fmt.Println("<hr>")
	fmt.Println("<div class='center content'>Message Body:</div>")
	fmt.Println("<pre class='content'>")

	req, err := http.NewRequest("POST", os.Getenv("REQUEST_URI"), ioutil.NopCloser(os.Stdin))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", os.Getenv("CONTENT_TYPE"))
	req.Header.Add("Content-Length", os.Getenv("CONTENT_LENGTH"))

	req.ParseForm()

	fmt.Println(req.PostForm.Encode())

	fmt.Println("</pre>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}

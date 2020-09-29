package demoapi

import (
	"fmt"
	"net/http"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	fmt.Fprint(response, "<b><i><u>Hello World</u></i></b>")
}
package main

import (
	"fmt"
	"net/http"
	"os"
)

//TODO find correct ports and return format

const DefaultPort = "7893"

func findPort() string {
	//find server port
	var port string = os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}

	return DefaultPort
}

func main() {

	var port string = findPort()
	fmt.Println("server starting on port", port)

	serverHandler := func(writer http.ResponseWriter, request *http.Request) {
		//echo
		//requestdata data should be read before writing the response

		//Writeheader(statusCode int)
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		writer.WriteHeader(http.StatusOK)

		//Write([]byte) - this should be the request
		request.Write(writer)
	}

	http.HandleFunc("/", serverHandler)

	http.ListenAndServe(":"+port, nil)

}

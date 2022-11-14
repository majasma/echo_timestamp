package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//TODO find correct ports

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

	ServerHandler := func(writer http.ResponseWriter, request *http.Request) {
		//echo
		//requestdata data should be read before writing the response

		//Writeheader(statusCode int)
		writer.Header().Set("Content-Type", "application/json") // normal header

		b, err := io.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}

		//Write([]byte) - this should be the request
		writer.Write(b)
		fmt.Println("served request")
	}

	http.HandleFunc("/", ServerHandler)

	http.ListenAndServe(":"+port, nil)

}

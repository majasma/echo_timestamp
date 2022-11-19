package main

import (
	"fmt"
	"io"
	"net/http"
)


const Port = "54321"


func main() {

	fmt.Println("Server starting on port", Port)

	ServerHandler := func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json") 

		b, err := io.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}

		writer.Write(b)
		fmt.Println("Served request")
	}

	http.HandleFunc("/", ServerHandler)

	http.ListenAndServe(":"+Port, nil)

}

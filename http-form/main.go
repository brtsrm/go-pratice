package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "Sayfa bulunamadı", http.StatusNotFound)
	}
	if request.Method != "GET" {
		http.Error(writer, "method is not suppoerted", http.StatusNotFound)
	}
	_, err := fmt.Fprintf(writer, "Hello!")
	if err != nil {
		return
	}
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "Error %v", err)
		return
	}
	fmt.Fprintf(writer, "Post request Successful")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(writer, "Name = %v", name)
	fmt.Fprintf(writer, "Address = %v ", address)

}

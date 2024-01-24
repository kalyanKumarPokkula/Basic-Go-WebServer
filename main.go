package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter , req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(w , "404 not found" , http.StatusNotFound)
	}

	if req.Method != "GET" {
		http.Error(w , "Method is not supported" , http.StatusNotFound)
	}

	fmt.Fprintf(w , "Hello Welcome to Go server\n" )
	fmt.Fprintf(w , "Thank you for visiting to over Go server")
	

}

func formHandler(w http.ResponseWriter , req *http.Request){
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(w , "ParseForm() err: %v", err)
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := req.FormValue("name")
	email := req.FormValue("email")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
}


func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/" , fileServer)
	http.HandleFunc("/form" , formHandler)
	http.HandleFunc("/hello" , helloHandler)

	fmt.Printf("Starting Server at port 8080\n")
	err := http.ListenAndServe(":8080" , nil)
	if err != nil {
		log.Fatal(err)
	}


}
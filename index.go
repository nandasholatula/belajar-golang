package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Define a handler function for incoming HTTP requests

	// Use the http.HandleFunc function to associate the handler with a route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "john wick",
			"Message": "have a nice day",
		}
		t, err := template.ParseFiles("index.html")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data)

	})

	// Start the HTTP server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

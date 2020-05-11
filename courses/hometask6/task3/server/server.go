// This is a http server - any request is handled by using web browser: http://localhost:8080

package main

import (
	"fmt"
	"net/http"
)

// Store the cookies, all Names are unique
var Cookies = map[string]string{"Name": "Address"}

// This is the handler function
func hello(w http.ResponseWriter, r *http.Request) {

	// 404 to all except root
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get the form from getForm.html file
	if r.Method == "GET" {
		http.ServeFile(w, r, "getForm.html")
	}

	// Take the data from the form, save it, post the data
	if r.Method == "POST" {

		// We parse the values
		ok := r.ParseForm()
		if ok != nil {
			fmt.Println(ok)
			return
		}

		// Read data from the form and store it as key:value pair
		Cookies[r.FormValue("name")] = r.FormValue("address")

		// Write cookies to the form
		message := ""
		for key, value := range Cookies {
			if key == "Name" {
				continue
			}
			cookie := "Name: " + key + " Address: " + value + "\n"
			message = message + cookie
		}

		fmt.Fprintf(w, message)

	}

}

func main() {

	fmt.Println("Server is running")

	http.HandleFunc("/", hello)

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.New("index.html").ParseFiles(
		"index.html",
	))

	err := t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	//	http.HandleFunc("/", helloHandler)

	http.Handle("/", http.FileServer(http.Dir("./")))

	//	http.Handle("/", http.StripPrefix("./public/", http.FileServer(http.Dir("./public"))))

	http.ListenAndServe(":3000", nil)
}

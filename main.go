package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	str := base64.StdEncoding.EncodeToString(iconData)

	t := template.Must(template.New("index.html").Parse(htmlTemplate))

	data := map[string]interface{}{"Image": str}
	err := t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	http.HandleFunc("/", helloHandler)

	http.ListenAndServe(":3000", nil)
}

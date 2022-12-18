package serv

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Serv() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http handle for the site, this will just get the html page from the templates folder and send it (same as get_handle.go in go-web)
		path := r.URL.Path[1:]

		if path == "" {
			http.ServeFile(w, r, "templates/html/index.html")
		}

		// return html files
		// this if statement is the same for the others just replace "html" with what ever the .Contains is looking for
		if strings.Contains(path, ".html") {

			path = "templates/html/" + path
			http.ServeFile(w, r, path)

		} else if strings.Contains(path, ".css") {
			if strings.Contains(path, "admin/") {
				path = "templates/" + path
			} else {
				path = "templates/css/" + path
			}
			http.ServeFile(w, r, path)

		} else if strings.Contains(path, "img") {
			path = "templates/" + path

			content, err := ioutil.ReadFile(path)

			if err != nil {
				fmt.Println("error: ", err)
				fmt.Fprintf(w, "the page your looking for doesn't exist!")
			} else {
				w.Write(content)
			}
		} else if strings.Contains(path, "vid") {
			// this one is abit diffrent, it uses http.ServeFile to give the user thier vidoe data
			path = "templates/" + path

			http.ServeFile(w, r, path)
		} else {
			fmt.Fprintf(w, "the page your looking for doesn't exist!")
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

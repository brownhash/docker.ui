package internals

import (
	"fmt"
	"html/template"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if request.Method == "GET" {
		t, err := template.ParseFiles("templates/dashboard.html")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			fmt.Fprintf(w, "Unable to fetch template:" + err.Error())
			fmt.Println(err)
		}

		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			fmt.Fprintf(w, "Unable to render template:" + err.Error())
			fmt.Println(err)
		}
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

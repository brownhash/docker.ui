package internals

import (
	"./errors"
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
			errors.InternalServerError(w, err)
		}

		err = t.Execute(w, nil)
		if err != nil {
			errors.InternalServerError(w, err)
		}
	} else {
		errors.MethodNotAllowed(w)
	}
}

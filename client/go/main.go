package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request)){
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse Form data", http.StatusInternalServerError)
			return
		}

		//Get form values
		firstName := r.Form.Get("firstname")
		surname := r.Form.Get("surname")
		email := r.Form.Get("email")
		age := r.Form.Get("age")
		password = r.Form.Get("password")

		//Write data to cvs File (Excel -Com[patible)
		file,err :=os.OpenFile("form_data.cvs", os.0_APPEND|os.0_CREATE|os.0_WRONLY, 0644)
		if err != nil {
			http.Error(w, "Failed to open file", http.StatusIntenalServerError)
			return
		}
		defer file.Close()

		writetr :=cvs.NewWriter(file)
		defer writer.Flush()

		err =writer.Write([]string{name, email})
		if err != nil {
			http.Error(w, "Failed to write data to file", http.StatusInternalServerError)
			return
		}
	}
}
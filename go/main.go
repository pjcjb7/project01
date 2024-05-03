package main

import (
	"encoding/csv"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
			return
		}

		//Get form values
		firstName := r.Form.Get("firstname")
		surname := r.Form.Get("surname")
		email := r.Form.Get("email")
		age := r.Form.Get("age")
		password := r.Form.Get("password")

		// Write data to CSV file (Excel-compatible)
		file, err := os.OpenFile("form_data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			http.Error(w, "Failed to open file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		err = writer.Write([]string{firstName, surname, age, email, password})
		if err != nil {
			http.Error(w, "Failed to write data to file", http.StatusInternalServerError)
			return
		}

		// Optionally, you can send a response back to the client
		w.Write([]byte("Data saved successfully!"))
	})

	http.ListenAndServe(":8080", nil)
}

package main

import (
//   "fmt"
  "html/template"
  "log"
  "net/http"
  "strconv"
)

// PageData contains the data to be displayed on the web page.
type PageData struct {
  X        float64
  Y        float64
  Result   float64
  Operator string
}

// HandleCalculate handles the calculation of the two input numbers based on the selected operator.
func HandleCalculate(w http.ResponseWriter, r *http.Request) {
  // Parse the input numbers and operator from the form data.
  x, err := strconv.ParseFloat(r.FormValue("x"), 64)
  if err != nil {
    http.Error(w, "Invalid input for x", http.StatusBadRequest)
    return
  }
  y, err := strconv.ParseFloat(r.FormValue("y"), 64)
  if err != nil {
    http.Error(w, "Invalid input for y", http.StatusBadRequest)
    return
  }
  operator := r.FormValue("operator")

  // Perform the calculation based on the selected operator.
  var result float64
  switch operator {
  case "+":
    result = x + y
  case "-":
    result = x - y
  case "*":
    result = x * y
  case "/":
    result = x / y
  default:
    http.Error(w, "Invalid operator", http.StatusBadRequest)
    return
  }

  // Execute the template with the calculated result.
  data := PageData{X: x, Y: y, Result: result, Operator: operator}
  tmpl := template.Must(template.ParseFiles("calculator.html"))
  if err := tmpl.Execute(w, data); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func main() {
  // Serve static files from the "static" directory.
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  // Use the HandleCalculate function to handle requests to the "/calculate" URL.
  http.HandleFunc("/calculate", HandleCalculate)

  // Start the server on port 8080.
  log.Println("Listening on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}

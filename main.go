package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul bem bonita", Valor: 99, Quantidade: 10},
		{Nome: "Bermuda", Descricao: "Preta bem bonita", Valor: 109.99, Quantidade: 90},
		{Nome: "PC", Descricao: "Preta tozera", Valor: 19.99, Quantidade: 3},
	}

	temp.ExecuteTemplate(w, "index", produtos)
}

package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	models "github.com/YuriBertoldi/Go-exloja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "index", models.BuscaTodosOsProdutos())
}

func NewProduto(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "newproduto", nil)
}

func InsertProduto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		valor, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))

		produto := new(models.Produto)
		produto.Descricao = descricao
		produto.Quantidade = quantidade
		produto.Nome = nome
		produto.Valor = valor

		models.GravarProduto(*produto)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

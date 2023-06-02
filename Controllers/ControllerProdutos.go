package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	Daos "github.com/YuriBertoldi/Go-exloja/daos"
	models "github.com/YuriBertoldi/Go-exloja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "index", Daos.BuscaTodosOsProdutos())
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

		Daos.GravarProduto(*produto)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		valor, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))

		produto := new(models.Produto)
		produto.Id = id
		produto.Descricao = descricao
		produto.Quantidade = quantidade
		produto.Nome = nome
		produto.Valor = valor

		Daos.AtualizarProduto(*produto)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	Daos.DeletarProduto(r.URL.Query().Get("id"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func EditProduto(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "editprod", Daos.CarregarProduto(r.URL.Query().Get("id")))
}

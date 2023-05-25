package models

import bd "github.com/YuriBertoldi/Go-exloja/bd"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Valor      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := bd.ConnectBD()

	QueryDeTodosProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for QueryDeTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var valor float64

		err = QueryDeTodosProdutos.Scan(&id, &nome, &descricao, &valor, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Valor = valor
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func GravarProduto(p Produto) {
	db := bd.ConnectBD()

	QueryGravarProdutos, err := db.Prepare("INSERT INTO PRODUTOS(nome, descricao, valor, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	QueryGravarProdutos.Exec(p.Nome, p.Descricao, p.Valor, p.Quantidade)
	defer db.Close()
}

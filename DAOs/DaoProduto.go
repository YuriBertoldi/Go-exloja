package daos

import (
	bd "github.com/YuriBertoldi/Go-exloja/bd"
	models "github.com/YuriBertoldi/Go-exloja/models"
)

func BuscaTodosOsProdutos() []models.Produto {
	db := bd.ConnectBD()

	QueryDeTodosProdutos, err := db.Query("select * from produtos order by id")
	if err != nil {
		panic(err.Error())
	}

	p := models.Produto{}

	produtos := []models.Produto{}

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

func GravarProduto(p models.Produto) {
	db := bd.ConnectBD()
	QueryGravarProdutos, err := db.Prepare("INSERT INTO PRODUTOS(nome, descricao, valor, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	QueryGravarProdutos.Exec(p.Nome, p.Descricao, p.Valor, p.Quantidade)
	defer db.Close()
}

func AtualizarProduto(p models.Produto) {
	db := bd.ConnectBD()
	QueryAtualizarProdutos, err := db.Prepare("UPDATE PRODUTOS SET nome = $1, descricao = $2, valor = $3, quantidade = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}
	QueryAtualizarProdutos.Exec(p.Nome, p.Descricao, p.Valor, p.Quantidade, p.Id)
	defer db.Close()
}

func DeletarProduto(id string) {
	db := bd.ConnectBD()

	QueryDeletaProdutos, err := db.Prepare("DELETE FROM PRODUTOS WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	QueryDeletaProdutos.Exec(id)
	defer db.Close()
}

func CarregarProduto(id string) models.Produto {
	db := bd.ConnectBD()

	QueryBuscarProdutos, err := db.Query("select * from produtos where id =" + id)
	if err != nil {
		panic(err.Error())
	}

	p := models.Produto{}

	for QueryBuscarProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var valor float64

		err = QueryBuscarProdutos.Scan(&id, &nome, &descricao, &valor, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Valor = valor
		p.Quantidade = quantidade
	}
	defer db.Close()
	return p
}
